package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	defaultAgentPort = "9091"
	defaultAgentHost = "localhost"
)

func GetConnection() *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	address := fmt.Sprintf("%s:%s", defaultAgentHost, defaultAgentPort)
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("error while calling agent service - %s", err)
	}

	return conn
}
