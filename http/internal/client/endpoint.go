package client

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
	baseUrl        string
	contentType    string
	timeoutSeconds int
}

func NewEndpointBuilder() *EndpointBuilder {
	return &EndpointBuilder{}
}

func (builder *EndpointBuilder) Name(name string) *EndpointBuilder {
	builder.name = name
	return builder
}

func (builder *EndpointBuilder) BaseUrl(baseUrl string) *EndpointBuilder {
	builder.baseUrl = baseUrl
	return builder
}

func (builder *EndpointBuilder) ContentType(contentType string) *EndpointBuilder {
	builder.contentType = contentType
	return builder
}

func (builder *EndpointBuilder) Timeout(timeoutSeconds int) *EndpointBuilder {
	builder.timeoutSeconds = timeoutSeconds
	return builder
}

func (builder *EndpointBuilder) Build() *Endpoint {
	client := grpc.GetClient()
	apiReq := &api.InitClientRequest{
		Name:           builder.name,
		BaseUrl:        builder.baseUrl,
		ContentType:    builder.contentType,
		TimeoutSeconds: int32(builder.timeoutSeconds),
	}

	ctx, cancel := context.WithTimeout(context.Background(), coreGrpc.DefaultTimeout)
	defer cancel()

	response, err := client.InitClientEndpoint(ctx, apiReq)
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
