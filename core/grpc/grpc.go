package grpc

import (
	"context"
	"fmt"
	"github.com/go-clarum/go-binding/core/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	defaultAgentPort = "9091"
	defaultAgentHost = "localhost"
	DefaultTimeout   = 5 * time.Second
)

func GetConnection() *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(unaryInterceptor),
	}

	address := fmt.Sprintf("%s:%s", defaultAgentHost, defaultAgentPort)
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("error while calling agent service - %s", err)
	}

	return conn
}

func unaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	logging.Debugf("executing GRPC method [%s]", method)

	err := invoker(ctx, method, req, reply, cc, opts...)

	logging.Debugf("finished GRPC method [%s], duration [%dms], error [%v]", method, time.Since(start).Milliseconds(), err)
	return err
}
