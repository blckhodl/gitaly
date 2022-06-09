// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: server.proto

package gitalypb

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

// This comment is left unintentionally blank.
type ServerInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerInfoRequest) Reset() {
	*x = ServerInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfoRequest) ProtoMessage() {}

func (x *ServerInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfoRequest.ProtoReflect.Descriptor instead.
func (*ServerInfoRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

// This comment is left unintentionally blank.
type ServerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This comment is left unintentionally blank.
	ServerVersion string `protobuf:"bytes,1,opt,name=server_version,json=serverVersion,proto3" json:"server_version,omitempty"`
	// This comment is left unintentionally blank.
	GitVersion string `protobuf:"bytes,2,opt,name=git_version,json=gitVersion,proto3" json:"git_version,omitempty"`
	// This comment is left unintentionally blank.
	StorageStatuses []*ServerInfoResponse_StorageStatus `protobuf:"bytes,3,rep,name=storage_statuses,json=storageStatuses,proto3" json:"storage_statuses,omitempty"`
}

func (x *ServerInfoResponse) Reset() {
	*x = ServerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfoResponse) ProtoMessage() {}

func (x *ServerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfoResponse.ProtoReflect.Descriptor instead.
func (*ServerInfoResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *ServerInfoResponse) GetServerVersion() string {
	if x != nil {
		return x.ServerVersion
	}
	return ""
}

func (x *ServerInfoResponse) GetGitVersion() string {
	if x != nil {
		return x.GitVersion
	}
	return ""
}

func (x *ServerInfoResponse) GetStorageStatuses() []*ServerInfoResponse_StorageStatus {
	if x != nil {
		return x.StorageStatuses
	}
	return nil
}

// This comment is left unintentionally blank.
type DiskStatisticsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DiskStatisticsRequest) Reset() {
	*x = DiskStatisticsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskStatisticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskStatisticsRequest) ProtoMessage() {}

func (x *DiskStatisticsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskStatisticsRequest.ProtoReflect.Descriptor instead.
func (*DiskStatisticsRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

// This comment is left unintentionally blank.
type DiskStatisticsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This comment is left unintentionally blank.
	StorageStatuses []*DiskStatisticsResponse_StorageStatus `protobuf:"bytes,1,rep,name=storage_statuses,json=storageStatuses,proto3" json:"storage_statuses,omitempty"`
}

func (x *DiskStatisticsResponse) Reset() {
	*x = DiskStatisticsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskStatisticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskStatisticsResponse) ProtoMessage() {}

func (x *DiskStatisticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskStatisticsResponse.ProtoReflect.Descriptor instead.
func (*DiskStatisticsResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3}
}

func (x *DiskStatisticsResponse) GetStorageStatuses() []*DiskStatisticsResponse_StorageStatus {
	if x != nil {
		return x.StorageStatuses
	}
	return nil
}

// This comment is left unintentionally blank.
type ClockSyncedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ntp_host is a URL to the external NTP service that should be used for clock sync check.
	// Default is ntp.pool.org
	NtpHost string `protobuf:"bytes,1,opt,name=ntp_host,json=ntpHost,proto3" json:"ntp_host,omitempty"`
	// drift_threshold_millis is an allowed drift from the NTP service in milliseconds.
	DriftThresholdMillis int64 `protobuf:"varint,2,opt,name=drift_threshold_millis,json=driftThresholdMillis,proto3" json:"drift_threshold_millis,omitempty"`
}

func (x *ClockSyncedRequest) Reset() {
	*x = ClockSyncedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockSyncedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockSyncedRequest) ProtoMessage() {}

func (x *ClockSyncedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockSyncedRequest.ProtoReflect.Descriptor instead.
func (*ClockSyncedRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{4}
}

func (x *ClockSyncedRequest) GetNtpHost() string {
	if x != nil {
		return x.NtpHost
	}
	return ""
}

func (x *ClockSyncedRequest) GetDriftThresholdMillis() int64 {
	if x != nil {
		return x.DriftThresholdMillis
	}
	return 0
}

// This comment is left unintentionally blank.
type ClockSyncedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// synced is set to true if system clock has an affordable drift compared to NTP service.
	Synced bool `protobuf:"varint,1,opt,name=synced,proto3" json:"synced,omitempty"`
}

func (x *ClockSyncedResponse) Reset() {
	*x = ClockSyncedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockSyncedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockSyncedResponse) ProtoMessage() {}

func (x *ClockSyncedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockSyncedResponse.ProtoReflect.Descriptor instead.
func (*ClockSyncedResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{5}
}

func (x *ClockSyncedResponse) GetSynced() bool {
	if x != nil {
		return x.Synced
	}
	return false
}

// This comment is left unintentionally blank.
type ServerInfoResponse_StorageStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This comment is left unintentionally blank.
	StorageName string `protobuf:"bytes,1,opt,name=storage_name,json=storageName,proto3" json:"storage_name,omitempty"`
	// This comment is left unintentionally blank.
	Readable bool `protobuf:"varint,2,opt,name=readable,proto3" json:"readable,omitempty"`
	// This comment is left unintentionally blank.
	Writeable bool `protobuf:"varint,3,opt,name=writeable,proto3" json:"writeable,omitempty"`
	// This comment is left unintentionally blank.
	FsType string `protobuf:"bytes,4,opt,name=fs_type,json=fsType,proto3" json:"fs_type,omitempty"`
	// This comment is left unintentionally blank.
	FilesystemId string `protobuf:"bytes,5,opt,name=filesystem_id,json=filesystemId,proto3" json:"filesystem_id,omitempty"`
	// This comment is left unintentionally blank.
	ReplicationFactor uint32 `protobuf:"varint,6,opt,name=replication_factor,json=replicationFactor,proto3" json:"replication_factor,omitempty"`
}

func (x *ServerInfoResponse_StorageStatus) Reset() {
	*x = ServerInfoResponse_StorageStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfoResponse_StorageStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfoResponse_StorageStatus) ProtoMessage() {}

func (x *ServerInfoResponse_StorageStatus) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfoResponse_StorageStatus.ProtoReflect.Descriptor instead.
func (*ServerInfoResponse_StorageStatus) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ServerInfoResponse_StorageStatus) GetStorageName() string {
	if x != nil {
		return x.StorageName
	}
	return ""
}

func (x *ServerInfoResponse_StorageStatus) GetReadable() bool {
	if x != nil {
		return x.Readable
	}
	return false
}

func (x *ServerInfoResponse_StorageStatus) GetWriteable() bool {
	if x != nil {
		return x.Writeable
	}
	return false
}

func (x *ServerInfoResponse_StorageStatus) GetFsType() string {
	if x != nil {
		return x.FsType
	}
	return ""
}

func (x *ServerInfoResponse_StorageStatus) GetFilesystemId() string {
	if x != nil {
		return x.FilesystemId
	}
	return ""
}

func (x *ServerInfoResponse_StorageStatus) GetReplicationFactor() uint32 {
	if x != nil {
		return x.ReplicationFactor
	}
	return 0
}

// This comment is left unintentionally blank.
type DiskStatisticsResponse_StorageStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When both available and used fields are equal 0 that means that
	// Gitaly was unable to determine storage stats.
	StorageName string `protobuf:"bytes,1,opt,name=storage_name,json=storageName,proto3" json:"storage_name,omitempty"`
	// This comment is left unintentionally blank.
	Available int64 `protobuf:"varint,2,opt,name=available,proto3" json:"available,omitempty"`
	// This comment is left unintentionally blank.
	Used int64 `protobuf:"varint,3,opt,name=used,proto3" json:"used,omitempty"`
}

func (x *DiskStatisticsResponse_StorageStatus) Reset() {
	*x = DiskStatisticsResponse_StorageStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskStatisticsResponse_StorageStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskStatisticsResponse_StorageStatus) ProtoMessage() {}

func (x *DiskStatisticsResponse_StorageStatus) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskStatisticsResponse_StorageStatus.ProtoReflect.Descriptor instead.
func (*DiskStatisticsResponse_StorageStatus) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3, 0}
}

func (x *DiskStatisticsResponse_StorageStatus) GetStorageName() string {
	if x != nil {
		return x.StorageName
	}
	return ""
}

func (x *DiskStatisticsResponse_StorageStatus) GetAvailable() int64 {
	if x != nil {
		return x.Available
	}
	return 0
}

func (x *DiskStatisticsResponse_StorageStatus) GetUsed() int64 {
	if x != nil {
		return x.Used
	}
	return 0
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x79, 0x1a, 0x0a, 0x6c, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8d, 0x03, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x69, 0x74, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x69, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x1a, 0xd9, 0x01, 0x0a, 0x0d,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x73,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x69, 0x73, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xd7, 0x01, 0x0a, 0x16, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x79, 0x2e, 0x44,
	0x69, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x65, 0x73, 0x1a, 0x64, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x22, 0x65, 0x0a, 0x12, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6e, 0x74, 0x70, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x74, 0x70, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x64,
	0x72, 0x69, 0x66, 0x74, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x6d,
	0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x64, 0x72, 0x69,
	0x66, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4d, 0x69, 0x6c, 0x6c, 0x69,
	0x73, 0x22, 0x2d, 0x0a, 0x13, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x63,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64,
	0x32, 0xf3, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x19, 0x2e, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x69,
	0x74, 0x61, 0x6c, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x79, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x79, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x43, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x12, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x79,
	0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x79, 0x2e, 0x43, 0x6c, 0x6f,
	0x63, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x1a, 0x04, 0xf0, 0x97, 0x28, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x79, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_server_proto_goTypes = []interface{}{
	(*ServerInfoRequest)(nil),                    // 0: gitaly.ServerInfoRequest
	(*ServerInfoResponse)(nil),                   // 1: gitaly.ServerInfoResponse
	(*DiskStatisticsRequest)(nil),                // 2: gitaly.DiskStatisticsRequest
	(*DiskStatisticsResponse)(nil),               // 3: gitaly.DiskStatisticsResponse
	(*ClockSyncedRequest)(nil),                   // 4: gitaly.ClockSyncedRequest
	(*ClockSyncedResponse)(nil),                  // 5: gitaly.ClockSyncedResponse
	(*ServerInfoResponse_StorageStatus)(nil),     // 6: gitaly.ServerInfoResponse.StorageStatus
	(*DiskStatisticsResponse_StorageStatus)(nil), // 7: gitaly.DiskStatisticsResponse.StorageStatus
}
var file_server_proto_depIdxs = []int32{
	6, // 0: gitaly.ServerInfoResponse.storage_statuses:type_name -> gitaly.ServerInfoResponse.StorageStatus
	7, // 1: gitaly.DiskStatisticsResponse.storage_statuses:type_name -> gitaly.DiskStatisticsResponse.StorageStatus
	0, // 2: gitaly.ServerService.ServerInfo:input_type -> gitaly.ServerInfoRequest
	2, // 3: gitaly.ServerService.DiskStatistics:input_type -> gitaly.DiskStatisticsRequest
	4, // 4: gitaly.ServerService.ClockSynced:input_type -> gitaly.ClockSyncedRequest
	1, // 5: gitaly.ServerService.ServerInfo:output_type -> gitaly.ServerInfoResponse
	3, // 6: gitaly.ServerService.DiskStatistics:output_type -> gitaly.DiskStatisticsResponse
	5, // 7: gitaly.ServerService.ClockSynced:output_type -> gitaly.ClockSyncedResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	file_lint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfoRequest); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfoResponse); i {
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
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskStatisticsRequest); i {
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
		file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskStatisticsResponse); i {
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
		file_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockSyncedRequest); i {
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
		file_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockSyncedResponse); i {
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
		file_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfoResponse_StorageStatus); i {
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
		file_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskStatisticsResponse_StorageStatus); i {
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
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
