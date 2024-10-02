package client

import (
	"context"
	"errors"
	api "github.com/go-clarum/go-binding/api/http"
	coreGrpc "github.com/go-clarum/go-binding/core/grpc"
	"github.com/go-clarum/go-binding/core/strings"
	"github.com/go-clarum/go-binding/http/internal/grpc"
	"github.com/go-clarum/go-binding/http/request"
	"testing"
)

// SendActionBuilder used to configure a send action on a client endpoint without the context of a test
// the method chain will end with the .Message() method which will return an error.
// The error will be a problem encountered during sending.
type SendActionBuilder struct {
	name     string
	endpoint *Endpoint
}

// TestSendActionBuilder used to configure a send action on a client endpoint with the context of a test
// the method chain will end with the .Message() method which will not return anything.
// Any error encountered during sending will fail the test by calling t.Error().
type TestSendActionBuilder struct {
	test *testing.T
	SendActionBuilder
}

func (testBuilder *TestSendActionBuilder) Name(name string) *TestSendActionBuilder {
	testBuilder.name = name
	return testBuilder
}

func (builder *SendActionBuilder) Name(name string) *SendActionBuilder {
	builder.name = name
	return builder
}

func (testBuilder *TestSendActionBuilder) Request(testReq *request.HttpRequest) {
	errorResponse := testBuilder.doClientSendRequest(testReq)

	if errorResponse != nil {
		testBuilder.test.Error(errorResponse)
	}
}

func (builder *SendActionBuilder) Request(testReq *request.HttpRequest) error {
	return builder.doClientSendRequest(testReq)
}

func (builder *SendActionBuilder) doClientSendRequest(testReq *request.HttpRequest) error {
	if testReq == nil {
		return errors.New("request to send cannot be nil")
	}

	client := grpc.GetClient()
	apiReq := &api.ClientSendActionRequest{
		Name:         builder.name,
		Url:          testReq.Url,
		Path:         testReq.Path,
		Method:       testReq.Method,
		QueryParams:  grpc.ParseQueryParams(testReq.QueryParams),
		Headers:      testReq.Headers,
		Payload:      testReq.MessagePayload,
		EndpointName: builder.endpoint.name,
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
