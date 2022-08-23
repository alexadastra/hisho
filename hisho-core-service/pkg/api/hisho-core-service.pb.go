// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: hisho-core-service.proto

package api

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Term int32

const (
	Term_OTHER Term = 0
	Term_WEEK  Term = 1
	Term_TODAY Term = 2
)

// Enum value maps for Term.
var (
	Term_name = map[int32]string{
		0: "OTHER",
		1: "WEEK",
		2: "TODAY",
	}
	Term_value = map[string]int32{
		"OTHER": 0,
		"WEEK":  1,
		"TODAY": 2,
	}
)

func (x Term) Enum() *Term {
	p := new(Term)
	*p = x
	return p
}

func (x Term) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Term) Descriptor() protoreflect.EnumDescriptor {
	return file_hisho_core_service_proto_enumTypes[0].Descriptor()
}

func (Term) Type() protoreflect.EnumType {
	return &file_hisho_core_service_proto_enumTypes[0]
}

func (x Term) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Term.Descriptor instead.
func (Term) EnumDescriptor() ([]byte, []int) {
	return file_hisho_core_service_proto_rawDescGZIP(), []int{0}
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hisho_core_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hisho_core_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_hisho_core_service_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hisho_core_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hisho_core_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_hisho_core_service_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PingResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Term      string                 `protobuf:"bytes,3,opt,name=term,proto3" json:"term,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DoneAt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=done_at,json=doneAt,proto3" json:"done_at,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hisho_core_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_hisho_core_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_hisho_core_service_proto_rawDescGZIP(), []int{2}
}

func (x *Task) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

func (x *Task) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Task) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Task) GetDoneAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DoneAt
	}
	return nil
}

type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term Term `protobuf:"varint,1,opt,name=term,proto3,enum=hisho_core_service.Term" json:"term,omitempty"`
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hisho_core_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hisho_core_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_hisho_core_service_proto_rawDescGZIP(), []int{3}
}

func (x *TaskRequest) GetTerm() Term {
	if x != nil {
		return x.Term
	}
	return Term_OTHER
}

type TasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TasksResponse) Reset() {
	*x = TasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hisho_core_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksResponse) ProtoMessage() {}

func (x *TasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hisho_core_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TasksResponse.ProtoReflect.Descriptor instead.
func (*TasksResponse) Descriptor() ([]byte, []int) {
	return file_hisho_core_service_proto_rawDescGZIP(), []int{4}
}

func (x *TasksResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hisho_core_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_hisho_core_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_hisho_core_service_proto_rawDescGZIP(), []int{5}
}

func (x *Msg) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_hisho_core_service_proto protoreflect.FileDescriptor

var file_hisho_core_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x68, 0x69, 0x73, 0x68, 0x6f, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x68, 0x69, 0x73, 0x68,
	0x6f, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0b, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x38, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xeb, 0x01, 0x0a, 0x04, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x64, 0x6f, 0x6e, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x06, 0x64, 0x6f, 0x6e, 0x65, 0x41, 0x74, 0x22, 0x3b, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x68, 0x69, 0x73, 0x68, 0x6f, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x22, 0x3f, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x69, 0x73, 0x68, 0x6f, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x1f, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x26, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x09,
	0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x45, 0x45,
	0x4b, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x4f, 0x44, 0x41, 0x59, 0x10, 0x02, 0x32, 0xc2,
	0x02, 0x0a, 0x10, 0x48, 0x69, 0x73, 0x68, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x68, 0x69,
	0x73, 0x68, 0x6f, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68,
	0x69, 0x73, 0x68, 0x6f, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x6e, 0x67,
	0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x42,
	0x79, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x1f, 0x2e, 0x68, 0x69, 0x73, 0x68, 0x6f, 0x5f, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x69, 0x73, 0x68, 0x6f, 0x5f, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x67, 0x65, 0x74,
	0x5f, 0x62, 0x79, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x3a, 0x01, 0x2a, 0x12, 0x56, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x2e, 0x68, 0x69, 0x73, 0x68, 0x6f, 0x5f, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x1a, 0x17, 0x2e, 0x68, 0x69, 0x73, 0x68, 0x6f, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x61, 0x64, 0x64,
	0x3a, 0x01, 0x2a, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hisho_core_service_proto_rawDescOnce sync.Once
	file_hisho_core_service_proto_rawDescData = file_hisho_core_service_proto_rawDesc
)

func file_hisho_core_service_proto_rawDescGZIP() []byte {
	file_hisho_core_service_proto_rawDescOnce.Do(func() {
		file_hisho_core_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_hisho_core_service_proto_rawDescData)
	})
	return file_hisho_core_service_proto_rawDescData
}

var file_hisho_core_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hisho_core_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_hisho_core_service_proto_goTypes = []interface{}{
	(Term)(0),                     // 0: hisho_core_service.Term
	(*PingRequest)(nil),           // 1: hisho_core_service.PingRequest
	(*PingResponse)(nil),          // 2: hisho_core_service.PingResponse
	(*Task)(nil),                  // 3: hisho_core_service.Task
	(*TaskRequest)(nil),           // 4: hisho_core_service.TaskRequest
	(*TasksResponse)(nil),         // 5: hisho_core_service.TasksResponse
	(*Msg)(nil),                   // 6: hisho_core_service.Msg
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_hisho_core_service_proto_depIdxs = []int32{
	7, // 0: hisho_core_service.Task.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: hisho_core_service.Task.updated_at:type_name -> google.protobuf.Timestamp
	7, // 2: hisho_core_service.Task.done_at:type_name -> google.protobuf.Timestamp
	0, // 3: hisho_core_service.TaskRequest.term:type_name -> hisho_core_service.Term
	3, // 4: hisho_core_service.TasksResponse.tasks:type_name -> hisho_core_service.Task
	1, // 5: hisho_core_service.HishoCoreService.Ping:input_type -> hisho_core_service.PingRequest
	4, // 6: hisho_core_service.HishoCoreService.GetTasksByTerm:input_type -> hisho_core_service.TaskRequest
	3, // 7: hisho_core_service.HishoCoreService.AddTask:input_type -> hisho_core_service.Task
	2, // 8: hisho_core_service.HishoCoreService.Ping:output_type -> hisho_core_service.PingResponse
	5, // 9: hisho_core_service.HishoCoreService.GetTasksByTerm:output_type -> hisho_core_service.TasksResponse
	6, // 10: hisho_core_service.HishoCoreService.AddTask:output_type -> hisho_core_service.Msg
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_hisho_core_service_proto_init() }
func file_hisho_core_service_proto_init() {
	if File_hisho_core_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hisho_core_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_hisho_core_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_hisho_core_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_hisho_core_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequest); i {
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
		file_hisho_core_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TasksResponse); i {
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
		file_hisho_core_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
			RawDescriptor: file_hisho_core_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hisho_core_service_proto_goTypes,
		DependencyIndexes: file_hisho_core_service_proto_depIdxs,
		EnumInfos:         file_hisho_core_service_proto_enumTypes,
		MessageInfos:      file_hisho_core_service_proto_msgTypes,
	}.Build()
	File_hisho_core_service_proto = out.File
	file_hisho_core_service_proto_rawDesc = nil
	file_hisho_core_service_proto_goTypes = nil
	file_hisho_core_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HishoCoreServiceClient is the client API for HishoCoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HishoCoreServiceClient interface {
	// ping pong
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// get tasks for terms (today, week, other etc.)
	GetTasksByTerm(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TasksResponse, error)
	// add new task
	AddTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Msg, error)
}

type hishoCoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHishoCoreServiceClient(cc grpc.ClientConnInterface) HishoCoreServiceClient {
	return &hishoCoreServiceClient{cc}
}

func (c *hishoCoreServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/hisho_core_service.HishoCoreService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hishoCoreServiceClient) GetTasksByTerm(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TasksResponse, error) {
	out := new(TasksResponse)
	err := c.cc.Invoke(ctx, "/hisho_core_service.HishoCoreService/GetTasksByTerm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hishoCoreServiceClient) AddTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, "/hisho_core_service.HishoCoreService/AddTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HishoCoreServiceServer is the server API for HishoCoreService service.
type HishoCoreServiceServer interface {
	// ping pong
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// get tasks for terms (today, week, other etc.)
	GetTasksByTerm(context.Context, *TaskRequest) (*TasksResponse, error)
	// add new task
	AddTask(context.Context, *Task) (*Msg, error)
}

// UnimplementedHishoCoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHishoCoreServiceServer struct {
}

func (*UnimplementedHishoCoreServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedHishoCoreServiceServer) GetTasksByTerm(context.Context, *TaskRequest) (*TasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasksByTerm not implemented")
}
func (*UnimplementedHishoCoreServiceServer) AddTask(context.Context, *Task) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTask not implemented")
}

func RegisterHishoCoreServiceServer(s *grpc.Server, srv HishoCoreServiceServer) {
	s.RegisterService(&_HishoCoreService_serviceDesc, srv)
}

func _HishoCoreService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HishoCoreServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hisho_core_service.HishoCoreService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HishoCoreServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HishoCoreService_GetTasksByTerm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HishoCoreServiceServer).GetTasksByTerm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hisho_core_service.HishoCoreService/GetTasksByTerm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HishoCoreServiceServer).GetTasksByTerm(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HishoCoreService_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HishoCoreServiceServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hisho_core_service.HishoCoreService/AddTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HishoCoreServiceServer).AddTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

var _HishoCoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hisho_core_service.HishoCoreService",
	HandlerType: (*HishoCoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _HishoCoreService_Ping_Handler,
		},
		{
			MethodName: "GetTasksByTerm",
			Handler:    _HishoCoreService_GetTasksByTerm_Handler,
		},
		{
			MethodName: "AddTask",
			Handler:    _HishoCoreService_AddTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hisho-core-service.proto",
}