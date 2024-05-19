package agent

import (
	"context"
	"errors"
	api "github.com/go-clarum/go-binding/agent/api/agent"
	coreGrpc "github.com/go-clarum/go-binding/core/grpc"
	"github.com/go-clarum/go-binding/core/logging"
	"google.golang.org/grpc"
	"os"
	"os/exec"
	"path"
	"time"
)

type AgentService interface {
	Initiate(agentExecutable []byte, fileName string)
	Shutdown()
}

type service struct {
	client api.AgentServiceClient
	conn   *grpc.ClientConn
	logger *logging.Logger
	cmd    *exec.Cmd
}

func NewAgentService() AgentService {
	conn := coreGrpc.GetConnection()
	client := getClient(conn)

	return &service{
		client: client,
		conn:   conn,
		logger: logging.NewLogger("AgentService"),
	}
}

func (s *service) Initiate(agentExecutable []byte, fileName string) {
	agentDir := path.Join(".clarum", "agent")
	if _, err := os.Stat(agentDir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(agentDir, os.ModePerm); err != nil {
			s.logger.Fatalf("unable to create agent dir - %s", err)
		}
	}

	agentFilePath := path.Join(agentDir, fileName)
	if _, err := os.Stat(agentFilePath); errors.Is(err, os.ErrNotExist) {
		if err := os.WriteFile(agentFilePath, agentExecutable, 0755); err != nil {
			s.logger.Fatalf("unable to write agent on disk - %s", err)
		}
	}

	s.cmd = exec.Command("./" + agentFilePath)
	s.logger.Info("starting clarum-agent")
	if err := s.cmd.Start(); err != nil {
		s.logger.Fatalf("unable to start agent - %s", err)
	}
	s.logger.Infof("clarum-agent started - pid %d", s.cmd.Process.Pid)

	// TODO: polling is be better than waiting
	time.Sleep(1 * time.Second)

	req := &api.StatusRequest{}
	// TODO: use timeout context
	res, err := s.client.Status(context.Background(), req)
	if err != nil {
		s.logger.Fatalf("unable to reach agent GRPC server - %s", err)
	}

	s.logger.Infof("connected to clarum-agent version %s", res.Version)
}

func (s *service) Shutdown() {
	s.logger.Info("shutting down agent")

	req := &api.ShutdownRequest{}
	_, err := s.client.Shutdown(context.Background(), req)

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
