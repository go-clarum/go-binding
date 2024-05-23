package http

import (
	"github.com/go-clarum/go-binding/http/internal/client"
	"github.com/go-clarum/go-binding/http/internal/server"
)

type EndpointBuilder struct {
}

func (heb *EndpointBuilder) Client() *client.EndpointBuilder {
	return client.NewEndpointBuilder()
}

func (heb *EndpointBuilder) Server() *server.EndpointBuilder {
	return server.NewEndpointBuilder()
}
