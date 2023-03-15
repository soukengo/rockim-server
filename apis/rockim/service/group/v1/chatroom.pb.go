// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/service/group/v1/chatroom.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	service "rockimserver/apis/rockim/service"
	types "rockimserver/apis/rockim/service/group/v1/types"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ChatRoomCreateRequest 创建聊天室请求
type ChatRoomCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基础参数
	Base *service.ServiceRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 自定义群组id
	CustomGroupId string `protobuf:"bytes,2,opt,name=custom_group_id,json=customGroupId,proto3" json:"custom_group_id,omitempty"`
	// 群组名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 图标
	IconUrl string `protobuf:"bytes,4,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
	// 扩展字段
	Fields map[string]string `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 群主UID
	Owner string `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *ChatRoomCreateRequest) Reset() {
	*x = ChatRoomCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_group_v1_chatroom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoomCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoomCreateRequest) ProtoMessage() {}

func (x *ChatRoomCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_group_v1_chatroom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRoomCreateRequest.ProtoReflect.Descriptor instead.
func (*ChatRoomCreateRequest) Descriptor() ([]byte, []int) {
	return file_rockim_service_group_v1_chatroom_proto_rawDescGZIP(), []int{0}
}

func (x *ChatRoomCreateRequest) GetBase() *service.ServiceRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *ChatRoomCreateRequest) GetCustomGroupId() string {
	if x != nil {
		return x.CustomGroupId
	}
	return ""
}

func (x *ChatRoomCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChatRoomCreateRequest) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *ChatRoomCreateRequest) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *ChatRoomCreateRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

// ChatRoomCreateResponse 创建聊天室响应
type ChatRoomCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *types.Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *ChatRoomCreateResponse) Reset() {
	*x = ChatRoomCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_group_v1_chatroom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoomCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoomCreateResponse) ProtoMessage() {}

func (x *ChatRoomCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_group_v1_chatroom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRoomCreateResponse.ProtoReflect.Descriptor instead.
func (*ChatRoomCreateResponse) Descriptor() ([]byte, []int) {
	return file_rockim_service_group_v1_chatroom_proto_rawDescGZIP(), []int{1}
}

func (x *ChatRoomCreateResponse) GetGroup() *types.Group {
	if x != nil {
		return x.Group
	}
	return nil
}

// ChatRoomDismissRequest 解散聊天室请求
type ChatRoomDismissRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基础参数
	Base *service.ServiceRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 群组id
	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *ChatRoomDismissRequest) Reset() {
	*x = ChatRoomDismissRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_group_v1_chatroom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoomDismissRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoomDismissRequest) ProtoMessage() {}

func (x *ChatRoomDismissRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_group_v1_chatroom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRoomDismissRequest.ProtoReflect.Descriptor instead.
func (*ChatRoomDismissRequest) Descriptor() ([]byte, []int) {
	return file_rockim_service_group_v1_chatroom_proto_rawDescGZIP(), []int{2}
}

func (x *ChatRoomDismissRequest) GetBase() *service.ServiceRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *ChatRoomDismissRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

// ChatRoomDismissResponse 解散聊天室响应
type ChatRoomDismissResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChatRoomDismissResponse) Reset() {
	*x = ChatRoomDismissResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_group_v1_chatroom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoomDismissResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoomDismissResponse) ProtoMessage() {}

func (x *ChatRoomDismissResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_group_v1_chatroom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRoomDismissResponse.ProtoReflect.Descriptor instead.
func (*ChatRoomDismissResponse) Descriptor() ([]byte, []int) {
	return file_rockim_service_group_v1_chatroom_proto_rawDescGZIP(), []int{3}
}

var File_rockim_service_group_v1_chatroom_proto protoreflect.FileDescriptor

var file_rockim_service_group_v1_chatroom_proto_rawDesc = []byte{
	0x0a, 0x26, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f,
	0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x03, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xfa, 0x42, 0x1a, 0x72, 0x18, 0x32, 0x16, 0x5e, 0x5b, 0x61,
	0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x2e, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x36,
	0x34, 0x7d, 0x24, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0e, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x40,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18,
	0x80, 0x02, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x52, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x54, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x7a, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f,
	0x6f, 0x6d, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3c, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x40, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x69,
	0x73, 0x6d, 0x69, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe6, 0x01,
	0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x50, 0x49, 0x12, 0x69, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x6d,
	0x69, 0x73, 0x73, 0x12, 0x2f, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_service_group_v1_chatroom_proto_rawDescOnce sync.Once
	file_rockim_service_group_v1_chatroom_proto_rawDescData = file_rockim_service_group_v1_chatroom_proto_rawDesc
)

func file_rockim_service_group_v1_chatroom_proto_rawDescGZIP() []byte {
	file_rockim_service_group_v1_chatroom_proto_rawDescOnce.Do(func() {
		file_rockim_service_group_v1_chatroom_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_service_group_v1_chatroom_proto_rawDescData)
	})
	return file_rockim_service_group_v1_chatroom_proto_rawDescData
}

var file_rockim_service_group_v1_chatroom_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rockim_service_group_v1_chatroom_proto_goTypes = []interface{}{
	(*ChatRoomCreateRequest)(nil),   // 0: rockim.service.group.v1.ChatRoomCreateRequest
	(*ChatRoomCreateResponse)(nil),  // 1: rockim.service.group.v1.ChatRoomCreateResponse
	(*ChatRoomDismissRequest)(nil),  // 2: rockim.service.group.v1.ChatRoomDismissRequest
	(*ChatRoomDismissResponse)(nil), // 3: rockim.service.group.v1.ChatRoomDismissResponse
	nil,                             // 4: rockim.service.group.v1.ChatRoomCreateRequest.FieldsEntry
	(*service.ServiceRequest)(nil),  // 5: rockim.service.ServiceRequest
	(*types.Group)(nil),             // 6: rockim.service.group.v1.types.Group
}
var file_rockim_service_group_v1_chatroom_proto_depIdxs = []int32{
	5, // 0: rockim.service.group.v1.ChatRoomCreateRequest.base:type_name -> rockim.service.ServiceRequest
	4, // 1: rockim.service.group.v1.ChatRoomCreateRequest.fields:type_name -> rockim.service.group.v1.ChatRoomCreateRequest.FieldsEntry
	6, // 2: rockim.service.group.v1.ChatRoomCreateResponse.group:type_name -> rockim.service.group.v1.types.Group
	5, // 3: rockim.service.group.v1.ChatRoomDismissRequest.base:type_name -> rockim.service.ServiceRequest
	0, // 4: rockim.service.group.v1.ChatRoomAPI.Create:input_type -> rockim.service.group.v1.ChatRoomCreateRequest
	2, // 5: rockim.service.group.v1.ChatRoomAPI.Dismiss:input_type -> rockim.service.group.v1.ChatRoomDismissRequest
	1, // 6: rockim.service.group.v1.ChatRoomAPI.Create:output_type -> rockim.service.group.v1.ChatRoomCreateResponse
	3, // 7: rockim.service.group.v1.ChatRoomAPI.Dismiss:output_type -> rockim.service.group.v1.ChatRoomDismissResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rockim_service_group_v1_chatroom_proto_init() }
func file_rockim_service_group_v1_chatroom_proto_init() {
	if File_rockim_service_group_v1_chatroom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_service_group_v1_chatroom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRoomCreateRequest); i {
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
		file_rockim_service_group_v1_chatroom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRoomCreateResponse); i {
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
		file_rockim_service_group_v1_chatroom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRoomDismissRequest); i {
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
		file_rockim_service_group_v1_chatroom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRoomDismissResponse); i {
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
			RawDescriptor: file_rockim_service_group_v1_chatroom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rockim_service_group_v1_chatroom_proto_goTypes,
		DependencyIndexes: file_rockim_service_group_v1_chatroom_proto_depIdxs,
		MessageInfos:      file_rockim_service_group_v1_chatroom_proto_msgTypes,
	}.Build()
	File_rockim_service_group_v1_chatroom_proto = out.File
	file_rockim_service_group_v1_chatroom_proto_rawDesc = nil
	file_rockim_service_group_v1_chatroom_proto_goTypes = nil
	file_rockim_service_group_v1_chatroom_proto_depIdxs = nil
}
