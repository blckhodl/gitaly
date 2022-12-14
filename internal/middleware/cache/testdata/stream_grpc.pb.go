// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: middleware/cache/testdata/stream.proto

package testdata

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

// InterceptedServiceClient is the client API for InterceptedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InterceptedServiceClient interface {
	IgnoredMethod(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type interceptedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInterceptedServiceClient(cc grpc.ClientConnInterface) InterceptedServiceClient {
	return &interceptedServiceClient{cc}
}

func (c *interceptedServiceClient) IgnoredMethod(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/testdata.InterceptedService/IgnoredMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InterceptedServiceServer is the server API for InterceptedService service.
// All implementations must embed UnimplementedInterceptedServiceServer
// for forward compatibility
type InterceptedServiceServer interface {
	IgnoredMethod(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedInterceptedServiceServer()
}

// UnimplementedInterceptedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInterceptedServiceServer struct {
}

func (UnimplementedInterceptedServiceServer) IgnoredMethod(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IgnoredMethod not implemented")
}
func (UnimplementedInterceptedServiceServer) mustEmbedUnimplementedInterceptedServiceServer() {}

// UnsafeInterceptedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InterceptedServiceServer will
// result in compilation errors.
type UnsafeInterceptedServiceServer interface {
	mustEmbedUnimplementedInterceptedServiceServer()
}

func RegisterInterceptedServiceServer(s grpc.ServiceRegistrar, srv InterceptedServiceServer) {
	s.RegisterService(&InterceptedService_ServiceDesc, srv)
}

func _InterceptedService_IgnoredMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterceptedServiceServer).IgnoredMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.InterceptedService/IgnoredMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterceptedServiceServer).IgnoredMethod(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// InterceptedService_ServiceDesc is the grpc.ServiceDesc for InterceptedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InterceptedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testdata.InterceptedService",
	HandlerType: (*InterceptedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IgnoredMethod",
			Handler:    _InterceptedService_IgnoredMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "middleware/cache/testdata/stream.proto",
}

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	ClientStreamRepoMutator(ctx context.Context, in *Request, opts ...grpc.CallOption) (TestService_ClientStreamRepoMutatorClient, error)
	ClientStreamRepoAccessor(ctx context.Context, in *Request, opts ...grpc.CallOption) (TestService_ClientStreamRepoAccessorClient, error)
	ClientStreamRepoMaintainer(ctx context.Context, in *Request, opts ...grpc.CallOption) (TestService_ClientStreamRepoMaintainerClient, error)
	ClientUnaryRepoMutator(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	ClientUnaryRepoAccessor(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	ClientUnaryRepoMaintainer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) ClientStreamRepoMutator(ctx context.Context, in *Request, opts ...grpc.CallOption) (TestService_ClientStreamRepoMutatorClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[0], "/testdata.TestService/ClientStreamRepoMutator", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceClientStreamRepoMutatorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_ClientStreamRepoMutatorClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type testServiceClientStreamRepoMutatorClient struct {
	grpc.ClientStream
}

func (x *testServiceClientStreamRepoMutatorClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) ClientStreamRepoAccessor(ctx context.Context, in *Request, opts ...grpc.CallOption) (TestService_ClientStreamRepoAccessorClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[1], "/testdata.TestService/ClientStreamRepoAccessor", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceClientStreamRepoAccessorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_ClientStreamRepoAccessorClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type testServiceClientStreamRepoAccessorClient struct {
	grpc.ClientStream
}

func (x *testServiceClientStreamRepoAccessorClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) ClientStreamRepoMaintainer(ctx context.Context, in *Request, opts ...grpc.CallOption) (TestService_ClientStreamRepoMaintainerClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[2], "/testdata.TestService/ClientStreamRepoMaintainer", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceClientStreamRepoMaintainerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_ClientStreamRepoMaintainerClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type testServiceClientStreamRepoMaintainerClient struct {
	grpc.ClientStream
}

func (x *testServiceClientStreamRepoMaintainerClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) ClientUnaryRepoMutator(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/testdata.TestService/ClientUnaryRepoMutator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) ClientUnaryRepoAccessor(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/testdata.TestService/ClientUnaryRepoAccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) ClientUnaryRepoMaintainer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/testdata.TestService/ClientUnaryRepoMaintainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	ClientStreamRepoMutator(*Request, TestService_ClientStreamRepoMutatorServer) error
	ClientStreamRepoAccessor(*Request, TestService_ClientStreamRepoAccessorServer) error
	ClientStreamRepoMaintainer(*Request, TestService_ClientStreamRepoMaintainerServer) error
	ClientUnaryRepoMutator(context.Context, *Request) (*Response, error)
	ClientUnaryRepoAccessor(context.Context, *Request) (*Response, error)
	ClientUnaryRepoMaintainer(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) ClientStreamRepoMutator(*Request, TestService_ClientStreamRepoMutatorServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamRepoMutator not implemented")
}
func (UnimplementedTestServiceServer) ClientStreamRepoAccessor(*Request, TestService_ClientStreamRepoAccessorServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamRepoAccessor not implemented")
}
func (UnimplementedTestServiceServer) ClientStreamRepoMaintainer(*Request, TestService_ClientStreamRepoMaintainerServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamRepoMaintainer not implemented")
}
func (UnimplementedTestServiceServer) ClientUnaryRepoMutator(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientUnaryRepoMutator not implemented")
}
func (UnimplementedTestServiceServer) ClientUnaryRepoAccessor(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientUnaryRepoAccessor not implemented")
}
func (UnimplementedTestServiceServer) ClientUnaryRepoMaintainer(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientUnaryRepoMaintainer not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_ClientStreamRepoMutator_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).ClientStreamRepoMutator(m, &testServiceClientStreamRepoMutatorServer{stream})
}

type TestService_ClientStreamRepoMutatorServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type testServiceClientStreamRepoMutatorServer struct {
	grpc.ServerStream
}

func (x *testServiceClientStreamRepoMutatorServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_ClientStreamRepoAccessor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).ClientStreamRepoAccessor(m, &testServiceClientStreamRepoAccessorServer{stream})
}

type TestService_ClientStreamRepoAccessorServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type testServiceClientStreamRepoAccessorServer struct {
	grpc.ServerStream
}

func (x *testServiceClientStreamRepoAccessorServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_ClientStreamRepoMaintainer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).ClientStreamRepoMaintainer(m, &testServiceClientStreamRepoMaintainerServer{stream})
}

type TestService_ClientStreamRepoMaintainerServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type testServiceClientStreamRepoMaintainerServer struct {
	grpc.ServerStream
}

func (x *testServiceClientStreamRepoMaintainerServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_ClientUnaryRepoMutator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).ClientUnaryRepoMutator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.TestService/ClientUnaryRepoMutator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).ClientUnaryRepoMutator(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_ClientUnaryRepoAccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).ClientUnaryRepoAccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.TestService/ClientUnaryRepoAccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).ClientUnaryRepoAccessor(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_ClientUnaryRepoMaintainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).ClientUnaryRepoMaintainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.TestService/ClientUnaryRepoMaintainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).ClientUnaryRepoMaintainer(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testdata.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClientUnaryRepoMutator",
			Handler:    _TestService_ClientUnaryRepoMutator_Handler,
		},
		{
			MethodName: "ClientUnaryRepoAccessor",
			Handler:    _TestService_ClientUnaryRepoAccessor_Handler,
		},
		{
			MethodName: "ClientUnaryRepoMaintainer",
			Handler:    _TestService_ClientUnaryRepoMaintainer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStreamRepoMutator",
			Handler:       _TestService_ClientStreamRepoMutator_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamRepoAccessor",
			Handler:       _TestService_ClientStreamRepoAccessor_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamRepoMaintainer",
			Handler:       _TestService_ClientStreamRepoMaintainer_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "middleware/cache/testdata/stream.proto",
}
