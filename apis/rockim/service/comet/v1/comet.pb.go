// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/service/comet/v1/comet.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	service "rockimserver/apis/rockim/service"
	types "rockimserver/apis/rockim/service/comet/v1/types"
	enums "rockimserver/apis/rockim/shared/enums"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DispatchRequest 分发数据请求
type DispatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基础参数
	Base *service.ServiceRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 连接id
	ChannelIds []string `protobuf:"bytes,2,rep,name=channel_ids,json=channelIds,proto3" json:"channel_ids,omitempty"`
	// 数据类型
	DataType enums.Comet_DataType `protobuf:"varint,3,opt,name=data_type,json=dataType,proto3,enum=rockim.shared.enums.Comet_DataType" json:"data_type,omitempty"`
	// 推送消息
	Push *types.PushMessage `protobuf:"bytes,4,opt,name=push,proto3" json:"push,omitempty"`
	// 控制消息
	Control *types.ControlMessage `protobuf:"bytes,5,opt,name=control,proto3" json:"control,omitempty"`
}

func (x *DispatchRequest) Reset() {
	*x = DispatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_comet_v1_comet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchRequest) ProtoMessage() {}

func (x *DispatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_comet_v1_comet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchRequest.ProtoReflect.Descriptor instead.
func (*DispatchRequest) Descriptor() ([]byte, []int) {
	return file_rockim_service_comet_v1_comet_proto_rawDescGZIP(), []int{0}
}

func (x *DispatchRequest) GetBase() *service.ServiceRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *DispatchRequest) GetChannelIds() []string {
	if x != nil {
		return x.ChannelIds
	}
	return nil
}

func (x *DispatchRequest) GetDataType() enums.Comet_DataType {
	if x != nil {
		return x.DataType
	}
	return enums.Comet_PUSH
}

func (x *DispatchRequest) GetPush() *types.PushMessage {
	if x != nil {
		return x.Push
	}
	return nil
}

func (x *DispatchRequest) GetControl() *types.ControlMessage {
	if x != nil {
		return x.Control
	}
	return nil
}

// DispatchResponse 分发数据响应
type DispatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DispatchResponse) Reset() {
	*x = DispatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_comet_v1_comet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchResponse) ProtoMessage() {}

func (x *DispatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_comet_v1_comet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchResponse.ProtoReflect.Descriptor instead.
func (*DispatchResponse) Descriptor() ([]byte, []int) {
	return file_rockim_service_comet_v1_comet_proto_rawDescGZIP(), []int{1}
}

// DispatchRoomRequest 分发房间数据请求
type DispatchRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基础参数
	Base *service.ServiceRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 房间
	Room *types.Room `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	// 推送消息
	Push *types.PushMessage `protobuf:"bytes,3,opt,name=push,proto3" json:"push,omitempty"`
}

func (x *DispatchRoomRequest) Reset() {
	*x = DispatchRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_comet_v1_comet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchRoomRequest) ProtoMessage() {}

func (x *DispatchRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_comet_v1_comet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchRoomRequest.ProtoReflect.Descriptor instead.
func (*DispatchRoomRequest) Descriptor() ([]byte, []int) {
	return file_rockim_service_comet_v1_comet_proto_rawDescGZIP(), []int{2}
}

func (x *DispatchRoomRequest) GetBase() *service.ServiceRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *DispatchRoomRequest) GetRoom() *types.Room {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *DispatchRoomRequest) GetPush() *types.PushMessage {
	if x != nil {
		return x.Push
	}
	return nil
}

// DispatchRoomResponse 分发房间数据响应
type DispatchRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DispatchRoomResponse) Reset() {
	*x = DispatchRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_comet_v1_comet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchRoomResponse) ProtoMessage() {}

func (x *DispatchRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_comet_v1_comet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchRoomResponse.ProtoReflect.Descriptor instead.
func (*DispatchRoomResponse) Descriptor() ([]byte, []int) {
	return file_rockim_service_comet_v1_comet_proto_rawDescGZIP(), []int{3}
}

var File_rockim_service_comet_v1_comet_proto protoreflect.FileDescriptor

var file_rockim_service_comet_v1_comet_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x73, 0x12, 0x40,
	0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x23, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x3e, 0x0a, 0x04, 0x70, 0x75, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x75, 0x73, 0x68,
	0x12, 0x47, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe0, 0x01,
	0x0a, 0x13, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x48, 0x0a, 0x04, 0x70, 0x75, 0x73, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x70, 0x75, 0x73, 0x68,
	0x22, 0x16, 0x0a, 0x14, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xda, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x50, 0x49, 0x12, 0x5f, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x28, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x2c, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_service_comet_v1_comet_proto_rawDescOnce sync.Once
	file_rockim_service_comet_v1_comet_proto_rawDescData = file_rockim_service_comet_v1_comet_proto_rawDesc
)

func file_rockim_service_comet_v1_comet_proto_rawDescGZIP() []byte {
	file_rockim_service_comet_v1_comet_proto_rawDescOnce.Do(func() {
		file_rockim_service_comet_v1_comet_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_service_comet_v1_comet_proto_rawDescData)
	})
	return file_rockim_service_comet_v1_comet_proto_rawDescData
}

var file_rockim_service_comet_v1_comet_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rockim_service_comet_v1_comet_proto_goTypes = []interface{}{
	(*DispatchRequest)(nil),        // 0: rockim.service.comet.v1.DispatchRequest
	(*DispatchResponse)(nil),       // 1: rockim.service.comet.v1.DispatchResponse
	(*DispatchRoomRequest)(nil),    // 2: rockim.service.comet.v1.DispatchRoomRequest
	(*DispatchRoomResponse)(nil),   // 3: rockim.service.comet.v1.DispatchRoomResponse
	(*service.ServiceRequest)(nil), // 4: rockim.service.ServiceRequest
	(enums.Comet_DataType)(0),      // 5: rockim.shared.enums.Comet.DataType
	(*types.PushMessage)(nil),      // 6: rockim.service.comet.v1.types.PushMessage
	(*types.ControlMessage)(nil),   // 7: rockim.service.comet.v1.types.ControlMessage
	(*types.Room)(nil),             // 8: rockim.service.comet.v1.types.Room
}
var file_rockim_service_comet_v1_comet_proto_depIdxs = []int32{
	4, // 0: rockim.service.comet.v1.DispatchRequest.base:type_name -> rockim.service.ServiceRequest
	5, // 1: rockim.service.comet.v1.DispatchRequest.data_type:type_name -> rockim.shared.enums.Comet.DataType
	6, // 2: rockim.service.comet.v1.DispatchRequest.push:type_name -> rockim.service.comet.v1.types.PushMessage
	7, // 3: rockim.service.comet.v1.DispatchRequest.control:type_name -> rockim.service.comet.v1.types.ControlMessage
	4, // 4: rockim.service.comet.v1.DispatchRoomRequest.base:type_name -> rockim.service.ServiceRequest
	8, // 5: rockim.service.comet.v1.DispatchRoomRequest.room:type_name -> rockim.service.comet.v1.types.Room
	6, // 6: rockim.service.comet.v1.DispatchRoomRequest.push:type_name -> rockim.service.comet.v1.types.PushMessage
	0, // 7: rockim.service.comet.v1.ChannelAPI.Dispatch:input_type -> rockim.service.comet.v1.DispatchRequest
	2, // 8: rockim.service.comet.v1.ChannelAPI.DispatchRoom:input_type -> rockim.service.comet.v1.DispatchRoomRequest
	1, // 9: rockim.service.comet.v1.ChannelAPI.Dispatch:output_type -> rockim.service.comet.v1.DispatchResponse
	3, // 10: rockim.service.comet.v1.ChannelAPI.DispatchRoom:output_type -> rockim.service.comet.v1.DispatchRoomResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_rockim_service_comet_v1_comet_proto_init() }
func file_rockim_service_comet_v1_comet_proto_init() {
	if File_rockim_service_comet_v1_comet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_service_comet_v1_comet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchRequest); i {
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
		file_rockim_service_comet_v1_comet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchResponse); i {
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
		file_rockim_service_comet_v1_comet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchRoomRequest); i {
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
		file_rockim_service_comet_v1_comet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchRoomResponse); i {
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
			RawDescriptor: file_rockim_service_comet_v1_comet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rockim_service_comet_v1_comet_proto_goTypes,
		DependencyIndexes: file_rockim_service_comet_v1_comet_proto_depIdxs,
		MessageInfos:      file_rockim_service_comet_v1_comet_proto_msgTypes,
	}.Build()
	File_rockim_service_comet_v1_comet_proto = out.File
	file_rockim_service_comet_v1_comet_proto_rawDesc = nil
	file_rockim_service_comet_v1_comet_proto_goTypes = nil
	file_rockim_service_comet_v1_comet_proto_depIdxs = nil
}
