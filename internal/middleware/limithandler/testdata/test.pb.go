// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: middleware/limithandler/testdata/test.proto

package testdata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UnaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnaryRequest) Reset() {
	*x = UnaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryRequest) ProtoMessage() {}

func (x *UnaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryRequest.ProtoReflect.Descriptor instead.
func (*UnaryRequest) Descriptor() ([]byte, []int) {
	return file_middleware_limithandler_testdata_test_proto_rawDescGZIP(), []int{0}
}

type UnaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *UnaryResponse) Reset() {
	*x = UnaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryResponse) ProtoMessage() {}

func (x *UnaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryResponse.ProtoReflect.Descriptor instead.
func (*UnaryResponse) Descriptor() ([]byte, []int) {
	return file_middleware_limithandler_testdata_test_proto_rawDescGZIP(), []int{1}
}

func (x *UnaryResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type StreamInputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamInputRequest) Reset() {
	*x = StreamInputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamInputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamInputRequest) ProtoMessage() {}

func (x *StreamInputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamInputRequest.ProtoReflect.Descriptor instead.
func (*StreamInputRequest) Descriptor() ([]byte, []int) {
	return file_middleware_limithandler_testdata_test_proto_rawDescGZIP(), []int{2}
}

type StreamInputResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *StreamInputResponse) Reset() {
	*x = StreamInputResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamInputResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamInputResponse) ProtoMessage() {}

func (x *StreamInputResponse) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamInputResponse.ProtoReflect.Descriptor instead.
func (*StreamInputResponse) Descriptor() ([]byte, []int) {
	return file_middleware_limithandler_testdata_test_proto_rawDescGZIP(), []int{3}
}

func (x *StreamInputResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type StreamOutputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamOutputRequest) Reset() {
	*x = StreamOutputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamOutputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamOutputRequest) ProtoMessage() {}

func (x *StreamOutputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamOutputRequest.ProtoReflect.Descriptor instead.
func (*StreamOutputRequest) Descriptor() ([]byte, []int) {
	return file_middleware_limithandler_testdata_test_proto_rawDescGZIP(), []int{4}
}

type StreamOutputResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *StreamOutputResponse) Reset() {
	*x = StreamOutputResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamOutputResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamOutputResponse) ProtoMessage() {}

func (x *StreamOutputResponse) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamOutputResponse.ProtoReflect.Descriptor instead.
func (*StreamOutputResponse) Descriptor() ([]byte, []int) {
	return file_middleware_limithandler_testdata_test_proto_rawDescGZIP(), []int{5}
}

func (x *StreamOutputResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type BidirectionalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BidirectionalRequest) Reset() {
	*x = BidirectionalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidirectionalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidirectionalRequest) ProtoMessage() {}

func (x *BidirectionalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidirectionalRequest.ProtoReflect.Descriptor instead.
func (*BidirectionalRequest) Descriptor() ([]byte, []int) {
	return file_middleware_limithandler_testdata_test_proto_rawDescGZIP(), []int{6}
}

type BidirectionalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *BidirectionalResponse) Reset() {
	*x = BidirectionalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidirectionalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidirectionalResponse) ProtoMessage() {}

func (x *BidirectionalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_limithandler_testdata_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidirectionalResponse.ProtoReflect.Descriptor instead.
func (*BidirectionalResponse) Descriptor() ([]byte, []int) {
	return file_middleware_limithandler_testdata_test_proto_rawDescGZIP(), []int{7}
}

func (x *BidirectionalResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_middleware_limithandler_testdata_test_proto protoreflect.FileDescriptor

var file_middleware_limithandler_testdata_test_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x22, 0x0e, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x1f, 0x0a, 0x0d, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f,
	0x6b, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x15,
	0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x16, 0x0a,
	0x14, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0x85,
	0x03, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x05, 0x55, 0x6e, 0x61, 0x72, 0x79,
	0x12, 0x1f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x25, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x63, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x26, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x68, 0x0a, 0x0d,
	0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x27, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2e, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x79, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_middleware_limithandler_testdata_test_proto_rawDescOnce sync.Once
	file_middleware_limithandler_testdata_test_proto_rawDescData = file_middleware_limithandler_testdata_test_proto_rawDesc
)

func file_middleware_limithandler_testdata_test_proto_rawDescGZIP() []byte {
	file_middleware_limithandler_testdata_test_proto_rawDescOnce.Do(func() {
		file_middleware_limithandler_testdata_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_middleware_limithandler_testdata_test_proto_rawDescData)
	})
	return file_middleware_limithandler_testdata_test_proto_rawDescData
}

var file_middleware_limithandler_testdata_test_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_middleware_limithandler_testdata_test_proto_goTypes = []interface{}{
	(*UnaryRequest)(nil),          // 0: test.limithandler.UnaryRequest
	(*UnaryResponse)(nil),         // 1: test.limithandler.UnaryResponse
	(*StreamInputRequest)(nil),    // 2: test.limithandler.StreamInputRequest
	(*StreamInputResponse)(nil),   // 3: test.limithandler.StreamInputResponse
	(*StreamOutputRequest)(nil),   // 4: test.limithandler.StreamOutputRequest
	(*StreamOutputResponse)(nil),  // 5: test.limithandler.StreamOutputResponse
	(*BidirectionalRequest)(nil),  // 6: test.limithandler.BidirectionalRequest
	(*BidirectionalResponse)(nil), // 7: test.limithandler.BidirectionalResponse
}
var file_middleware_limithandler_testdata_test_proto_depIdxs = []int32{
	0, // 0: test.limithandler.Test.Unary:input_type -> test.limithandler.UnaryRequest
	2, // 1: test.limithandler.Test.StreamInput:input_type -> test.limithandler.StreamInputRequest
	4, // 2: test.limithandler.Test.StreamOutput:input_type -> test.limithandler.StreamOutputRequest
	6, // 3: test.limithandler.Test.Bidirectional:input_type -> test.limithandler.BidirectionalRequest
	1, // 4: test.limithandler.Test.Unary:output_type -> test.limithandler.UnaryResponse
	3, // 5: test.limithandler.Test.StreamInput:output_type -> test.limithandler.StreamInputResponse
	5, // 6: test.limithandler.Test.StreamOutput:output_type -> test.limithandler.StreamOutputResponse
	7, // 7: test.limithandler.Test.Bidirectional:output_type -> test.limithandler.BidirectionalResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_middleware_limithandler_testdata_test_proto_init() }
func file_middleware_limithandler_testdata_test_proto_init() {
	if File_middleware_limithandler_testdata_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_middleware_limithandler_testdata_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_middleware_limithandler_testdata_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_middleware_limithandler_testdata_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamInputRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_middleware_limithandler_testdata_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamInputResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_middleware_limithandler_testdata_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamOutputRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_middleware_limithandler_testdata_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamOutputResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_middleware_limithandler_testdata_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidirectionalRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_middleware_limithandler_testdata_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidirectionalResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_middleware_limithandler_testdata_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_middleware_limithandler_testdata_test_proto_goTypes,
		DependencyIndexes: file_middleware_limithandler_testdata_test_proto_depIdxs,
		MessageInfos:      file_middleware_limithandler_testdata_test_proto_msgTypes,
	}.Build()
	File_middleware_limithandler_testdata_test_proto = out.File
	file_middleware_limithandler_testdata_test_proto_rawDesc = nil
	file_middleware_limithandler_testdata_test_proto_goTypes = nil
	file_middleware_limithandler_testdata_test_proto_depIdxs = nil
}
