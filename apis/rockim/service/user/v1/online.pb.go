// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/service/user/v1/online.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	service "rockimserver/apis/rockim/service"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// OnlineAddRequest 添加在线连接请求
type OnlineAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基础参数
	Base *service.ServiceRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 服务器id
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// 连接id
	ChannelId string `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// 访问令牌
	Token string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *OnlineAddRequest) Reset() {
	*x = OnlineAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_user_v1_online_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineAddRequest) ProtoMessage() {}

func (x *OnlineAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_user_v1_online_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineAddRequest.ProtoReflect.Descriptor instead.
func (*OnlineAddRequest) Descriptor() ([]byte, []int) {
	return file_rockim_service_user_v1_online_proto_rawDescGZIP(), []int{0}
}

func (x *OnlineAddRequest) GetBase() *service.ServiceRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OnlineAddRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *OnlineAddRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *OnlineAddRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// OnlineAddResponse 添加在线连接响应
type OnlineAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *OnlineAddResponse) Reset() {
	*x = OnlineAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_user_v1_online_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineAddResponse) ProtoMessage() {}

func (x *OnlineAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_user_v1_online_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineAddResponse.ProtoReflect.Descriptor instead.
func (*OnlineAddResponse) Descriptor() ([]byte, []int) {
	return file_rockim_service_user_v1_online_proto_rawDescGZIP(), []int{1}
}

func (x *OnlineAddResponse) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

// OnlineDeleteRequest 移除在线连接请求
type OnlineDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基础参数
	Base *service.ServiceRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 服务器id
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// 连接id
	ChannelId string `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *OnlineDeleteRequest) Reset() {
	*x = OnlineDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_user_v1_online_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineDeleteRequest) ProtoMessage() {}

func (x *OnlineDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_user_v1_online_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineDeleteRequest.ProtoReflect.Descriptor instead.
func (*OnlineDeleteRequest) Descriptor() ([]byte, []int) {
	return file_rockim_service_user_v1_online_proto_rawDescGZIP(), []int{2}
}

func (x *OnlineDeleteRequest) GetBase() *service.ServiceRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OnlineDeleteRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *OnlineDeleteRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

// OnlineDeleteResponse 移除在线连接响应
type OnlineDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnlineDeleteResponse) Reset() {
	*x = OnlineDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_user_v1_online_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineDeleteResponse) ProtoMessage() {}

func (x *OnlineDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_user_v1_online_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineDeleteResponse.ProtoReflect.Descriptor instead.
func (*OnlineDeleteResponse) Descriptor() ([]byte, []int) {
	return file_rockim_service_user_v1_online_proto_rawDescGZIP(), []int{3}
}

// OnlineRefreshRequest 刷新在线连接请求
type OnlineRefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基础参数
	Base *service.ServiceRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 服务器id
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// 连接id
	ChannelId string `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *OnlineRefreshRequest) Reset() {
	*x = OnlineRefreshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_user_v1_online_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineRefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineRefreshRequest) ProtoMessage() {}

func (x *OnlineRefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_user_v1_online_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineRefreshRequest.ProtoReflect.Descriptor instead.
func (*OnlineRefreshRequest) Descriptor() ([]byte, []int) {
	return file_rockim_service_user_v1_online_proto_rawDescGZIP(), []int{4}
}

func (x *OnlineRefreshRequest) GetBase() *service.ServiceRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OnlineRefreshRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *OnlineRefreshRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

// OnlineRefreshResponse 刷新在线连接响应
type OnlineRefreshResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnlineRefreshResponse) Reset() {
	*x = OnlineRefreshResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_user_v1_online_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineRefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineRefreshResponse) ProtoMessage() {}

func (x *OnlineRefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_user_v1_online_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineRefreshResponse.ProtoReflect.Descriptor instead.
func (*OnlineRefreshResponse) Descriptor() ([]byte, []int) {
	return file_rockim_service_user_v1_online_proto_rawDescGZIP(), []int{5}
}

var File_rockim_service_user_v1_online_proto protoreflect.FileDescriptor

var file_rockim_service_user_v1_online_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x10, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x25, 0x0a, 0x11, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0xa1, 0x01, 0x0a, 0x13,
	0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22,
	0x16, 0x0a, 0x14, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x14, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3c, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15,
	0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb4, 0x02, 0x0a, 0x09, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x41, 0x50, 0x49, 0x12, 0x5a, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x28, 0x2e, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x63, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x07, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12,
	0x2c, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rockim_service_user_v1_online_proto_rawDescOnce sync.Once
	file_rockim_service_user_v1_online_proto_rawDescData = file_rockim_service_user_v1_online_proto_rawDesc
)

func file_rockim_service_user_v1_online_proto_rawDescGZIP() []byte {
	file_rockim_service_user_v1_online_proto_rawDescOnce.Do(func() {
		file_rockim_service_user_v1_online_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_service_user_v1_online_proto_rawDescData)
	})
	return file_rockim_service_user_v1_online_proto_rawDescData
}

var file_rockim_service_user_v1_online_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rockim_service_user_v1_online_proto_goTypes = []interface{}{
	(*OnlineAddRequest)(nil),       // 0: rockim.service.user.v1.OnlineAddRequest
	(*OnlineAddResponse)(nil),      // 1: rockim.service.user.v1.OnlineAddResponse
	(*OnlineDeleteRequest)(nil),    // 2: rockim.service.user.v1.OnlineDeleteRequest
	(*OnlineDeleteResponse)(nil),   // 3: rockim.service.user.v1.OnlineDeleteResponse
	(*OnlineRefreshRequest)(nil),   // 4: rockim.service.user.v1.OnlineRefreshRequest
	(*OnlineRefreshResponse)(nil),  // 5: rockim.service.user.v1.OnlineRefreshResponse
	(*service.ServiceRequest)(nil), // 6: rockim.service.ServiceRequest
}
var file_rockim_service_user_v1_online_proto_depIdxs = []int32{
	6, // 0: rockim.service.user.v1.OnlineAddRequest.base:type_name -> rockim.service.ServiceRequest
	6, // 1: rockim.service.user.v1.OnlineDeleteRequest.base:type_name -> rockim.service.ServiceRequest
	6, // 2: rockim.service.user.v1.OnlineRefreshRequest.base:type_name -> rockim.service.ServiceRequest
	0, // 3: rockim.service.user.v1.OnlineAPI.Add:input_type -> rockim.service.user.v1.OnlineAddRequest
	2, // 4: rockim.service.user.v1.OnlineAPI.Delete:input_type -> rockim.service.user.v1.OnlineDeleteRequest
	4, // 5: rockim.service.user.v1.OnlineAPI.Refresh:input_type -> rockim.service.user.v1.OnlineRefreshRequest
	1, // 6: rockim.service.user.v1.OnlineAPI.Add:output_type -> rockim.service.user.v1.OnlineAddResponse
	3, // 7: rockim.service.user.v1.OnlineAPI.Delete:output_type -> rockim.service.user.v1.OnlineDeleteResponse
	5, // 8: rockim.service.user.v1.OnlineAPI.Refresh:output_type -> rockim.service.user.v1.OnlineRefreshResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rockim_service_user_v1_online_proto_init() }
func file_rockim_service_user_v1_online_proto_init() {
	if File_rockim_service_user_v1_online_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_service_user_v1_online_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineAddRequest); i {
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
		file_rockim_service_user_v1_online_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineAddResponse); i {
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
		file_rockim_service_user_v1_online_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineDeleteRequest); i {
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
		file_rockim_service_user_v1_online_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineDeleteResponse); i {
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
		file_rockim_service_user_v1_online_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineRefreshRequest); i {
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
		file_rockim_service_user_v1_online_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineRefreshResponse); i {
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
			RawDescriptor: file_rockim_service_user_v1_online_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rockim_service_user_v1_online_proto_goTypes,
		DependencyIndexes: file_rockim_service_user_v1_online_proto_depIdxs,
		MessageInfos:      file_rockim_service_user_v1_online_proto_msgTypes,
	}.Build()
	File_rockim_service_user_v1_online_proto = out.File
	file_rockim_service_user_v1_online_proto_rawDesc = nil
	file_rockim_service_user_v1_online_proto_goTypes = nil
	file_rockim_service_user_v1_online_proto_depIdxs = nil
}
