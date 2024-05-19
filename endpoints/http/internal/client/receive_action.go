package client

import (
	"context"
	"errors"
	api "github.com/go-clarum/go-binding/agent/api/http"
	coreGrpc "github.com/go-clarum/go-binding/core/grpc"
	"github.com/go-clarum/go-binding/core/strings"
	"github.com/go-clarum/go-binding/endpoints/http/internal/grpc"
	"github.com/go-clarum/go-binding/endpoints/http/internal/response"
	"testing"
)

// ReceiveActionBuilder used to configure a receive action on a client endpoint without the context of a test
// the method chain will end with the .Message() method which will return an error.
// The error will be a problem encountered during receiving or a validation error.
type ReceiveActionBuilder struct {
	endpoint    *Endpoint
	payloadType api.PayloadType
}

// TestReceiveActionBuilder used to configure a receive action on a client endpoint with the context of a test
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

func (testBuilder *TestReceiveActionBuilder) Response(testRes *response.HttpResponse) {
	if err := testBuilder.doClientReceiveRequest(testRes); err != nil {
		testBuilder.test.Error(err)
	}
}

func (builder *ReceiveActionBuilder) Response(testRes *response.HttpResponse) error {
	return builder.doClientReceiveRequest(testRes)
}

func (builder *ReceiveActionBuilder) doClientReceiveRequest(testRes *response.HttpResponse) error {
	client := grpc.GetClient()

	apiReq := &api.ClientReceiveActionRequest{
		Name:         "not yet implemented",
		StatusCode:   int32(testRes.StatusCode),
		Headers:      testRes.Headers,
		Payload:      testRes.MessagePayload,
		PayloadType:  builder.payloadType,
		EndpointName: builder.endpoint.name,
	}

	ctx, cancel := context.WithTimeout(context.Background(), coreGrpc.DefaultTimeout)
	defer cancel()

	res, err := client.ClientReceiveAction(ctx, apiReq)

	if err != nil {
		return err
	}
	if strings.IsNotBlank(res.Error) {
		return errors.New(res.Error)
	}

	return nil
}
