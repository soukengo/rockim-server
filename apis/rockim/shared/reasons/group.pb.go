// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/shared/reasons/group.proto

package reasons

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

type Group_ErrorReason int32

const (
	// 群组不存在
	Group_GROUP_NOT_FOUND Group_ErrorReason = 0
	// 群组重复
	Group_GROUP_DUPLICATED Group_ErrorReason = 1
	// 群组成员不存在
	Group_GROUP_MEMBER_NOT_FOUND Group_ErrorReason = 2
	// 自定义群组格式不正确
	Group_CUSTOM_GROUP_ID_INVALID Group_ErrorReason = 3
)

// Enum value maps for Group_ErrorReason.
var (
	Group_ErrorReason_name = map[int32]string{
		0: "GROUP_NOT_FOUND",
		1: "GROUP_DUPLICATED",
		2: "GROUP_MEMBER_NOT_FOUND",
		3: "CUSTOM_GROUP_ID_INVALID",
	}
	Group_ErrorReason_value = map[string]int32{
		"GROUP_NOT_FOUND":         0,
		"GROUP_DUPLICATED":        1,
		"GROUP_MEMBER_NOT_FOUND":  2,
		"CUSTOM_GROUP_ID_INVALID": 3,
	}
)

func (x Group_ErrorReason) Enum() *Group_ErrorReason {
	p := new(Group_ErrorReason)
	*p = x
	return p
}

func (x Group_ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Group_ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_rockim_shared_reasons_group_proto_enumTypes[0].Descriptor()
}

func (Group_ErrorReason) Type() protoreflect.EnumType {
	return &file_rockim_shared_reasons_group_proto_enumTypes[0]
}

func (x Group_ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Group_ErrorReason.Descriptor instead.
func (Group_ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_rockim_shared_reasons_group_proto_rawDescGZIP(), []int{0, 0}
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_shared_reasons_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_shared_reasons_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_rockim_shared_reasons_group_proto_rawDescGZIP(), []int{0}
}

var File_rockim_shared_reasons_group_proto protoreflect.FileDescriptor

var file_rockim_shared_reasons_group_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x7a, 0x0a, 0x05, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x22, 0x71, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x5f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1a, 0x0a,
	0x16, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x55, 0x53,
	0x54, 0x4f, 0x4d, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x49, 0x44, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x03, 0x42, 0x5d, 0x0a, 0x18, 0x63, 0x6e, 0x2e, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x73, 0x50, 0x01, 0x5a, 0x27, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0xaa, 0x02, 0x15,
	0x52, 0x6f, 0x63, 0x6b, 0x49, 0x4d, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_shared_reasons_group_proto_rawDescOnce sync.Once
	file_rockim_shared_reasons_group_proto_rawDescData = file_rockim_shared_reasons_group_proto_rawDesc
)

func file_rockim_shared_reasons_group_proto_rawDescGZIP() []byte {
	file_rockim_shared_reasons_group_proto_rawDescOnce.Do(func() {
		file_rockim_shared_reasons_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_shared_reasons_group_proto_rawDescData)
	})
	return file_rockim_shared_reasons_group_proto_rawDescData
}

var file_rockim_shared_reasons_group_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rockim_shared_reasons_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rockim_shared_reasons_group_proto_goTypes = []interface{}{
	(Group_ErrorReason)(0), // 0: rockim.shared.reasons.Group.ErrorReason
	(*Group)(nil),          // 1: rockim.shared.reasons.Group
}
var file_rockim_shared_reasons_group_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rockim_shared_reasons_group_proto_init() }
func file_rockim_shared_reasons_group_proto_init() {
	if File_rockim_shared_reasons_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_shared_reasons_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
			RawDescriptor: file_rockim_shared_reasons_group_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rockim_shared_reasons_group_proto_goTypes,
		DependencyIndexes: file_rockim_shared_reasons_group_proto_depIdxs,
		EnumInfos:         file_rockim_shared_reasons_group_proto_enumTypes,
		MessageInfos:      file_rockim_shared_reasons_group_proto_msgTypes,
	}.Build()
	File_rockim_shared_reasons_group_proto = out.File
	file_rockim_shared_reasons_group_proto_rawDesc = nil
	file_rockim_shared_reasons_group_proto_goTypes = nil
	file_rockim_shared_reasons_group_proto_depIdxs = nil
}
