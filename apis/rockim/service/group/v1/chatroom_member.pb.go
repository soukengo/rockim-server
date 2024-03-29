// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/service/group/v1/chatroom_member.proto

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

// ChatRoomJoinRequest 加入聊天室请求
type ChatRoomJoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基础参数
	Base *service.ServiceRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 群组id
	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// 用户账号
	Uid string `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ChatRoomJoinRequest) Reset() {
	*x = ChatRoomJoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_group_v1_chatroom_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoomJoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoomJoinRequest) ProtoMessage() {}

func (x *ChatRoomJoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_group_v1_chatroom_member_proto_msgTypes[0]
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
	return file_rockim_service_group_v1_chatroom_member_proto_rawDescGZIP(), []int{0}
}

func (x *ChatRoomJoinRequest) GetBase() *service.ServiceRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *ChatRoomJoinRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ChatRoomJoinRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

// ChatRoomJoinResponse 加入聊天室响应
type ChatRoomJoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChatRoomJoinResponse) Reset() {
	*x = ChatRoomJoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_group_v1_chatroom_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoomJoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoomJoinResponse) ProtoMessage() {}

func (x *ChatRoomJoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_group_v1_chatroom_member_proto_msgTypes[1]
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
	return file_rockim_service_group_v1_chatroom_member_proto_rawDescGZIP(), []int{1}
}

// ChatRoomQuitRequest 退出聊天室请求
type ChatRoomQuitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基础参数
	Base *service.ServiceRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 群组id
	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// 用户账号
	Uid string `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ChatRoomQuitRequest) Reset() {
	*x = ChatRoomQuitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_group_v1_chatroom_member_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoomQuitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoomQuitRequest) ProtoMessage() {}

func (x *ChatRoomQuitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_group_v1_chatroom_member_proto_msgTypes[2]
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
	return file_rockim_service_group_v1_chatroom_member_proto_rawDescGZIP(), []int{2}
}

func (x *ChatRoomQuitRequest) GetBase() *service.ServiceRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *ChatRoomQuitRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ChatRoomQuitRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

// ChatRoomQuitResponse 退出聊天室响应
type ChatRoomQuitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChatRoomQuitResponse) Reset() {
	*x = ChatRoomQuitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_group_v1_chatroom_member_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoomQuitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoomQuitResponse) ProtoMessage() {}

func (x *ChatRoomQuitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_group_v1_chatroom_member_proto_msgTypes[3]
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
	return file_rockim_service_group_v1_chatroom_member_proto_rawDescGZIP(), []int{3}
}

var File_rockim_service_group_v1_chatroom_member_proto protoreflect.FileDescriptor

var file_rockim_service_group_v1_chatroom_member_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x99, 0x01, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x18, 0x40, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d,
	0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x40, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x16, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x51, 0x75, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xdd, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0x63, 0x0a,
	0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x2c, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x63, 0x0a, 0x04, 0x51, 0x75, 0x69, 0x74, 0x12, 0x2c, 0x2e, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x51, 0x75, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x51, 0x75, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_service_group_v1_chatroom_member_proto_rawDescOnce sync.Once
	file_rockim_service_group_v1_chatroom_member_proto_rawDescData = file_rockim_service_group_v1_chatroom_member_proto_rawDesc
)

func file_rockim_service_group_v1_chatroom_member_proto_rawDescGZIP() []byte {
	file_rockim_service_group_v1_chatroom_member_proto_rawDescOnce.Do(func() {
		file_rockim_service_group_v1_chatroom_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_service_group_v1_chatroom_member_proto_rawDescData)
	})
	return file_rockim_service_group_v1_chatroom_member_proto_rawDescData
}

var file_rockim_service_group_v1_chatroom_member_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rockim_service_group_v1_chatroom_member_proto_goTypes = []interface{}{
	(*ChatRoomJoinRequest)(nil),    // 0: rockim.service.group.v1.ChatRoomJoinRequest
	(*ChatRoomJoinResponse)(nil),   // 1: rockim.service.group.v1.ChatRoomJoinResponse
	(*ChatRoomQuitRequest)(nil),    // 2: rockim.service.group.v1.ChatRoomQuitRequest
	(*ChatRoomQuitResponse)(nil),   // 3: rockim.service.group.v1.ChatRoomQuitResponse
	(*service.ServiceRequest)(nil), // 4: rockim.service.ServiceRequest
}
var file_rockim_service_group_v1_chatroom_member_proto_depIdxs = []int32{
	4, // 0: rockim.service.group.v1.ChatRoomJoinRequest.base:type_name -> rockim.service.ServiceRequest
	4, // 1: rockim.service.group.v1.ChatRoomQuitRequest.base:type_name -> rockim.service.ServiceRequest
	0, // 2: rockim.service.group.v1.ChatRoomMemberAPI.Join:input_type -> rockim.service.group.v1.ChatRoomJoinRequest
	2, // 3: rockim.service.group.v1.ChatRoomMemberAPI.Quit:input_type -> rockim.service.group.v1.ChatRoomQuitRequest
	1, // 4: rockim.service.group.v1.ChatRoomMemberAPI.Join:output_type -> rockim.service.group.v1.ChatRoomJoinResponse
	3, // 5: rockim.service.group.v1.ChatRoomMemberAPI.Quit:output_type -> rockim.service.group.v1.ChatRoomQuitResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rockim_service_group_v1_chatroom_member_proto_init() }
func file_rockim_service_group_v1_chatroom_member_proto_init() {
	if File_rockim_service_group_v1_chatroom_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_service_group_v1_chatroom_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_rockim_service_group_v1_chatroom_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_rockim_service_group_v1_chatroom_member_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_rockim_service_group_v1_chatroom_member_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_rockim_service_group_v1_chatroom_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rockim_service_group_v1_chatroom_member_proto_goTypes,
		DependencyIndexes: file_rockim_service_group_v1_chatroom_member_proto_depIdxs,
		MessageInfos:      file_rockim_service_group_v1_chatroom_member_proto_msgTypes,
	}.Build()
	File_rockim_service_group_v1_chatroom_member_proto = out.File
	file_rockim_service_group_v1_chatroom_member_proto_rawDesc = nil
	file_rockim_service_group_v1_chatroom_member_proto_goTypes = nil
	file_rockim_service_group_v1_chatroom_member_proto_depIdxs = nil
}
