package grpc

import (
	api "github.com/go-clarum/go-binding/api/http"
	coreGrpc "github.com/go-clarum/go-binding/core/grpc"
)

var client api.HttpApiClient

// initiate client on demand
func GetClient() api.HttpApiClient {
	if client == nil {
		client = createClient()
	}

	return client
}

func createClient() api.HttpApiClient {
	return api.NewHttpApiClient(coreGrpc.GetConnection())
}

func ParseQueryParams(apiQueryParams map[string][]string) map[string]*api.StringsList {
	result := make(map[string]*api.StringsList)

	if apiQueryParams != nil {
		for key, value := range apiQueryParams {
			newList := &api.StringsList{
				Values: value,
			}
			result[key] = newList
		}
	}

	return result
}
