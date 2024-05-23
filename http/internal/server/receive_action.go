package server

import (
	"context"
	"errors"
	api "github.com/go-clarum/go-binding/api/agent/api/http"
	coreGrpc "github.com/go-clarum/go-binding/core/grpc"
	"github.com/go-clarum/go-binding/core/strings"
	"github.com/go-clarum/go-binding/http/internal/grpc"
	"github.com/go-clarum/go-binding/http/request"
	"testing"
)

// ReceiveActionBuilder used to configure a receive action on a server endpoint without the context of a test
// the method chain will end with the .Message() method which will return an error.
// The error will be a problem encountered during receiving or a validation error.
type ReceiveActionBuilder struct {
	name        string
	endpoint    *Endpoint
	payloadType api.PayloadType
}

// TestReceiveActionBuilder used to configure a receive action on a server endpoint with the context of a test
// the method chain will end with the .Message() method which will not return anything.
// Any error encountered during receiving or validating will fail the test by calling t.Error().
type TestReceiveActionBuilder struct {
	test *testing.T
	ReceiveActionBuilder
}

func (testBuilder *TestReceiveActionBuilder) Json() *TestReceiveActionBuilder {
	testBuilder.payloadType = api.PayloadType_Json
	return testBuilder
}

func (builder *ReceiveActionBuilder) Json() *ReceiveActionBuilder {
	builder.payloadType = api.PayloadType_Json
	return builder
}

func (testBuilder *TestReceiveActionBuilder) Name(name string) *TestReceiveActionBuilder {
	testBuilder.name = name
	return testBuilder
}

func (builder *ReceiveActionBuilder) Name(name string) *ReceiveActionBuilder {
	builder.name = name
	return builder
}

func (testBuilder *TestReceiveActionBuilder) Request(testReq *request.HttpRequest) {
	if err := testBuilder.doServerReceiveRequest(testReq); err != nil {
		testBuilder.test.Error(err)
	}
}

func (builder *ReceiveActionBuilder) Request(testReq *request.HttpRequest) error {
	return builder.doServerReceiveRequest(testReq)
}

func (builder *ReceiveActionBuilder) doServerReceiveRequest(testReq *request.HttpRequest) error {
	client := grpc.GetClient()
	apiReq := &api.ServerReceiveActionRequest{
		Name:         builder.name,
		Method:       testReq.Method,
		Path:         testReq.Path,
		Url:          testReq.Url,
		QueryParams:  grpc.ParseQueryParams(testReq.QueryParams),
		Headers:      testReq.Headers,
		Payload:      testReq.MessagePayload,
		PayloadType:  builder.payloadType,
		EndpointName: builder.endpoint.name,
	}

	ctx, cancel := context.WithTimeout(context.Background(), coreGrpc.DefaultTimeout)
	defer cancel()

	res, err := client.ServerReceiveAction(ctx, apiReq)

	if err != nil {
		return err
	}
	if strings.IsNotBlank(res.Error) {
		return errors.New(res.Error)
	}

	return nil
}
