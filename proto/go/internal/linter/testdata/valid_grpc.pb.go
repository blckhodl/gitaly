// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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
	TestMethod(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error)
}

type interceptedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInterceptedServiceClient(cc grpc.ClientConnInterface) InterceptedServiceClient {
	return &interceptedServiceClient{cc}
}

func (c *interceptedServiceClient) TestMethod(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error) {
	out := new(ValidResponse)
	err := c.cc.Invoke(ctx, "/test.InterceptedService/TestMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InterceptedServiceServer is the server API for InterceptedService service.
// All implementations must embed UnimplementedInterceptedServiceServer
// for forward compatibility
type InterceptedServiceServer interface {
	TestMethod(context.Context, *ValidRequest) (*ValidResponse, error)
	mustEmbedUnimplementedInterceptedServiceServer()
}

// UnimplementedInterceptedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInterceptedServiceServer struct {
}

func (UnimplementedInterceptedServiceServer) TestMethod(context.Context, *ValidRequest) (*ValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMethod not implemented")
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

func _InterceptedService_TestMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterceptedServiceServer).TestMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.InterceptedService/TestMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterceptedServiceServer).TestMethod(ctx, req.(*ValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InterceptedService_ServiceDesc is the grpc.ServiceDesc for InterceptedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InterceptedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.InterceptedService",
	HandlerType: (*InterceptedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestMethod",
			Handler:    _InterceptedService_TestMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go/internal/linter/testdata/valid.proto",
}

// ValidServiceClient is the client API for ValidService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ValidServiceClient interface {
	TestMethod(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error)
	TestMethod2(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error)
	TestMethod3(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error)
	TestMethod5(ctx context.Context, in *ValidNestedRequest, opts ...grpc.CallOption) (*ValidResponse, error)
	TestMethod6(ctx context.Context, in *ValidNestedSharedRequest, opts ...grpc.CallOption) (*ValidResponse, error)
	TestMethod7(ctx context.Context, in *ValidInnerNestedRequest, opts ...grpc.CallOption) (*ValidResponse, error)
	TestMethod8(ctx context.Context, in *ValidStorageRequest, opts ...grpc.CallOption) (*ValidResponse, error)
	TestMethod9(ctx context.Context, in *ValidStorageNestedRequest, opts ...grpc.CallOption) (*ValidResponse, error)
}

type validServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewValidServiceClient(cc grpc.ClientConnInterface) ValidServiceClient {
	return &validServiceClient{cc}
}

func (c *validServiceClient) TestMethod(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error) {
	out := new(ValidResponse)
	err := c.cc.Invoke(ctx, "/test.ValidService/TestMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validServiceClient) TestMethod2(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error) {
	out := new(ValidResponse)
	err := c.cc.Invoke(ctx, "/test.ValidService/TestMethod2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validServiceClient) TestMethod3(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error) {
	out := new(ValidResponse)
	err := c.cc.Invoke(ctx, "/test.ValidService/TestMethod3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validServiceClient) TestMethod5(ctx context.Context, in *ValidNestedRequest, opts ...grpc.CallOption) (*ValidResponse, error) {
	out := new(ValidResponse)
	err := c.cc.Invoke(ctx, "/test.ValidService/TestMethod5", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validServiceClient) TestMethod6(ctx context.Context, in *ValidNestedSharedRequest, opts ...grpc.CallOption) (*ValidResponse, error) {
	out := new(ValidResponse)
	err := c.cc.Invoke(ctx, "/test.ValidService/TestMethod6", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validServiceClient) TestMethod7(ctx context.Context, in *ValidInnerNestedRequest, opts ...grpc.CallOption) (*ValidResponse, error) {
	out := new(ValidResponse)
	err := c.cc.Invoke(ctx, "/test.ValidService/TestMethod7", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validServiceClient) TestMethod8(ctx context.Context, in *ValidStorageRequest, opts ...grpc.CallOption) (*ValidResponse, error) {
	out := new(ValidResponse)
	err := c.cc.Invoke(ctx, "/test.ValidService/TestMethod8", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validServiceClient) TestMethod9(ctx context.Context, in *ValidStorageNestedRequest, opts ...grpc.CallOption) (*ValidResponse, error) {
	out := new(ValidResponse)
	err := c.cc.Invoke(ctx, "/test.ValidService/TestMethod9", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidServiceServer is the server API for ValidService service.
// All implementations must embed UnimplementedValidServiceServer
// for forward compatibility
type ValidServiceServer interface {
	TestMethod(context.Context, *ValidRequest) (*ValidResponse, error)
	TestMethod2(context.Context, *ValidRequest) (*ValidResponse, error)
	TestMethod3(context.Context, *ValidRequest) (*ValidResponse, error)
	TestMethod5(context.Context, *ValidNestedRequest) (*ValidResponse, error)
	TestMethod6(context.Context, *ValidNestedSharedRequest) (*ValidResponse, error)
	TestMethod7(context.Context, *ValidInnerNestedRequest) (*ValidResponse, error)
	TestMethod8(context.Context, *ValidStorageRequest) (*ValidResponse, error)
	TestMethod9(context.Context, *ValidStorageNestedRequest) (*ValidResponse, error)
	mustEmbedUnimplementedValidServiceServer()
}

// UnimplementedValidServiceServer must be embedded to have forward compatible implementations.
type UnimplementedValidServiceServer struct {
}

func (UnimplementedValidServiceServer) TestMethod(context.Context, *ValidRequest) (*ValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMethod not implemented")
}
func (UnimplementedValidServiceServer) TestMethod2(context.Context, *ValidRequest) (*ValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMethod2 not implemented")
}
func (UnimplementedValidServiceServer) TestMethod3(context.Context, *ValidRequest) (*ValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMethod3 not implemented")
}
func (UnimplementedValidServiceServer) TestMethod5(context.Context, *ValidNestedRequest) (*ValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMethod5 not implemented")
}
func (UnimplementedValidServiceServer) TestMethod6(context.Context, *ValidNestedSharedRequest) (*ValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMethod6 not implemented")
}
func (UnimplementedValidServiceServer) TestMethod7(context.Context, *ValidInnerNestedRequest) (*ValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMethod7 not implemented")
}
func (UnimplementedValidServiceServer) TestMethod8(context.Context, *ValidStorageRequest) (*ValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMethod8 not implemented")
}
func (UnimplementedValidServiceServer) TestMethod9(context.Context, *ValidStorageNestedRequest) (*ValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMethod9 not implemented")
}
func (UnimplementedValidServiceServer) mustEmbedUnimplementedValidServiceServer() {}

// UnsafeValidServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ValidServiceServer will
// result in compilation errors.
type UnsafeValidServiceServer interface {
	mustEmbedUnimplementedValidServiceServer()
}

func RegisterValidServiceServer(s grpc.ServiceRegistrar, srv ValidServiceServer) {
	s.RegisterService(&ValidService_ServiceDesc, srv)
}

func _ValidService_TestMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidServiceServer).TestMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.ValidService/TestMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidServiceServer).TestMethod(ctx, req.(*ValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidService_TestMethod2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidServiceServer).TestMethod2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.ValidService/TestMethod2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidServiceServer).TestMethod2(ctx, req.(*ValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidService_TestMethod3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidServiceServer).TestMethod3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.ValidService/TestMethod3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidServiceServer).TestMethod3(ctx, req.(*ValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidService_TestMethod5_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidNestedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidServiceServer).TestMethod5(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.ValidService/TestMethod5",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidServiceServer).TestMethod5(ctx, req.(*ValidNestedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidService_TestMethod6_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidNestedSharedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidServiceServer).TestMethod6(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.ValidService/TestMethod6",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidServiceServer).TestMethod6(ctx, req.(*ValidNestedSharedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidService_TestMethod7_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidInnerNestedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidServiceServer).TestMethod7(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.ValidService/TestMethod7",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidServiceServer).TestMethod7(ctx, req.(*ValidInnerNestedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidService_TestMethod8_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidServiceServer).TestMethod8(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.ValidService/TestMethod8",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidServiceServer).TestMethod8(ctx, req.(*ValidStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidService_TestMethod9_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidStorageNestedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidServiceServer).TestMethod9(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.ValidService/TestMethod9",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidServiceServer).TestMethod9(ctx, req.(*ValidStorageNestedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ValidService_ServiceDesc is the grpc.ServiceDesc for ValidService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ValidService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.ValidService",
	HandlerType: (*ValidServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestMethod",
			Handler:    _ValidService_TestMethod_Handler,
		},
		{
			MethodName: "TestMethod2",
			Handler:    _ValidService_TestMethod2_Handler,
		},
		{
			MethodName: "TestMethod3",
			Handler:    _ValidService_TestMethod3_Handler,
		},
		{
			MethodName: "TestMethod5",
			Handler:    _ValidService_TestMethod5_Handler,
		},
		{
			MethodName: "TestMethod6",
			Handler:    _ValidService_TestMethod6_Handler,
		},
		{
			MethodName: "TestMethod7",
			Handler:    _ValidService_TestMethod7_Handler,
		},
		{
			MethodName: "TestMethod8",
			Handler:    _ValidService_TestMethod8_Handler,
		},
		{
			MethodName: "TestMethod9",
			Handler:    _ValidService_TestMethod9_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go/internal/linter/testdata/valid.proto",
}