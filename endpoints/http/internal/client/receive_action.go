package client

import (
	"context"
	"errors"
	api "github.com/go-clarum/go-binding/agent/api/http"
	"github.com/go-clarum/go-binding/core/strings"
	"github.com/go-clarum/go-binding/endpoints/http/internal/grpc"
	"github.com/go-clarum/go-binding/endpoints/http/internal/response"
	"testing"
)

// ReceiveActionBuilder used to configure a receive action on a client endpoint without the context of a test
// the method chain will end with the .Message() method which will return an error.
// The error will be a problem encountered during receiving or a validation error.
type ReceiveActionBuilder struct {
	endpoint *Endpoint
	request  *api.ClientReceiveActionRequest
}

// TestReceiveActionBuilder used to configure a receive action on a client endpoint with the context of a test
// the method chain will end with the .Message() method which will not return anything.
// Any error encountered during receiving or validating will fail the test by calling t.Error().
type TestReceiveActionBuilder struct {
	test *testing.T
	ReceiveActionBuilder
}

func (testBuilder *TestReceiveActionBuilder) Json() *TestReceiveActionBuilder {
	testBuilder.request.PayloadType = api.PayloadType_Json
	return testBuilder
}

func (builder *ReceiveActionBuilder) Json() *ReceiveActionBuilder {
	builder.request.PayloadType = api.PayloadType_Json
	return builder
}

func (testBuilder *TestReceiveActionBuilder) Response(testRes *response.HttpResponse) {
	if err := doClientReceiveRequest(testRes, testBuilder.request, testBuilder.endpoint.name); err != nil {
		testBuilder.test.Error(err)
	}
}

// TODO: this previously also returned the HttpResponse
func (builder *ReceiveActionBuilder) Response(testRes *response.HttpResponse) error {
	return doClientReceiveRequest(testRes, builder.request, builder.endpoint.name)
}

func doClientReceiveRequest(testRes *response.HttpResponse, apiReq *api.ClientReceiveActionRequest,
	endpointName string) error {
	client := grpc.GetClient()

	apiReq.Name = "not yet implemented"
	apiReq.StatusCode = int32(testRes.StatusCode)
	apiReq.Headers = testRes.Headers
	apiReq.Payload = testRes.MessagePayload
	apiReq.EndpointName = endpointName

	// TODO: timeout context
	res, err := client.ClientReceiveAction(context.Background(), apiReq)

	if err != nil {
		return err
	}
	if strings.IsNotBlank(res.Error) {
		return errors.New(res.Error)
	}

	return nil
}
