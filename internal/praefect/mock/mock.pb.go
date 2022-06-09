//
//This file is a mock gRPC service used for validating the various types of
//gRPC methods that Praefect is expected to reverse proxy. It is intended to keep
//tests simple and keep Praefect decoupled from specific gRPC services.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: praefect/mock/mock.proto

package mock

import (
	gitalypb "gitlab.com/gitlab-org/gitaly/v15/proto/go/gitalypb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo *gitalypb.Repository `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
}

func (x *RepoRequest) Reset() {
	*x = RepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_praefect_mock_mock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoRequest) ProtoMessage() {}

func (x *RepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_praefect_mock_mock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoRequest.ProtoReflect.Descriptor instead.
func (*RepoRequest) Descriptor() ([]byte, []int) {
	return file_praefect_mock_mock_proto_rawDescGZIP(), []int{0}
}

func (x *RepoRequest) GetRepo() *gitalypb.Repository {
	if x != nil {
		return x.Repo
	}
	return nil
}

var File_praefect_mock_mock_proto protoreflect.FileDescriptor

var file_praefect_mock_mock_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x61, 0x65, 0x66, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x2f,
	0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x6f, 0x63, 0x6b,
	0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x6c, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x79, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x04, 0x98, 0xc6, 0x2c, 0x01, 0x52, 0x04,
	0x72, 0x65, 0x70, 0x6f, 0x32, 0xe9, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x11, 0x2e, 0x6d, 0x6f,
	0x63, 0x6b, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x06, 0xfa, 0x97, 0x28, 0x02, 0x08, 0x02, 0x12, 0x45,
	0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x6e, 0x61,
	0x72, 0x79, 0x12, 0x11, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x06, 0xfa,
	0x97, 0x28, 0x02, 0x08, 0x01, 0x12, 0x49, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6f, 0x4d, 0x61, 0x69,
	0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x11, 0x2e,
	0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x06, 0xfa, 0x97, 0x28, 0x02, 0x08, 0x03,
	0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x79,
	0x2f, 0x76, 0x31, 0x35, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x61, 0x65, 0x66, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_praefect_mock_mock_proto_rawDescOnce sync.Once
	file_praefect_mock_mock_proto_rawDescData = file_praefect_mock_mock_proto_rawDesc
)

func file_praefect_mock_mock_proto_rawDescGZIP() []byte {
	file_praefect_mock_mock_proto_rawDescOnce.Do(func() {
		file_praefect_mock_mock_proto_rawDescData = protoimpl.X.CompressGZIP(file_praefect_mock_mock_proto_rawDescData)
	})
	return file_praefect_mock_mock_proto_rawDescData
}

var file_praefect_mock_mock_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_praefect_mock_mock_proto_goTypes = []interface{}{
	(*RepoRequest)(nil),         // 0: mock.RepoRequest
	(*gitalypb.Repository)(nil), // 1: gitaly.Repository
	(*emptypb.Empty)(nil),       // 2: google.protobuf.Empty
}
var file_praefect_mock_mock_proto_depIdxs = []int32{
	1, // 0: mock.RepoRequest.repo:type_name -> gitaly.Repository
	0, // 1: mock.SimpleService.RepoAccessorUnary:input_type -> mock.RepoRequest
	0, // 2: mock.SimpleService.RepoMutatorUnary:input_type -> mock.RepoRequest
	0, // 3: mock.SimpleService.RepoMaintenanceUnary:input_type -> mock.RepoRequest
	2, // 4: mock.SimpleService.RepoAccessorUnary:output_type -> google.protobuf.Empty
	2, // 5: mock.SimpleService.RepoMutatorUnary:output_type -> google.protobuf.Empty
	2, // 6: mock.SimpleService.RepoMaintenanceUnary:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_praefect_mock_mock_proto_init() }
func file_praefect_mock_mock_proto_init() {
	if File_praefect_mock_mock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_praefect_mock_mock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoRequest); i {
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
			RawDescriptor: file_praefect_mock_mock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_praefect_mock_mock_proto_goTypes,
		DependencyIndexes: file_praefect_mock_mock_proto_depIdxs,
		MessageInfos:      file_praefect_mock_mock_proto_msgTypes,
	}.Build()
	File_praefect_mock_mock_proto = out.File
	file_praefect_mock_mock_proto_rawDesc = nil
	file_praefect_mock_mock_proto_goTypes = nil
	file_praefect_mock_mock_proto_depIdxs = nil
}
