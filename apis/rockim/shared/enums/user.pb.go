// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/shared/enums/user.proto

package enums

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

// Status 用户状态
type User_Status int32

const (
	// 无效
	User_USER_STATUS_INVALID User_Status = 0
	// 正常
	User_USER_STATUS_NORMAL User_Status = 1
)

// Enum value maps for User_Status.
var (
	User_Status_name = map[int32]string{
		0: "USER_STATUS_INVALID",
		1: "USER_STATUS_NORMAL",
	}
	User_Status_value = map[string]int32{
		"USER_STATUS_INVALID": 0,
		"USER_STATUS_NORMAL":  1,
	}
)

func (x User_Status) Enum() *User_Status {
	p := new(User_Status)
	*p = x
	return p
}

func (x User_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_rockim_shared_enums_user_proto_enumTypes[0].Descriptor()
}

func (User_Status) Type() protoreflect.EnumType {
	return &file_rockim_shared_enums_user_proto_enumTypes[0]
}

func (x User_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_Status.Descriptor instead.
func (User_Status) EnumDescriptor() ([]byte, []int) {
	return file_rockim_shared_enums_user_proto_rawDescGZIP(), []int{0, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_shared_enums_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_shared_enums_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_rockim_shared_enums_user_proto_rawDescGZIP(), []int{0}
}

var File_rockim_shared_enums_user_proto protoreflect.FileDescriptor

var file_rockim_shared_enums_user_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x41, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x39, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x16, 0x0a, 0x12, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x42, 0x41, 0x0a, 0x16, 0x63, 0x6e, 0x2e, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x65, 0x6e, 0x6d,
	0x75, 0x73, 0x50, 0x01, 0x5a, 0x25, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rockim_shared_enums_user_proto_rawDescOnce sync.Once
	file_rockim_shared_enums_user_proto_rawDescData = file_rockim_shared_enums_user_proto_rawDesc
)

func file_rockim_shared_enums_user_proto_rawDescGZIP() []byte {
	file_rockim_shared_enums_user_proto_rawDescOnce.Do(func() {
		file_rockim_shared_enums_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_shared_enums_user_proto_rawDescData)
	})
	return file_rockim_shared_enums_user_proto_rawDescData
}

var file_rockim_shared_enums_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rockim_shared_enums_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rockim_shared_enums_user_proto_goTypes = []interface{}{
	(User_Status)(0), // 0: rockim.shared.enums.User.Status
	(*User)(nil),     // 1: rockim.shared.enums.User
}
var file_rockim_shared_enums_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rockim_shared_enums_user_proto_init() }
func file_rockim_shared_enums_user_proto_init() {
	if File_rockim_shared_enums_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_shared_enums_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_rockim_shared_enums_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rockim_shared_enums_user_proto_goTypes,
		DependencyIndexes: file_rockim_shared_enums_user_proto_depIdxs,
		EnumInfos:         file_rockim_shared_enums_user_proto_enumTypes,
		MessageInfos:      file_rockim_shared_enums_user_proto_msgTypes,
	}.Build()
	File_rockim_shared_enums_user_proto = out.File
	file_rockim_shared_enums_user_proto_rawDesc = nil
	file_rockim_shared_enums_user_proto_goTypes = nil
	file_rockim_shared_enums_user_proto_depIdxs = nil
}
