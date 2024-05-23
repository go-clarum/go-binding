package server

import (
	"context"
	api "github.com/go-clarum/go-binding/api/agent/api/http"
	coreGrpc "github.com/go-clarum/go-binding/core/grpc"
	"github.com/go-clarum/go-binding/core/logging"
	"github.com/go-clarum/go-binding/core/strings"
	"github.com/go-clarum/go-binding/http/internal/grpc"
)

type Endpoint struct {
	name string
}

type EndpointBuilder struct {
	name           string
	port           int
	timeoutSeconds int
	contentType    string
}

func NewEndpointBuilder() *EndpointBuilder {
	return &EndpointBuilder{}
}

func (builder *EndpointBuilder) Name(name string) *EndpointBuilder {
	builder.name = name
	return builder
}

func (builder *EndpointBuilder) Port(port int) *EndpointBuilder {
	builder.port = port
	return builder
}

func (builder *EndpointBuilder) Timeout(timeoutSeconds int) *EndpointBuilder {
	builder.timeoutSeconds = timeoutSeconds
	return builder
}

func (builder *EndpointBuilder) ContentType(contentType string) *EndpointBuilder {
	builder.contentType = contentType
	return builder
}

func (builder *EndpointBuilder) Build() *Endpoint {
	client := grpc.GetClient()
	apiReq := &api.InitServerRequest{
		Name:           builder.name,
		Port:           int32(builder.port),
		ContentType:    builder.contentType,
		TimeoutSeconds: int32(builder.timeoutSeconds),
	}

	ctx, cancel := context.WithTimeout(context.Background(), coreGrpc.DefaultTimeout)
	defer cancel()

	response, err := client.InitServerEndpoint(ctx, apiReq)
	if err != nil {
		logging.Fatalf("connection to agent failed - %s", err)
	}

	if strings.IsNotBlank(response.Error) {
		logging.Errorf("error while creating endpoint - %s", response.Error)
	}

	return &Endpoint{
		name: builder.name,
	}
}
