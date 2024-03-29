// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/api/client/v1/protocol/socket/protocol.proto

package socket

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	enums "rockimserver/apis/rockim/shared/enums"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Operation int32

const (
	// 无效值
	Operation_INVALID_REQUEST Operation = 0
	// 授权包
	Operation_AUTH Operation = 1
	// 心跳
	Operation_HEARTBEAT Operation = 2
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0: "INVALID_REQUEST",
		1: "AUTH",
		2: "HEARTBEAT",
	}
	Operation_value = map[string]int32{
		"INVALID_REQUEST": 0,
		"AUTH":            1,
		"HEARTBEAT":       2,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_rockim_api_client_v1_protocol_socket_protocol_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_rockim_api_client_v1_protocol_socket_protocol_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescGZIP(), []int{0}
}

// RequestPacketHeader 客户端请求包头部
type RequestPacketHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=rockim.api.client.v1.protocol.socket.Operation" json:"operation,omitempty"`
	// 请求id
	RequestId int32 `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *RequestPacketHeader) Reset() {
	*x = RequestPacketHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPacketHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPacketHeader) ProtoMessage() {}

func (x *RequestPacketHeader) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPacketHeader.ProtoReflect.Descriptor instead.
func (*RequestPacketHeader) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *RequestPacketHeader) GetOperation() Operation {
	if x != nil {
		return x.Operation
	}
	return Operation_INVALID_REQUEST
}

func (x *RequestPacketHeader) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

// RequestPacketBody 客户端请求包内容
type RequestPacketBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body []byte `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *RequestPacketBody) Reset() {
	*x = RequestPacketBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPacketBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPacketBody) ProtoMessage() {}

func (x *RequestPacketBody) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPacketBody.ProtoReflect.Descriptor instead.
func (*RequestPacketBody) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *RequestPacketBody) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// ResponsePacketHeader 服务端响应包头部
type ResponsePacketHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=rockim.api.client.v1.protocol.socket.Operation" json:"operation,omitempty"`
	// 请求id
	RequestId int32 `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// 是否操作成功
	Success bool `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	// 链路追踪ID
	TraceId string `protobuf:"bytes,4,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (x *ResponsePacketHeader) Reset() {
	*x = ResponsePacketHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePacketHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePacketHeader) ProtoMessage() {}

func (x *ResponsePacketHeader) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePacketHeader.ProtoReflect.Descriptor instead.
func (*ResponsePacketHeader) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *ResponsePacketHeader) GetOperation() Operation {
	if x != nil {
		return x.Operation
	}
	return Operation_INVALID_REQUEST
}

func (x *ResponsePacketHeader) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *ResponsePacketHeader) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ResponsePacketHeader) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

// ResponsePacketBody 服务端响应包内容
type ResponsePacketBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body []byte `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ResponsePacketBody) Reset() {
	*x = ResponsePacketBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePacketBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePacketBody) ProtoMessage() {}

func (x *ResponsePacketBody) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePacketBody.ProtoReflect.Descriptor instead.
func (*ResponsePacketBody) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescGZIP(), []int{3}
}

func (x *ResponsePacketBody) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// PushPacketHeader 服务端推送包头部
type PushPacketHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation enums.Comet_PushOperation `protobuf:"varint,1,opt,name=operation,proto3,enum=rockim.shared.enums.Comet_PushOperation" json:"operation,omitempty"`
}

func (x *PushPacketHeader) Reset() {
	*x = PushPacketHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushPacketHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushPacketHeader) ProtoMessage() {}

func (x *PushPacketHeader) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushPacketHeader.ProtoReflect.Descriptor instead.
func (*PushPacketHeader) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescGZIP(), []int{4}
}

func (x *PushPacketHeader) GetOperation() enums.Comet_PushOperation {
	if x != nil {
		return x.Operation
	}
	return enums.Comet_INVALID_PUSH
}

// PushPacketBody 服务端推送包内容
type PushPacketBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body []byte `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PushPacketBody) Reset() {
	*x = PushPacketBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushPacketBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushPacketBody) ProtoMessage() {}

func (x *PushPacketBody) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushPacketBody.ProtoReflect.Descriptor instead.
func (*PushPacketBody) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescGZIP(), []int{5}
}

func (x *PushPacketBody) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_rockim_api_client_v1_protocol_socket_protocol_proto protoreflect.FileDescriptor

var file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDesc = []byte{
	0x0a, 0x33, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x1f, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a,
	0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x22, 0x27, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xb9, 0x01, 0x0a, 0x14,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0x5a, 0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a,
	0x0e, 0x50, 0x75, 0x73, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x2a, 0x39, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x55, 0x54, 0x48, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x10, 0x02, 0x42, 0x66,
	0x5a, 0x3d, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x3b, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0xaa,
	0x02, 0x24, 0x52, 0x6f, 0x63, 0x6b, 0x49, 0x4d, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescOnce sync.Once
	file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescData = file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDesc
)

func file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescGZIP() []byte {
	file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescOnce.Do(func() {
		file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescData)
	})
	return file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDescData
}

var file_rockim_api_client_v1_protocol_socket_protocol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rockim_api_client_v1_protocol_socket_protocol_proto_goTypes = []interface{}{
	(Operation)(0),                 // 0: rockim.api.client.v1.protocol.socket.Operation
	(*RequestPacketHeader)(nil),    // 1: rockim.api.client.v1.protocol.socket.RequestPacketHeader
	(*RequestPacketBody)(nil),      // 2: rockim.api.client.v1.protocol.socket.RequestPacketBody
	(*ResponsePacketHeader)(nil),   // 3: rockim.api.client.v1.protocol.socket.ResponsePacketHeader
	(*ResponsePacketBody)(nil),     // 4: rockim.api.client.v1.protocol.socket.ResponsePacketBody
	(*PushPacketHeader)(nil),       // 5: rockim.api.client.v1.protocol.socket.PushPacketHeader
	(*PushPacketBody)(nil),         // 6: rockim.api.client.v1.protocol.socket.PushPacketBody
	(enums.Comet_PushOperation)(0), // 7: rockim.shared.enums.Comet.PushOperation
}
var file_rockim_api_client_v1_protocol_socket_protocol_proto_depIdxs = []int32{
	0, // 0: rockim.api.client.v1.protocol.socket.RequestPacketHeader.operation:type_name -> rockim.api.client.v1.protocol.socket.Operation
	0, // 1: rockim.api.client.v1.protocol.socket.ResponsePacketHeader.operation:type_name -> rockim.api.client.v1.protocol.socket.Operation
	7, // 2: rockim.api.client.v1.protocol.socket.PushPacketHeader.operation:type_name -> rockim.shared.enums.Comet.PushOperation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rockim_api_client_v1_protocol_socket_protocol_proto_init() }
func file_rockim_api_client_v1_protocol_socket_protocol_proto_init() {
	if File_rockim_api_client_v1_protocol_socket_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPacketHeader); i {
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
		file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPacketBody); i {
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
		file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsePacketHeader); i {
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
		file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsePacketBody); i {
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
		file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushPacketHeader); i {
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
		file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushPacketBody); i {
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
			RawDescriptor: file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rockim_api_client_v1_protocol_socket_protocol_proto_goTypes,
		DependencyIndexes: file_rockim_api_client_v1_protocol_socket_protocol_proto_depIdxs,
		EnumInfos:         file_rockim_api_client_v1_protocol_socket_protocol_proto_enumTypes,
		MessageInfos:      file_rockim_api_client_v1_protocol_socket_protocol_proto_msgTypes,
	}.Build()
	File_rockim_api_client_v1_protocol_socket_protocol_proto = out.File
	file_rockim_api_client_v1_protocol_socket_protocol_proto_rawDesc = nil
	file_rockim_api_client_v1_protocol_socket_protocol_proto_goTypes = nil
	file_rockim_api_client_v1_protocol_socket_protocol_proto_depIdxs = nil
}
func (x *RequestPacketHeader) SetOperation(operation Operation) {
	if x == nil {
		return
	}
	x.Operation = operation
}

func (x *RequestPacketHeader) SetRequestId(requestId int32) {
	if x == nil {
		return
	}
	x.RequestId = requestId
}

func (x *RequestPacketBody) SetBody(body []byte) {
	if x == nil {
		return
	}
	x.Body = body
}

func (x *ResponsePacketHeader) SetOperation(operation Operation) {
	if x == nil {
		return
	}
	x.Operation = operation
}

func (x *ResponsePacketHeader) SetRequestId(requestId int32) {
	if x == nil {
		return
	}
	x.RequestId = requestId
}

func (x *ResponsePacketHeader) SetSuccess(success bool) {
	if x == nil {
		return
	}
	x.Success = success
}

func (x *ResponsePacketHeader) SetTraceId(traceId string) {
	if x == nil {
		return
	}
	x.TraceId = traceId
}

func (x *ResponsePacketBody) SetBody(body []byte) {
	if x == nil {
		return
	}
	x.Body = body
}

func (x *PushPacketHeader) SetOperation(operation enums.Comet_PushOperation) {
	if x == nil {
		return
	}
	x.Operation = operation
}

func (x *PushPacketBody) SetBody(body []byte) {
	if x == nil {
		return
	}
	x.Body = body
}
