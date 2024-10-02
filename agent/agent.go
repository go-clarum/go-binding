package agent

import (
	"context"
	"errors"
	"fmt"
	api "github.com/go-clarum/go-binding/api/agent"
	coreGrpc "github.com/go-clarum/go-binding/core/grpc"
	"github.com/go-clarum/go-binding/core/logging"
	"google.golang.org/grpc"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type AgentService interface {
	Initiate()
	Logs()
	Shutdown()
}

type service struct {
	client         api.AgentServiceClient
	conn           *grpc.ClientConn
	logger         *logging.Logger
	cmd            *exec.Cmd
	shutdownSignal chan bool
}

func NewAgentService() AgentService {
	conn := coreGrpc.GetConnection()
	client := getClient(conn)
	signal := make(chan bool)

	return &service{
		client:         client,
		conn:           conn,
		logger:         logging.NewLogger("AgentService"),
		shutdownSignal: signal,
	}
}

func (s *service) Initiate() {
	agentDir := filepath.Join(".clarum", "agent")
	if _, err := os.Stat(agentDir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(agentDir, os.ModePerm); err != nil {
			s.logger.Fatalf("unable to create agent dir - %s", err)
		}
	}

	// TODO: at this point the agent decide on two things based on current config:
	// 		- if set to mode local - the agent will be downloaded for the current OS & ARCH
	//		- if set to mode remote - a gRPC connection is attempted to the configured URL
	agentFilePath := filepath.Join(agentDir, "clarum-agent") // add current version to file path
	// TODO: implement local

	s.cmd = exec.Command(filepath.Join(".", agentFilePath))
	s.logger.Info("starting clarum-agent")
	if err := s.cmd.Start(); err != nil {
		s.logger.Fatalf("unable to start agent - %s", err)
	}
	s.logger.Infof("clarum-agent started - pid %d", s.cmd.Process.Pid)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	timeout := time.After(5 * time.Second)
	for {
		select {
		case <-timeout:
			s.logger.Fatalf("timeout reached, unable to connect to agent GRPC server")
		case <-ticker.C:
			req := &api.StatusRequest{}
			res, err := s.client.Status(ctx, req)
			if err == nil {
				s.logger.Infof("connected to clarum-agent version %s", res.Version)
				return
			}
			s.logger.Warnf("unable to connect to clarum-agent - %s", err)
		}
	}
}

func (s *service) Logs() {
	ctx := context.Background()
	req := &api.LogsRequest{ListenerName: "go-binding"}
	stream, err := s.client.Logs(ctx, req)

	if err != nil {
		msg := fmt.Sprintf("error while creating logs channel: %s", err)
		logging.Fatalf(msg)
		panic(msg)
	}

	go func() {
		for {
			select {
			case <-s.shutdownSignal:
				err := stream.CloseSend()
				if err != nil {
					logging.Fatalf("failed to close logs channel: %s", err)
				}
				return
			default:
				res, err := stream.Recv()
				if err != nil {
					logging.Fatalf("error while receiving logs: %s", err)
				}

				logging.Log(res.Message)
			}
		}
	}()
}

func (s *service) Shutdown() {
	s.logger.Info("sending shutdown signal to agent")
	s.shutdownSignal <- true

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &api.ShutdownRequest{}
	_, err := s.client.Shutdown(ctx, req)

	if err != nil {
		s.logger.Errorf("unable to shut down agent - %s", err)
		s.handleFailedAgentShutdown()
	}

	if err := s.conn.Close(); err != nil {
		s.logger.Errorf("error while closing grpc connection to agent - %s", err)
	}
}

func (s *service) handleFailedAgentShutdown() {
	s.logger.Debugf("attempting to kill process %d", s.cmd.Process.Pid)
	err := s.cmd.Process.Kill()
	if err != nil {
		s.logger.Errorf("error while killing process - %s", err)
	}
}

func getClient(connection *grpc.ClientConn) api.AgentServiceClient {
	return api.NewAgentServiceClient(connection)
}
