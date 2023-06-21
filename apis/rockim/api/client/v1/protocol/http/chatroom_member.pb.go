// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/api/client/v1/protocol/http/chatroom_member.proto

package http

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// ChatRoomJoinRequest 聊天室加入请求
type ChatRoomJoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基本参数
	Base *APIRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 自定义群组ID
	CustomGroupId string `protobuf:"bytes,2,opt,name=custom_group_id,json=customGroupId,proto3" json:"custom_group_id,omitempty"`
}

func (x *ChatRoomJoinRequest) Reset() {
	*x = ChatRoomJoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoomJoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoomJoinRequest) ProtoMessage() {}

func (x *ChatRoomJoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRoomJoinRequest.ProtoReflect.Descriptor instead.
func (*ChatRoomJoinRequest) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDescGZIP(), []int{0}
}

func (x *ChatRoomJoinRequest) GetBase() *APIRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *ChatRoomJoinRequest) GetCustomGroupId() string {
	if x != nil {
		return x.CustomGroupId
	}
	return ""
}

// ChatRoomJoinResponse 聊天室加入响应
type ChatRoomJoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChatRoomJoinResponse) Reset() {
	*x = ChatRoomJoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoomJoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoomJoinResponse) ProtoMessage() {}

func (x *ChatRoomJoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRoomJoinResponse.ProtoReflect.Descriptor instead.
func (*ChatRoomJoinResponse) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDescGZIP(), []int{1}
}

// ChatRoomQuitRequest 聊天室退出请求
type ChatRoomQuitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基本参数
	Base *APIRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 自定义群组ID
	CustomGroupId string `protobuf:"bytes,3,opt,name=custom_group_id,json=customGroupId,proto3" json:"custom_group_id,omitempty"`
}

func (x *ChatRoomQuitRequest) Reset() {
	*x = ChatRoomQuitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoomQuitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoomQuitRequest) ProtoMessage() {}

func (x *ChatRoomQuitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRoomQuitRequest.ProtoReflect.Descriptor instead.
func (*ChatRoomQuitRequest) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDescGZIP(), []int{2}
}

func (x *ChatRoomQuitRequest) GetBase() *APIRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *ChatRoomQuitRequest) GetCustomGroupId() string {
	if x != nil {
		return x.CustomGroupId
	}
	return ""
}

// ChatRoomQuitResponse 聊天室退出响应
type ChatRoomQuitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChatRoomQuitResponse) Reset() {
	*x = ChatRoomQuitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoomQuitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoomQuitResponse) ProtoMessage() {}

func (x *ChatRoomQuitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRoomQuitResponse.ProtoReflect.Descriptor instead.
func (*ChatRoomQuitResponse) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDescGZIP(), []int{3}
}

var File_rockim_api_client_v1_protocol_http_chatroom_member_proto protoreflect.FileDescriptor

var file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDesc = []byte{
	0x0a, 0x38, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d,
	0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x41,
	0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xe3, 0x02, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0xa5, 0x01, 0x0a, 0x04, 0x4a, 0x6f, 0x69,
	0x6e, 0x12, 0x37, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f,
	0x6d, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x6a, 0x6f, 0x69, 0x6e, 0x3a, 0x01, 0x2a,
	0x12, 0xa5, 0x01, 0x0a, 0x04, 0x51, 0x75, 0x69, 0x74, 0x12, 0x37, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x38, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d,
	0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x2f, 0x71, 0x75, 0x69, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x60, 0x5a, 0x39, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f,
	0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x3b, 0x68, 0x74, 0x74, 0x70, 0xaa, 0x02, 0x22, 0x52, 0x6f, 0x63, 0x6b, 0x49, 0x4d, 0x2e, 0x41,
	0x70, 0x69, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDescOnce sync.Once
	file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDescData = file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDesc
)

func file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDescGZIP() []byte {
	file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDescOnce.Do(func() {
		file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDescData)
	})
	return file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDescData
}

var file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rockim_api_client_v1_protocol_http_chatroom_member_proto_goTypes = []interface{}{
	(*ChatRoomJoinRequest)(nil),  // 0: rockim.api.client.v1.protocol.http.ChatRoomJoinRequest
	(*ChatRoomJoinResponse)(nil), // 1: rockim.api.client.v1.protocol.http.ChatRoomJoinResponse
	(*ChatRoomQuitRequest)(nil),  // 2: rockim.api.client.v1.protocol.http.ChatRoomQuitRequest
	(*ChatRoomQuitResponse)(nil), // 3: rockim.api.client.v1.protocol.http.ChatRoomQuitResponse
	(*APIRequest)(nil),           // 4: rockim.api.client.v1.protocol.http.APIRequest
}
var file_rockim_api_client_v1_protocol_http_chatroom_member_proto_depIdxs = []int32{
	4, // 0: rockim.api.client.v1.protocol.http.ChatRoomJoinRequest.base:type_name -> rockim.api.client.v1.protocol.http.APIRequest
	4, // 1: rockim.api.client.v1.protocol.http.ChatRoomQuitRequest.base:type_name -> rockim.api.client.v1.protocol.http.APIRequest
	0, // 2: rockim.api.client.v1.protocol.http.ChatRoomMemberAPI.Join:input_type -> rockim.api.client.v1.protocol.http.ChatRoomJoinRequest
	2, // 3: rockim.api.client.v1.protocol.http.ChatRoomMemberAPI.Quit:input_type -> rockim.api.client.v1.protocol.http.ChatRoomQuitRequest
	1, // 4: rockim.api.client.v1.protocol.http.ChatRoomMemberAPI.Join:output_type -> rockim.api.client.v1.protocol.http.ChatRoomJoinResponse
	3, // 5: rockim.api.client.v1.protocol.http.ChatRoomMemberAPI.Quit:output_type -> rockim.api.client.v1.protocol.http.ChatRoomQuitResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rockim_api_client_v1_protocol_http_chatroom_member_proto_init() }
func file_rockim_api_client_v1_protocol_http_chatroom_member_proto_init() {
	if File_rockim_api_client_v1_protocol_http_chatroom_member_proto != nil {
		return
	}
	file_rockim_api_client_v1_protocol_http_http_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRoomJoinRequest); i {
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
		file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRoomJoinResponse); i {
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
		file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRoomQuitRequest); i {
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
		file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRoomQuitResponse); i {
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
			RawDescriptor: file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rockim_api_client_v1_protocol_http_chatroom_member_proto_goTypes,
		DependencyIndexes: file_rockim_api_client_v1_protocol_http_chatroom_member_proto_depIdxs,
		MessageInfos:      file_rockim_api_client_v1_protocol_http_chatroom_member_proto_msgTypes,
	}.Build()
	File_rockim_api_client_v1_protocol_http_chatroom_member_proto = out.File
	file_rockim_api_client_v1_protocol_http_chatroom_member_proto_rawDesc = nil
	file_rockim_api_client_v1_protocol_http_chatroom_member_proto_goTypes = nil
	file_rockim_api_client_v1_protocol_http_chatroom_member_proto_depIdxs = nil
}
func (x *ChatRoomJoinRequest) SetBase(base *APIRequest) {
	if x == nil {
		return
	}
	x.Base = base
}

func (x *ChatRoomJoinRequest) SetCustomGroupId(customGroupId string) {
	if x == nil {
		return
	}
	x.CustomGroupId = customGroupId
}

func (x *ChatRoomQuitRequest) SetBase(base *APIRequest) {
	if x == nil {
		return
	}
	x.Base = base
}

func (x *ChatRoomQuitRequest) SetCustomGroupId(customGroupId string) {
	if x == nil {
		return
	}
	x.CustomGroupId = customGroupId
}
