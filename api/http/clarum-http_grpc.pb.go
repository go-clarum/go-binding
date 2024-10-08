// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: api/http/clarum-http.proto

package http

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

// HttpServiceClient is the client API for HttpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HttpServiceClient interface {
	InitClientEndpoint(ctx context.Context, in *InitClientRequest, opts ...grpc.CallOption) (*InitClientResponse, error)
	InitServerEndpoint(ctx context.Context, in *InitServerRequest, opts ...grpc.CallOption) (*InitServerResponse, error)
	ClientSendAction(ctx context.Context, in *ClientSendActionRequest, opts ...grpc.CallOption) (*ClientSendActionResponse, error)
	ClientReceiveAction(ctx context.Context, in *ClientReceiveActionRequest, opts ...grpc.CallOption) (*ClientReceiveActionResponse, error)
	ServerSendAction(ctx context.Context, in *ServerSendActionRequest, opts ...grpc.CallOption) (*ServerSendActionResponse, error)
	ServerReceiveAction(ctx context.Context, in *ServerReceiveActionRequest, opts ...grpc.CallOption) (*ServerReceiveActionResponse, error)
}

type httpServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHttpServiceClient(cc grpc.ClientConnInterface) HttpServiceClient {
	return &httpServiceClient{cc}
}

func (c *httpServiceClient) InitClientEndpoint(ctx context.Context, in *InitClientRequest, opts ...grpc.CallOption) (*InitClientResponse, error) {
	out := new(InitClientResponse)
	err := c.cc.Invoke(ctx, "/HttpService/InitClientEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpServiceClient) InitServerEndpoint(ctx context.Context, in *InitServerRequest, opts ...grpc.CallOption) (*InitServerResponse, error) {
	out := new(InitServerResponse)
	err := c.cc.Invoke(ctx, "/HttpService/InitServerEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpServiceClient) ClientSendAction(ctx context.Context, in *ClientSendActionRequest, opts ...grpc.CallOption) (*ClientSendActionResponse, error) {
	out := new(ClientSendActionResponse)
	err := c.cc.Invoke(ctx, "/HttpService/ClientSendAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpServiceClient) ClientReceiveAction(ctx context.Context, in *ClientReceiveActionRequest, opts ...grpc.CallOption) (*ClientReceiveActionResponse, error) {
	out := new(ClientReceiveActionResponse)
	err := c.cc.Invoke(ctx, "/HttpService/ClientReceiveAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpServiceClient) ServerSendAction(ctx context.Context, in *ServerSendActionRequest, opts ...grpc.CallOption) (*ServerSendActionResponse, error) {
	out := new(ServerSendActionResponse)
	err := c.cc.Invoke(ctx, "/HttpService/ServerSendAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpServiceClient) ServerReceiveAction(ctx context.Context, in *ServerReceiveActionRequest, opts ...grpc.CallOption) (*ServerReceiveActionResponse, error) {
	out := new(ServerReceiveActionResponse)
	err := c.cc.Invoke(ctx, "/HttpService/ServerReceiveAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HttpServiceServer is the server API for HttpService service.
// All implementations must embed UnimplementedHttpServiceServer
// for forward compatibility
type HttpServiceServer interface {
	InitClientEndpoint(context.Context, *InitClientRequest) (*InitClientResponse, error)
	InitServerEndpoint(context.Context, *InitServerRequest) (*InitServerResponse, error)
	ClientSendAction(context.Context, *ClientSendActionRequest) (*ClientSendActionResponse, error)
	ClientReceiveAction(context.Context, *ClientReceiveActionRequest) (*ClientReceiveActionResponse, error)
	ServerSendAction(context.Context, *ServerSendActionRequest) (*ServerSendActionResponse, error)
	ServerReceiveAction(context.Context, *ServerReceiveActionRequest) (*ServerReceiveActionResponse, error)
	mustEmbedUnimplementedHttpServiceServer()
}

// UnimplementedHttpServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHttpServiceServer struct {
}

func (UnimplementedHttpServiceServer) InitClientEndpoint(context.Context, *InitClientRequest) (*InitClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitClientEndpoint not implemented")
}
func (UnimplementedHttpServiceServer) InitServerEndpoint(context.Context, *InitServerRequest) (*InitServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitServerEndpoint not implemented")
}
func (UnimplementedHttpServiceServer) ClientSendAction(context.Context, *ClientSendActionRequest) (*ClientSendActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientSendAction not implemented")
}
func (UnimplementedHttpServiceServer) ClientReceiveAction(context.Context, *ClientReceiveActionRequest) (*ClientReceiveActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientReceiveAction not implemented")
}
func (UnimplementedHttpServiceServer) ServerSendAction(context.Context, *ServerSendActionRequest) (*ServerSendActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerSendAction not implemented")
}
func (UnimplementedHttpServiceServer) ServerReceiveAction(context.Context, *ServerReceiveActionRequest) (*ServerReceiveActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerReceiveAction not implemented")
}
func (UnimplementedHttpServiceServer) mustEmbedUnimplementedHttpServiceServer() {}

// UnsafeHttpServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HttpServiceServer will
// result in compilation errors.
type UnsafeHttpServiceServer interface {
	mustEmbedUnimplementedHttpServiceServer()
}

func RegisterHttpServiceServer(s grpc.ServiceRegistrar, srv HttpServiceServer) {
	s.RegisterService(&HttpService_ServiceDesc, srv)
}

func _HttpService_InitClientEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpServiceServer).InitClientEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HttpService/InitClientEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpServiceServer).InitClientEndpoint(ctx, req.(*InitClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpService_InitServerEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpServiceServer).InitServerEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HttpService/InitServerEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpServiceServer).InitServerEndpoint(ctx, req.(*InitServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpService_ClientSendAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientSendActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpServiceServer).ClientSendAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HttpService/ClientSendAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpServiceServer).ClientSendAction(ctx, req.(*ClientSendActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpService_ClientReceiveAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientReceiveActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpServiceServer).ClientReceiveAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HttpService/ClientReceiveAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpServiceServer).ClientReceiveAction(ctx, req.(*ClientReceiveActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpService_ServerSendAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerSendActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpServiceServer).ServerSendAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HttpService/ServerSendAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpServiceServer).ServerSendAction(ctx, req.(*ServerSendActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpService_ServerReceiveAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerReceiveActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpServiceServer).ServerReceiveAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HttpService/ServerReceiveAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpServiceServer).ServerReceiveAction(ctx, req.(*ServerReceiveActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HttpService_ServiceDesc is the grpc.ServiceDesc for HttpService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HttpService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HttpService",
	HandlerType: (*HttpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitClientEndpoint",
			Handler:    _HttpService_InitClientEndpoint_Handler,
		},
		{
			MethodName: "InitServerEndpoint",
			Handler:    _HttpService_InitServerEndpoint_Handler,
		},
		{
			MethodName: "ClientSendAction",
			Handler:    _HttpService_ClientSendAction_Handler,
		},
		{
			MethodName: "ClientReceiveAction",
			Handler:    _HttpService_ClientReceiveAction_Handler,
		},
		{
			MethodName: "ServerSendAction",
			Handler:    _HttpService_ServerSendAction_Handler,
		},
		{
			MethodName: "ServerReceiveAction",
			Handler:    _HttpService_ServerReceiveAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/http/clarum-http.proto",
}
