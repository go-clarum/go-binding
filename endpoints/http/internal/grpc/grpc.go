package grpc

import (
	api "github.com/go-clarum/go-binding/agent/api/http"
	coreGrpc "github.com/go-clarum/go-binding/core/grpc"
)

var client api.HttpServiceClient

// initiate client on demand
func GetClient() api.HttpServiceClient {
	if client == nil {
		client = createClient()
	}

	return client
}

func createClient() api.HttpServiceClient {
	return api.NewHttpServiceClient(coreGrpc.GetConnection())
}
