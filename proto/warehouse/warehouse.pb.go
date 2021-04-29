// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/warehouse/warehouse.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total  int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{0}
}

func (x *Pagination) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Pagination) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type WHTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UploadId   int64                  `protobuf:"varint,2,opt,name=upload_id,json=uploadId,proto3" json:"upload_id,omitempty"`
	Name       string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Status     string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Error      string                 `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	LastExecAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_exec_at,json=lastExecAt,proto3" json:"last_exec_at,omitempty"`
	Count      int32                  `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"`
	Duration   int32                  `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *WHTable) Reset() {
	*x = WHTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WHTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WHTable) ProtoMessage() {}

func (x *WHTable) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WHTable.ProtoReflect.Descriptor instead.
func (*WHTable) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{1}
}

func (x *WHTable) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WHTable) GetUploadId() int64 {
	if x != nil {
		return x.UploadId
	}
	return 0
}

func (x *WHTable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WHTable) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WHTable) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *WHTable) GetLastExecAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastExecAt
	}
	return nil
}

func (x *WHTable) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *WHTable) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type WHUploadsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId        string `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	DestinationId   string `protobuf:"bytes,2,opt,name=destination_id,json=destinationId,proto3" json:"destination_id,omitempty"`
	DestinationType string `protobuf:"bytes,3,opt,name=destination_type,json=destinationType,proto3" json:"destination_type,omitempty"`
	Status          string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Limit           int32  `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset          int32  `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	WorkspaceId     string `protobuf:"bytes,7,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
}

func (x *WHUploadsRequest) Reset() {
	*x = WHUploadsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WHUploadsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WHUploadsRequest) ProtoMessage() {}

func (x *WHUploadsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WHUploadsRequest.ProtoReflect.Descriptor instead.
func (*WHUploadsRequest) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{2}
}

func (x *WHUploadsRequest) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *WHUploadsRequest) GetDestinationId() string {
	if x != nil {
		return x.DestinationId
	}
	return ""
}

func (x *WHUploadsRequest) GetDestinationType() string {
	if x != nil {
		return x.DestinationType
	}
	return ""
}

func (x *WHUploadsRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WHUploadsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *WHUploadsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *WHUploadsRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type WHUploadsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uploads    []*WHUploadResponse `protobuf:"bytes,1,rep,name=uploads,proto3" json:"uploads,omitempty"`
	Pagination *Pagination         `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *WHUploadsResponse) Reset() {
	*x = WHUploadsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WHUploadsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WHUploadsResponse) ProtoMessage() {}

func (x *WHUploadsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WHUploadsResponse.ProtoReflect.Descriptor instead.
func (*WHUploadsResponse) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{3}
}

func (x *WHUploadsResponse) GetUploads() []*WHUploadResponse {
	if x != nil {
		return x.Uploads
	}
	return nil
}

func (x *WHUploadsResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type WHUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadId    int64  `protobuf:"varint,1,opt,name=upload_id,json=uploadId,proto3" json:"upload_id,omitempty"`
	WorkspaceId string `protobuf:"bytes,2,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
}

func (x *WHUploadRequest) Reset() {
	*x = WHUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WHUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WHUploadRequest) ProtoMessage() {}

func (x *WHUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WHUploadRequest.ProtoReflect.Descriptor instead.
func (*WHUploadRequest) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{4}
}

func (x *WHUploadRequest) GetUploadId() int64 {
	if x != nil {
		return x.UploadId
	}
	return 0
}

func (x *WHUploadRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type WHUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SourceId        string                 `protobuf:"bytes,2,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	DestinationId   string                 `protobuf:"bytes,3,opt,name=destination_id,json=destinationId,proto3" json:"destination_id,omitempty"`
	DestinationType string                 `protobuf:"bytes,4,opt,name=destination_type,json=destinationType,proto3" json:"destination_type,omitempty"`
	Namespace       string                 `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Error           string                 `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	Attempt         int32                  `protobuf:"varint,7,opt,name=attempt,proto3" json:"attempt,omitempty"`
	Status          string                 `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	FirstEventAt    *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=first_event_at,json=firstEventAt,proto3" json:"first_event_at,omitempty"`
	LastEventAt     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=last_event_at,json=lastEventAt,proto3" json:"last_event_at,omitempty"`
	LastExecAt      *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=last_exec_at,json=lastExecAt,proto3" json:"last_exec_at,omitempty"`
	NextRetryTime   *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=next_retry_time,json=nextRetryTime,proto3" json:"next_retry_time,omitempty"`
	Duration        int32                  `protobuf:"varint,14,opt,name=duration,proto3" json:"duration,omitempty"`
	Tables          []*WHTable             `protobuf:"bytes,15,rep,name=tables,proto3" json:"tables,omitempty"`
}

func (x *WHUploadResponse) Reset() {
	*x = WHUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WHUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WHUploadResponse) ProtoMessage() {}

func (x *WHUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WHUploadResponse.ProtoReflect.Descriptor instead.
func (*WHUploadResponse) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{5}
}

func (x *WHUploadResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WHUploadResponse) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *WHUploadResponse) GetDestinationId() string {
	if x != nil {
		return x.DestinationId
	}
	return ""
}

func (x *WHUploadResponse) GetDestinationType() string {
	if x != nil {
		return x.DestinationType
	}
	return ""
}

func (x *WHUploadResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *WHUploadResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *WHUploadResponse) GetAttempt() int32 {
	if x != nil {
		return x.Attempt
	}
	return 0
}

func (x *WHUploadResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WHUploadResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *WHUploadResponse) GetFirstEventAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FirstEventAt
	}
	return nil
}

func (x *WHUploadResponse) GetLastEventAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastEventAt
	}
	return nil
}

func (x *WHUploadResponse) GetLastExecAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastExecAt
	}
	return nil
}

func (x *WHUploadResponse) GetNextRetryTime() *timestamppb.Timestamp {
	if x != nil {
		return x.NextRetryTime
	}
	return nil
}

func (x *WHUploadResponse) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *WHUploadResponse) GetTables() []*WHTable {
	if x != nil {
		return x.Tables
	}
	return nil
}

var File_proto_warehouse_warehouse_proto protoreflect.FileDescriptor

var file_proto_warehouse_warehouse_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xe8, 0x01, 0x0a, 0x07, 0x57, 0x48, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x63,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x41,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xea, 0x01, 0x0a, 0x10, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x22, 0x79, 0x0a, 0x11, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57,
	0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x07, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x0f, 0x57,
	0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0xfa,
	0x04, 0x0a, 0x10, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x41, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x78,
	0x65, 0x63, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x65,
	0x63, 0x41, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x0f, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x48, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x32, 0xcf, 0x01, 0x0a, 0x09,
	0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x48, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x48, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_warehouse_warehouse_proto_rawDescOnce sync.Once
	file_proto_warehouse_warehouse_proto_rawDescData = file_proto_warehouse_warehouse_proto_rawDesc
)

func file_proto_warehouse_warehouse_proto_rawDescGZIP() []byte {
	file_proto_warehouse_warehouse_proto_rawDescOnce.Do(func() {
		file_proto_warehouse_warehouse_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_warehouse_warehouse_proto_rawDescData)
	})
	return file_proto_warehouse_warehouse_proto_rawDescData
}

var file_proto_warehouse_warehouse_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_warehouse_warehouse_proto_goTypes = []interface{}{
	(*Pagination)(nil),            // 0: proto.Pagination
	(*WHTable)(nil),               // 1: proto.WHTable
	(*WHUploadsRequest)(nil),      // 2: proto.WHUploadsRequest
	(*WHUploadsResponse)(nil),     // 3: proto.WHUploadsResponse
	(*WHUploadRequest)(nil),       // 4: proto.WHUploadRequest
	(*WHUploadResponse)(nil),      // 5: proto.WHUploadResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
	(*wrapperspb.BoolValue)(nil),  // 8: google.protobuf.BoolValue
}
var file_proto_warehouse_warehouse_proto_depIdxs = []int32{
	6,  // 0: proto.WHTable.last_exec_at:type_name -> google.protobuf.Timestamp
	5,  // 1: proto.WHUploadsResponse.uploads:type_name -> proto.WHUploadResponse
	0,  // 2: proto.WHUploadsResponse.pagination:type_name -> proto.Pagination
	6,  // 3: proto.WHUploadResponse.created_at:type_name -> google.protobuf.Timestamp
	6,  // 4: proto.WHUploadResponse.first_event_at:type_name -> google.protobuf.Timestamp
	6,  // 5: proto.WHUploadResponse.last_event_at:type_name -> google.protobuf.Timestamp
	6,  // 6: proto.WHUploadResponse.last_exec_at:type_name -> google.protobuf.Timestamp
	6,  // 7: proto.WHUploadResponse.next_retry_time:type_name -> google.protobuf.Timestamp
	1,  // 8: proto.WHUploadResponse.tables:type_name -> proto.WHTable
	7,  // 9: proto.Warehouse.GetHealth:input_type -> google.protobuf.Empty
	2,  // 10: proto.Warehouse.GetWHUploads:input_type -> proto.WHUploadsRequest
	4,  // 11: proto.Warehouse.GetWHUpload:input_type -> proto.WHUploadRequest
	8,  // 12: proto.Warehouse.GetHealth:output_type -> google.protobuf.BoolValue
	3,  // 13: proto.Warehouse.GetWHUploads:output_type -> proto.WHUploadsResponse
	5,  // 14: proto.Warehouse.GetWHUpload:output_type -> proto.WHUploadResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_warehouse_warehouse_proto_init() }
func file_proto_warehouse_warehouse_proto_init() {
	if File_proto_warehouse_warehouse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_warehouse_warehouse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WHTable); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WHUploadsRequest); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WHUploadsResponse); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WHUploadRequest); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WHUploadResponse); i {
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
			RawDescriptor: file_proto_warehouse_warehouse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_warehouse_warehouse_proto_goTypes,
		DependencyIndexes: file_proto_warehouse_warehouse_proto_depIdxs,
		MessageInfos:      file_proto_warehouse_warehouse_proto_msgTypes,
	}.Build()
	File_proto_warehouse_warehouse_proto = out.File
	file_proto_warehouse_warehouse_proto_rawDesc = nil
	file_proto_warehouse_warehouse_proto_goTypes = nil
	file_proto_warehouse_warehouse_proto_depIdxs = nil
}
