// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssh.proto

package gitalypb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SSHUploadPackRequest struct {
	// 'repository' must be present in the first message.
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// A chunk of raw data to be copied to 'git upload-pack' standard input
	Stdin []byte `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
	// Parameters to use with git -c (key=value pairs)
	GitConfigOptions []string `protobuf:"bytes,4,rep,name=git_config_options,json=gitConfigOptions" json:"git_config_options,omitempty"`
	// Git protocol version
	GitProtocol string `protobuf:"bytes,5,opt,name=git_protocol,json=gitProtocol" json:"git_protocol,omitempty"`
}

func (m *SSHUploadPackRequest) Reset()                    { *m = SSHUploadPackRequest{} }
func (m *SSHUploadPackRequest) String() string            { return proto.CompactTextString(m) }
func (*SSHUploadPackRequest) ProtoMessage()               {}
func (*SSHUploadPackRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *SSHUploadPackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *SSHUploadPackRequest) GetStdin() []byte {
	if m != nil {
		return m.Stdin
	}
	return nil
}

func (m *SSHUploadPackRequest) GetGitConfigOptions() []string {
	if m != nil {
		return m.GitConfigOptions
	}
	return nil
}

func (m *SSHUploadPackRequest) GetGitProtocol() string {
	if m != nil {
		return m.GitProtocol
	}
	return ""
}

type SSHUploadPackResponse struct {
	// A chunk of raw data from 'git upload-pack' standard output
	Stdout []byte `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	// A chunk of raw data from 'git upload-pack' standard error
	Stderr []byte `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// This field may be nil. This is intentional: only when the remote
	// command has finished can we return its exit status.
	ExitStatus *ExitStatus `protobuf:"bytes,3,opt,name=exit_status,json=exitStatus" json:"exit_status,omitempty"`
}

func (m *SSHUploadPackResponse) Reset()                    { *m = SSHUploadPackResponse{} }
func (m *SSHUploadPackResponse) String() string            { return proto.CompactTextString(m) }
func (*SSHUploadPackResponse) ProtoMessage()               {}
func (*SSHUploadPackResponse) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *SSHUploadPackResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *SSHUploadPackResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *SSHUploadPackResponse) GetExitStatus() *ExitStatus {
	if m != nil {
		return m.ExitStatus
	}
	return nil
}

type SSHReceivePackRequest struct {
	// 'repository' must be present in the first message.
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// A chunk of raw data to be copied to 'git upload-pack' standard input
	Stdin []byte `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
	// Contents of GL_ID, GL_REPOSITORY, and GL_USERNAME environment variables
	// for 'git receive-pack'
	GlId         string `protobuf:"bytes,3,opt,name=gl_id,json=glId" json:"gl_id,omitempty"`
	GlRepository string `protobuf:"bytes,4,opt,name=gl_repository,json=glRepository" json:"gl_repository,omitempty"`
	GlUsername   string `protobuf:"bytes,5,opt,name=gl_username,json=glUsername" json:"gl_username,omitempty"`
	// Git protocol version
	GitProtocol string `protobuf:"bytes,6,opt,name=git_protocol,json=gitProtocol" json:"git_protocol,omitempty"`
	// Parameters to use with git -c (key=value pairs)
	GitConfigOptions []string `protobuf:"bytes,7,rep,name=git_config_options,json=gitConfigOptions" json:"git_config_options,omitempty"`
}

func (m *SSHReceivePackRequest) Reset()                    { *m = SSHReceivePackRequest{} }
func (m *SSHReceivePackRequest) String() string            { return proto.CompactTextString(m) }
func (*SSHReceivePackRequest) ProtoMessage()               {}
func (*SSHReceivePackRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *SSHReceivePackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *SSHReceivePackRequest) GetStdin() []byte {
	if m != nil {
		return m.Stdin
	}
	return nil
}

func (m *SSHReceivePackRequest) GetGlId() string {
	if m != nil {
		return m.GlId
	}
	return ""
}

func (m *SSHReceivePackRequest) GetGlRepository() string {
	if m != nil {
		return m.GlRepository
	}
	return ""
}

func (m *SSHReceivePackRequest) GetGlUsername() string {
	if m != nil {
		return m.GlUsername
	}
	return ""
}

func (m *SSHReceivePackRequest) GetGitProtocol() string {
	if m != nil {
		return m.GitProtocol
	}
	return ""
}

func (m *SSHReceivePackRequest) GetGitConfigOptions() []string {
	if m != nil {
		return m.GitConfigOptions
	}
	return nil
}

type SSHReceivePackResponse struct {
	// A chunk of raw data from 'git receive-pack' standard output
	Stdout []byte `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	// A chunk of raw data from 'git receive-pack' standard error
	Stderr []byte `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// This field may be nil. This is intentional: only when the remote
	// command has finished can we return its exit status.
	ExitStatus *ExitStatus `protobuf:"bytes,3,opt,name=exit_status,json=exitStatus" json:"exit_status,omitempty"`
}

func (m *SSHReceivePackResponse) Reset()                    { *m = SSHReceivePackResponse{} }
func (m *SSHReceivePackResponse) String() string            { return proto.CompactTextString(m) }
func (*SSHReceivePackResponse) ProtoMessage()               {}
func (*SSHReceivePackResponse) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *SSHReceivePackResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *SSHReceivePackResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *SSHReceivePackResponse) GetExitStatus() *ExitStatus {
	if m != nil {
		return m.ExitStatus
	}
	return nil
}

type SSHUploadArchiveRequest struct {
	// 'repository' must be present in the first message.
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// A chunk of raw data to be copied to 'git upload-archive' standard input
	Stdin []byte `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
}

func (m *SSHUploadArchiveRequest) Reset()                    { *m = SSHUploadArchiveRequest{} }
func (m *SSHUploadArchiveRequest) String() string            { return proto.CompactTextString(m) }
func (*SSHUploadArchiveRequest) ProtoMessage()               {}
func (*SSHUploadArchiveRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

func (m *SSHUploadArchiveRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *SSHUploadArchiveRequest) GetStdin() []byte {
	if m != nil {
		return m.Stdin
	}
	return nil
}

type SSHUploadArchiveResponse struct {
	// A chunk of raw data from 'git upload-archive' standard output
	Stdout []byte `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	// A chunk of raw data from 'git upload-archive' standard error
	Stderr []byte `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// This value will only be set on the last message
	ExitStatus *ExitStatus `protobuf:"bytes,3,opt,name=exit_status,json=exitStatus" json:"exit_status,omitempty"`
}

func (m *SSHUploadArchiveResponse) Reset()                    { *m = SSHUploadArchiveResponse{} }
func (m *SSHUploadArchiveResponse) String() string            { return proto.CompactTextString(m) }
func (*SSHUploadArchiveResponse) ProtoMessage()               {}
func (*SSHUploadArchiveResponse) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

func (m *SSHUploadArchiveResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *SSHUploadArchiveResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *SSHUploadArchiveResponse) GetExitStatus() *ExitStatus {
	if m != nil {
		return m.ExitStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*SSHUploadPackRequest)(nil), "gitaly.SSHUploadPackRequest")
	proto.RegisterType((*SSHUploadPackResponse)(nil), "gitaly.SSHUploadPackResponse")
	proto.RegisterType((*SSHReceivePackRequest)(nil), "gitaly.SSHReceivePackRequest")
	proto.RegisterType((*SSHReceivePackResponse)(nil), "gitaly.SSHReceivePackResponse")
	proto.RegisterType((*SSHUploadArchiveRequest)(nil), "gitaly.SSHUploadArchiveRequest")
	proto.RegisterType((*SSHUploadArchiveResponse)(nil), "gitaly.SSHUploadArchiveResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SSHService service

type SSHServiceClient interface {
	// To forward 'git upload-pack' to Gitaly for SSH sessions
	SSHUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHUploadPackClient, error)
	// To forward 'git receive-pack' to Gitaly for SSH sessions
	SSHReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHReceivePackClient, error)
	// To forward 'git upload-archive' to Gitaly for SSH sessions
	SSHUploadArchive(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHUploadArchiveClient, error)
}

type sSHServiceClient struct {
	cc *grpc.ClientConn
}

func NewSSHServiceClient(cc *grpc.ClientConn) SSHServiceClient {
	return &sSHServiceClient{cc}
}

func (c *sSHServiceClient) SSHUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHUploadPackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SSHService_serviceDesc.Streams[0], c.cc, "/gitaly.SSHService/SSHUploadPack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHServiceSSHUploadPackClient{stream}
	return x, nil
}

type SSHService_SSHUploadPackClient interface {
	Send(*SSHUploadPackRequest) error
	Recv() (*SSHUploadPackResponse, error)
	grpc.ClientStream
}

type sSHServiceSSHUploadPackClient struct {
	grpc.ClientStream
}

func (x *sSHServiceSSHUploadPackClient) Send(m *SSHUploadPackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHServiceSSHUploadPackClient) Recv() (*SSHUploadPackResponse, error) {
	m := new(SSHUploadPackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sSHServiceClient) SSHReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHReceivePackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SSHService_serviceDesc.Streams[1], c.cc, "/gitaly.SSHService/SSHReceivePack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHServiceSSHReceivePackClient{stream}
	return x, nil
}

type SSHService_SSHReceivePackClient interface {
	Send(*SSHReceivePackRequest) error
	Recv() (*SSHReceivePackResponse, error)
	grpc.ClientStream
}

type sSHServiceSSHReceivePackClient struct {
	grpc.ClientStream
}

func (x *sSHServiceSSHReceivePackClient) Send(m *SSHReceivePackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHServiceSSHReceivePackClient) Recv() (*SSHReceivePackResponse, error) {
	m := new(SSHReceivePackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sSHServiceClient) SSHUploadArchive(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHUploadArchiveClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SSHService_serviceDesc.Streams[2], c.cc, "/gitaly.SSHService/SSHUploadArchive", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHServiceSSHUploadArchiveClient{stream}
	return x, nil
}

type SSHService_SSHUploadArchiveClient interface {
	Send(*SSHUploadArchiveRequest) error
	Recv() (*SSHUploadArchiveResponse, error)
	grpc.ClientStream
}

type sSHServiceSSHUploadArchiveClient struct {
	grpc.ClientStream
}

func (x *sSHServiceSSHUploadArchiveClient) Send(m *SSHUploadArchiveRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHServiceSSHUploadArchiveClient) Recv() (*SSHUploadArchiveResponse, error) {
	m := new(SSHUploadArchiveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SSHService service

type SSHServiceServer interface {
	// To forward 'git upload-pack' to Gitaly for SSH sessions
	SSHUploadPack(SSHService_SSHUploadPackServer) error
	// To forward 'git receive-pack' to Gitaly for SSH sessions
	SSHReceivePack(SSHService_SSHReceivePackServer) error
	// To forward 'git upload-archive' to Gitaly for SSH sessions
	SSHUploadArchive(SSHService_SSHUploadArchiveServer) error
}

func RegisterSSHServiceServer(s *grpc.Server, srv SSHServiceServer) {
	s.RegisterService(&_SSHService_serviceDesc, srv)
}

func _SSHService_SSHUploadPack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServiceServer).SSHUploadPack(&sSHServiceSSHUploadPackServer{stream})
}

type SSHService_SSHUploadPackServer interface {
	Send(*SSHUploadPackResponse) error
	Recv() (*SSHUploadPackRequest, error)
	grpc.ServerStream
}

type sSHServiceSSHUploadPackServer struct {
	grpc.ServerStream
}

func (x *sSHServiceSSHUploadPackServer) Send(m *SSHUploadPackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHServiceSSHUploadPackServer) Recv() (*SSHUploadPackRequest, error) {
	m := new(SSHUploadPackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SSHService_SSHReceivePack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServiceServer).SSHReceivePack(&sSHServiceSSHReceivePackServer{stream})
}

type SSHService_SSHReceivePackServer interface {
	Send(*SSHReceivePackResponse) error
	Recv() (*SSHReceivePackRequest, error)
	grpc.ServerStream
}

type sSHServiceSSHReceivePackServer struct {
	grpc.ServerStream
}

func (x *sSHServiceSSHReceivePackServer) Send(m *SSHReceivePackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHServiceSSHReceivePackServer) Recv() (*SSHReceivePackRequest, error) {
	m := new(SSHReceivePackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SSHService_SSHUploadArchive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServiceServer).SSHUploadArchive(&sSHServiceSSHUploadArchiveServer{stream})
}

type SSHService_SSHUploadArchiveServer interface {
	Send(*SSHUploadArchiveResponse) error
	Recv() (*SSHUploadArchiveRequest, error)
	grpc.ServerStream
}

type sSHServiceSSHUploadArchiveServer struct {
	grpc.ServerStream
}

func (x *sSHServiceSSHUploadArchiveServer) Send(m *SSHUploadArchiveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHServiceSSHUploadArchiveServer) Recv() (*SSHUploadArchiveRequest, error) {
	m := new(SSHUploadArchiveRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SSHService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.SSHService",
	HandlerType: (*SSHServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SSHUploadPack",
			Handler:       _SSHService_SSHUploadPack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SSHReceivePack",
			Handler:       _SSHService_SSHReceivePack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SSHUploadArchive",
			Handler:       _SSHService_SSHUploadArchive_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ssh.proto",
}

func init() { proto.RegisterFile("ssh.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x89, 0x13, 0xc8, 0xc4, 0x45, 0xd1, 0xd2, 0x16, 0x2b, 0x02, 0x6a, 0xcc, 0xc5, 0x07,
	0x14, 0xa1, 0xf4, 0x0b, 0x10, 0x42, 0x2a, 0x5c, 0xa8, 0xd6, 0xca, 0x89, 0x83, 0x65, 0xec, 0x61,
	0xb3, 0x62, 0xeb, 0x35, 0xbb, 0x9b, 0xa8, 0x95, 0x40, 0x7c, 0x01, 0x37, 0xbe, 0x8b, 0x6f, 0x42,
	0xac, 0x4d, 0xb0, 0xe3, 0xfa, 0x06, 0xb9, 0x79, 0xe6, 0x8d, 0xdf, 0xcc, 0xbc, 0x37, 0x0b, 0x13,
	0xad, 0xd7, 0x8b, 0x52, 0x49, 0x23, 0xc9, 0x98, 0x71, 0x93, 0x8a, 0x9b, 0xb9, 0xa7, 0xd7, 0xa9,
	0xc2, 0xbc, 0xca, 0x86, 0x3f, 0x1d, 0x38, 0x8e, 0xe3, 0x8b, 0x55, 0x29, 0x64, 0x9a, 0x5f, 0xa6,
	0xd9, 0x27, 0x8a, 0x9f, 0x37, 0xa8, 0x0d, 0x59, 0x02, 0x28, 0x2c, 0xa5, 0xe6, 0x46, 0xaa, 0x1b,
	0xdf, 0x09, 0x9c, 0x68, 0xba, 0x24, 0x8b, 0x8a, 0x63, 0x41, 0x77, 0x08, 0x6d, 0x54, 0x91, 0x63,
	0x18, 0x69, 0x93, 0xf3, 0xc2, 0x1f, 0x04, 0x4e, 0xe4, 0xd1, 0x2a, 0x20, 0xcf, 0x81, 0x30, 0x6e,
	0x92, 0x4c, 0x16, 0x1f, 0x39, 0x4b, 0x64, 0x69, 0xb8, 0x2c, 0xb4, 0xef, 0x06, 0xc3, 0x68, 0x42,
	0x67, 0x8c, 0x9b, 0x57, 0x16, 0x78, 0x57, 0xe5, 0xc9, 0x53, 0xf0, 0x7e, 0x57, 0xdb, 0xe9, 0x32,
	0x29, 0xfc, 0x51, 0xe0, 0x44, 0x13, 0x3a, 0x65, 0xdc, 0x5c, 0xd6, 0xa9, 0xb7, 0xee, 0xbd, 0xe1,
	0xcc, 0xa5, 0x27, 0x0d, 0xd2, 0x32, 0x55, 0xe9, 0x15, 0x1a, 0x54, 0x3a, 0xfc, 0x02, 0x27, 0x7b,
	0xfb, 0xe8, 0x52, 0x16, 0x1a, 0xc9, 0x29, 0x8c, 0xb5, 0xc9, 0xe5, 0xc6, 0xd8, 0x65, 0x3c, 0x5a,
	0x47, 0x75, 0x1e, 0x95, 0xaa, 0xa7, 0xae, 0x23, 0x72, 0x0e, 0x53, 0xbc, 0xe6, 0x26, 0xd1, 0x26,
	0x35, 0x1b, 0xed, 0x0f, 0xdb, 0x0a, 0xbc, 0xbe, 0xe6, 0x26, 0xb6, 0x08, 0x05, 0xdc, 0x7d, 0x87,
	0xdf, 0x07, 0xb6, 0x3d, 0xc5, 0x0c, 0xf9, 0x16, 0xff, 0x8f, 0x9e, 0x0f, 0x60, 0xc4, 0x44, 0xc2,
	0x73, 0x3b, 0xd2, 0x84, 0xba, 0x4c, 0xbc, 0xc9, 0xc9, 0x33, 0x38, 0x62, 0x22, 0x69, 0x74, 0x70,
	0x2d, 0xe8, 0x31, 0xf1, 0x97, 0x9b, 0x9c, 0xc1, 0x94, 0x89, 0x64, 0xa3, 0x51, 0x15, 0xe9, 0x15,
	0xd6, 0xd2, 0x02, 0x13, 0xab, 0x3a, 0xd3, 0x11, 0x7f, 0xdc, 0x11, 0xbf, 0xc7, 0xcd, 0xbb, 0xb7,
	0xbb, 0x19, 0x7e, 0x85, 0xd3, 0x7d, 0x39, 0x0e, 0x69, 0x47, 0x06, 0x0f, 0x77, 0xc7, 0xf0, 0x52,
	0x65, 0x6b, 0xbe, 0xc5, 0x7f, 0xee, 0x47, 0xf8, 0x0d, 0xfc, 0x6e, 0x93, 0x03, 0x6e, 0xb9, 0xfc,
	0x31, 0x00, 0x88, 0xe3, 0x8b, 0x18, 0xd5, 0x96, 0x67, 0x48, 0x28, 0x1c, 0xb5, 0x5e, 0x00, 0x79,
	0xf4, 0xe7, 0xff, 0xdb, 0x1e, 0xfa, 0xfc, 0x71, 0x0f, 0x5a, 0x6d, 0x10, 0xde, 0x89, 0x9c, 0x17,
	0x0e, 0x59, 0xc1, 0xfd, 0xb6, 0x8f, 0xa4, 0xf9, 0x5b, 0xf7, 0xdc, 0xe7, 0x4f, 0xfa, 0xe0, 0x16,
	0xed, 0x7b, 0x98, 0xed, 0x4b, 0x47, 0xce, 0x3a, 0xf3, 0xb4, 0x9d, 0x9b, 0x07, 0xfd, 0x05, 0x4d,
	0xf2, 0x0f, 0x63, 0x7b, 0xc6, 0xe7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x98, 0x8e, 0xd7,
	0x04, 0x05, 0x00, 0x00,
}
