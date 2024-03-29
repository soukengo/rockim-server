// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/service/message/v1/types/shared.proto

package types

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// ConversationID 会话ID
type ConversationID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 目标类型
	Category enums.MessageTarget_Category `protobuf:"varint,1,opt,name=category,proto3,enum=rockim.shared.enums.MessageTarget_Category" json:"category,omitempty"`
	// 值（群聊为groupId，私聊为用户uid1#用户uid2）
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ConversationID) Reset() {
	*x = ConversationID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_message_v1_types_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationID) ProtoMessage() {}

func (x *ConversationID) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_message_v1_types_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationID.ProtoReflect.Descriptor instead.
func (*ConversationID) Descriptor() ([]byte, []int) {
	return file_rockim_service_message_v1_types_shared_proto_rawDescGZIP(), []int{0}
}

func (x *ConversationID) GetCategory() enums.MessageTarget_Category {
	if x != nil {
		return x.Category
	}
	return enums.MessageTarget_UNKNOWN
}

func (x *ConversationID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// TargetID 目标ID
type TargetID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 目标类型
	Category enums.MessageTarget_Category `protobuf:"varint,1,opt,name=category,proto3,enum=rockim.shared.enums.MessageTarget_Category" json:"category,omitempty"`
	// 值（群聊为customGroupId，私聊为用户account）
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TargetID) Reset() {
	*x = TargetID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_message_v1_types_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetID) ProtoMessage() {}

func (x *TargetID) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_message_v1_types_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetID.ProtoReflect.Descriptor instead.
func (*TargetID) Descriptor() ([]byte, []int) {
	return file_rockim_service_message_v1_types_shared_proto_rawDescGZIP(), []int{1}
}

func (x *TargetID) GetCategory() enums.MessageTarget_Category {
	if x != nil {
		return x.Category
	}
	return enums.MessageTarget_UNKNOWN
}

func (x *TargetID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_rockim_service_message_v1_types_shared_proto protoreflect.FileDescriptor

var file_rockim_service_message_v1_types_shared_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x51,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2b, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x7c, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x12, 0x51, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b,
	0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x39,
	0x5a, 0x37, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rockim_service_message_v1_types_shared_proto_rawDescOnce sync.Once
	file_rockim_service_message_v1_types_shared_proto_rawDescData = file_rockim_service_message_v1_types_shared_proto_rawDesc
)

func file_rockim_service_message_v1_types_shared_proto_rawDescGZIP() []byte {
	file_rockim_service_message_v1_types_shared_proto_rawDescOnce.Do(func() {
		file_rockim_service_message_v1_types_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_service_message_v1_types_shared_proto_rawDescData)
	})
	return file_rockim_service_message_v1_types_shared_proto_rawDescData
}

var file_rockim_service_message_v1_types_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rockim_service_message_v1_types_shared_proto_goTypes = []interface{}{
	(*ConversationID)(nil),            // 0: rockim.service.message.v1.types.ConversationID
	(*TargetID)(nil),                  // 1: rockim.service.message.v1.types.TargetID
	(enums.MessageTarget_Category)(0), // 2: rockim.shared.enums.MessageTarget.Category
}
var file_rockim_service_message_v1_types_shared_proto_depIdxs = []int32{
	2, // 0: rockim.service.message.v1.types.ConversationID.category:type_name -> rockim.shared.enums.MessageTarget.Category
	2, // 1: rockim.service.message.v1.types.TargetID.category:type_name -> rockim.shared.enums.MessageTarget.Category
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rockim_service_message_v1_types_shared_proto_init() }
func file_rockim_service_message_v1_types_shared_proto_init() {
	if File_rockim_service_message_v1_types_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_service_message_v1_types_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationID); i {
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
		file_rockim_service_message_v1_types_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetID); i {
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
			RawDescriptor: file_rockim_service_message_v1_types_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rockim_service_message_v1_types_shared_proto_goTypes,
		DependencyIndexes: file_rockim_service_message_v1_types_shared_proto_depIdxs,
		MessageInfos:      file_rockim_service_message_v1_types_shared_proto_msgTypes,
	}.Build()
	File_rockim_service_message_v1_types_shared_proto = out.File
	file_rockim_service_message_v1_types_shared_proto_rawDesc = nil
	file_rockim_service_message_v1_types_shared_proto_goTypes = nil
	file_rockim_service_message_v1_types_shared_proto_depIdxs = nil
}
