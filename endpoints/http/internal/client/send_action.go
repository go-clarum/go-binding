package client

import (
	"context"
	"errors"
	api "github.com/go-clarum/go-binding/agent/api/http"
	coreGrpc "github.com/go-clarum/go-binding/core/grpc"
	"github.com/go-clarum/go-binding/core/strings"
	"github.com/go-clarum/go-binding/endpoints/http/internal/grpc"
	"github.com/go-clarum/go-binding/endpoints/http/internal/request"
	"testing"
)

// SendActionBuilder used to configure a send action on a client endpoint without the context of a test
// the method chain will end with the .Message() method which will return an error.
// The error will be a problem encountered during sending.
type SendActionBuilder struct {
	endpoint *Endpoint
}

// TestSendActionBuilder used to configure a send action on a client endpoint with the context of a test
// the method chain will end with the .Message() method which will not return anything.
// Any error encountered during sending will fail the test by calling t.Error().
type TestSendActionBuilder struct {
	test *testing.T
	SendActionBuilder
}

func (testBuilder *TestSendActionBuilder) Request(testReq *request.HttpRequest) {
	errorResponse := doClientSendRequest(testReq, testBuilder.endpoint.name)

	if errorResponse != nil {
		testBuilder.test.Error(errorResponse)
	}
}

func (builder *SendActionBuilder) Request(testReq *request.HttpRequest) error {
	return doClientSendRequest(testReq, builder.endpoint.name)
}

func doClientSendRequest(testReq *request.HttpRequest, endpointName string) error {
	client := grpc.GetClient()
	apiReq := &api.ClientSendActionRequest{
		Name:         "not yet implemented",
		Url:          testReq.Url,
		Path:         testReq.Path,
		Method:       testReq.Method,
		QueryParams:  grpc.ParseQueryParams(testReq.QueryParams),
		Headers:      testReq.Headers,
		Payload:      testReq.MessagePayload,
		EndpointName: endpointName,
	}

	ctx, cancel := context.WithTimeout(context.Background(), coreGrpc.DefaultTimeout)
	defer cancel()

	res, err := client.ClientSendAction(ctx, apiReq)

	if err != nil {
		return err
	}
	if strings.IsNotBlank(res.Error) {
		return errors.New(res.Error)
	}

	return nil
}
