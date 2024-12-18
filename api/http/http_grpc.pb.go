// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.3
// source: application/api/http/http.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HttpApiClient is the client API for HttpApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HttpApiClient interface {
	InitClientEndpoint(ctx context.Context, in *InitClientRequest, opts ...grpc.CallOption) (*InitClientResponse, error)
	InitServerEndpoint(ctx context.Context, in *InitServerRequest, opts ...grpc.CallOption) (*InitServerResponse, error)
	ClientSendAction(ctx context.Context, in *ClientSendActionRequest, opts ...grpc.CallOption) (*ClientSendActionResponse, error)
	ClientReceiveAction(ctx context.Context, in *ClientReceiveActionRequest, opts ...grpc.CallOption) (*ClientReceiveActionResponse, error)
	ServerSendAction(ctx context.Context, in *ServerSendActionRequest, opts ...grpc.CallOption) (*ServerSendActionResponse, error)
	ServerReceiveAction(ctx context.Context, in *ServerReceiveActionRequest, opts ...grpc.CallOption) (*ServerReceiveActionResponse, error)
}

type httpApiClient struct {
	cc grpc.ClientConnInterface
}

func NewHttpApiClient(cc grpc.ClientConnInterface) HttpApiClient {
	return &httpApiClient{cc}
}

func (c *httpApiClient) InitClientEndpoint(ctx context.Context, in *InitClientRequest, opts ...grpc.CallOption) (*InitClientResponse, error) {
	out := new(InitClientResponse)
	err := c.cc.Invoke(ctx, "/HttpApi/InitClientEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpApiClient) InitServerEndpoint(ctx context.Context, in *InitServerRequest, opts ...grpc.CallOption) (*InitServerResponse, error) {
	out := new(InitServerResponse)
	err := c.cc.Invoke(ctx, "/HttpApi/InitServerEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpApiClient) ClientSendAction(ctx context.Context, in *ClientSendActionRequest, opts ...grpc.CallOption) (*ClientSendActionResponse, error) {
	out := new(ClientSendActionResponse)
	err := c.cc.Invoke(ctx, "/HttpApi/ClientSendAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpApiClient) ClientReceiveAction(ctx context.Context, in *ClientReceiveActionRequest, opts ...grpc.CallOption) (*ClientReceiveActionResponse, error) {
	out := new(ClientReceiveActionResponse)
	err := c.cc.Invoke(ctx, "/HttpApi/ClientReceiveAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpApiClient) ServerSendAction(ctx context.Context, in *ServerSendActionRequest, opts ...grpc.CallOption) (*ServerSendActionResponse, error) {
	out := new(ServerSendActionResponse)
	err := c.cc.Invoke(ctx, "/HttpApi/ServerSendAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpApiClient) ServerReceiveAction(ctx context.Context, in *ServerReceiveActionRequest, opts ...grpc.CallOption) (*ServerReceiveActionResponse, error) {
	out := new(ServerReceiveActionResponse)
	err := c.cc.Invoke(ctx, "/HttpApi/ServerReceiveAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HttpApiServer is the server API for HttpApi service.
// All implementations must embed UnimplementedHttpApiServer
// for forward compatibility
type HttpApiServer interface {
	InitClientEndpoint(context.Context, *InitClientRequest) (*InitClientResponse, error)
	InitServerEndpoint(context.Context, *InitServerRequest) (*InitServerResponse, error)
	ClientSendAction(context.Context, *ClientSendActionRequest) (*ClientSendActionResponse, error)
	ClientReceiveAction(context.Context, *ClientReceiveActionRequest) (*ClientReceiveActionResponse, error)
	ServerSendAction(context.Context, *ServerSendActionRequest) (*ServerSendActionResponse, error)
	ServerReceiveAction(context.Context, *ServerReceiveActionRequest) (*ServerReceiveActionResponse, error)
	mustEmbedUnimplementedHttpApiServer()
}

// UnimplementedHttpApiServer must be embedded to have forward compatible implementations.
type UnimplementedHttpApiServer struct {
}

func (UnimplementedHttpApiServer) InitClientEndpoint(context.Context, *InitClientRequest) (*InitClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitClientEndpoint not implemented")
}
func (UnimplementedHttpApiServer) InitServerEndpoint(context.Context, *InitServerRequest) (*InitServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitServerEndpoint not implemented")
}
func (UnimplementedHttpApiServer) ClientSendAction(context.Context, *ClientSendActionRequest) (*ClientSendActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientSendAction not implemented")
}
func (UnimplementedHttpApiServer) ClientReceiveAction(context.Context, *ClientReceiveActionRequest) (*ClientReceiveActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientReceiveAction not implemented")
}
func (UnimplementedHttpApiServer) ServerSendAction(context.Context, *ServerSendActionRequest) (*ServerSendActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerSendAction not implemented")
}
func (UnimplementedHttpApiServer) ServerReceiveAction(context.Context, *ServerReceiveActionRequest) (*ServerReceiveActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerReceiveAction not implemented")
}
func (UnimplementedHttpApiServer) mustEmbedUnimplementedHttpApiServer() {}

// UnsafeHttpApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HttpApiServer will
// result in compilation errors.
type UnsafeHttpApiServer interface {
	mustEmbedUnimplementedHttpApiServer()
}

func RegisterHttpApiServer(s grpc.ServiceRegistrar, srv HttpApiServer) {
	s.RegisterService(&HttpApi_ServiceDesc, srv)
}

func _HttpApi_InitClientEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpApiServer).InitClientEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HttpApi/InitClientEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpApiServer).InitClientEndpoint(ctx, req.(*InitClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpApi_InitServerEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpApiServer).InitServerEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HttpApi/InitServerEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpApiServer).InitServerEndpoint(ctx, req.(*InitServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpApi_ClientSendAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientSendActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpApiServer).ClientSendAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HttpApi/ClientSendAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpApiServer).ClientSendAction(ctx, req.(*ClientSendActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpApi_ClientReceiveAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientReceiveActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpApiServer).ClientReceiveAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HttpApi/ClientReceiveAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpApiServer).ClientReceiveAction(ctx, req.(*ClientReceiveActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpApi_ServerSendAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerSendActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpApiServer).ServerSendAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HttpApi/ServerSendAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpApiServer).ServerSendAction(ctx, req.(*ServerSendActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpApi_ServerReceiveAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerReceiveActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpApiServer).ServerReceiveAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HttpApi/ServerReceiveAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpApiServer).ServerReceiveAction(ctx, req.(*ServerReceiveActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HttpApi_ServiceDesc is the grpc.ServiceDesc for HttpApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HttpApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HttpApi",
	HandlerType: (*HttpApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitClientEndpoint",
			Handler:    _HttpApi_InitClientEndpoint_Handler,
		},
		{
			MethodName: "InitServerEndpoint",
			Handler:    _HttpApi_InitServerEndpoint_Handler,
		},
		{
			MethodName: "ClientSendAction",
			Handler:    _HttpApi_ClientSendAction_Handler,
		},
		{
			MethodName: "ClientReceiveAction",
			Handler:    _HttpApi_ClientReceiveAction_Handler,
		},
		{
			MethodName: "ServerSendAction",
			Handler:    _HttpApi_ServerSendAction_Handler,
		},
		{
			MethodName: "ServerReceiveAction",
			Handler:    _HttpApi_ServerReceiveAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application/api/http/http.proto",
}
