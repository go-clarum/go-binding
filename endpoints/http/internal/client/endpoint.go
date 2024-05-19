package client

import (
	"context"
	api "github.com/go-clarum/go-binding/agent/api/http"
	"github.com/go-clarum/go-binding/endpoints/http/internal/grpc"
	"log"
)

type Endpoint struct {
	name string
}

type EndpointBuilder struct {
	request *api.InitClientRequest
}

func NewEndpointBuilder() *EndpointBuilder {
	return &EndpointBuilder{}
}

func (builder *EndpointBuilder) Name(name string) *EndpointBuilder {
	builder.request.Name = name
	return builder
}

func (builder *EndpointBuilder) BaseUrl(baseUrl string) *EndpointBuilder {
	builder.request.BaseUrl = baseUrl
	return builder
}

func (builder *EndpointBuilder) ContentType(contentType string) *EndpointBuilder {
	builder.request.ContentType = contentType
	return builder
}

func (builder *EndpointBuilder) Timeout(timeout int32) *EndpointBuilder {
	builder.request.TimeoutSeconds = timeout
	return builder
}

func (builder *EndpointBuilder) Build() *Endpoint {
	client := grpc.GetClient()
	response, err := client.InitClientEndpoint(context.Background(), builder.request)
	if err != nil {
		log.Fatalf("connection to agent failed - %s", err)
	}

	if len(response.Error) != 0 {
		log.Printf("error while creating endpoint - %s", response.Error)
	}

	return &Endpoint{
		name: builder.request.Name,
	}
}
