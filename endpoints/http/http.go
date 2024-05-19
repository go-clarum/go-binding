package http

import (
	"github.com/go-clarum/go-binding/endpoints/http/internal/client"
	"github.com/go-clarum/go-binding/endpoints/http/internal/server"
)

type EndpointBuilder struct {
}

func Http() *EndpointBuilder {
	return &EndpointBuilder{}
}

func (heb *EndpointBuilder) Client() *client.EndpointBuilder {
	return client.NewEndpointBuilder()
}

func (heb *EndpointBuilder) Server() *server.EndpointBuilder {
	return server.NewEndpointBuilder()
}
