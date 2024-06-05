package server

import (
	"context"
	"errors"
	api "github.com/go-clarum/go-binding/api/http"
	coreGrpc "github.com/go-clarum/go-binding/core/grpc"
	"github.com/go-clarum/go-binding/core/strings"
	"github.com/go-clarum/go-binding/http/internal/grpc"
	"github.com/go-clarum/go-binding/http/response"
	"testing"
)

// SendActionBuilder used to configure a send action on a server endpoint without the context of a test
// the method chain will end with the .Message() method which will return an error.
// The error will be a problem encountered during sending.
type SendActionBuilder struct {
	name     string
	endpoint *Endpoint
}

// TestSendActionBuilder used to configure a send action on a server endpoint with the context of a test
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

func (testBuilder *TestSendActionBuilder) Response(testRes *response.HttpResponse) {
	if err := testBuilder.doServerSendResponse(testRes); err != nil {
		testBuilder.test.Error(err)
	}
}

func (builder *SendActionBuilder) Response(testRes *response.HttpResponse) error {
	return builder.doServerSendResponse(testRes)
}

func (builder *SendActionBuilder) doServerSendResponse(testRes *response.HttpResponse) error {
	client := grpc.GetClient()

	apiReq := &api.ServerSendActionRequest{
		Name:         builder.name,
		StatusCode:   int32(testRes.StatusCode),
		Headers:      testRes.Headers,
		Payload:      testRes.MessagePayload,
		EndpointName: builder.endpoint.name,
	}

	ctx, cancel := context.WithTimeout(context.Background(), coreGrpc.DefaultTimeout)
	defer cancel()

	res, err := client.ServerSendAction(ctx, apiReq)
	if err != nil {
		return err
	}
	if strings.IsNotBlank(res.Error) {
		return errors.New(res.Error)
	}

	return nil
}
