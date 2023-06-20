// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/shared/enums/admin.proto

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

// 资源类型
type Admin_ResourceCategory int32

const (
	// 无效
	Admin_RESOURCE_CATEGORY_INVALID Admin_ResourceCategory = 0
	// 菜单
	Admin_RESOURCE_CATEGORY_MENU Admin_ResourceCategory = 1
	// 页面
	Admin_RESOURCE_CATEGORY_PAGE Admin_ResourceCategory = 2
	// 功能
	Admin_RESOURCE_CATEGORY_FUNCTION Admin_ResourceCategory = 3
)

// Enum value maps for Admin_ResourceCategory.
var (
	Admin_ResourceCategory_name = map[int32]string{
		0: "RESOURCE_CATEGORY_INVALID",
		1: "RESOURCE_CATEGORY_MENU",
		2: "RESOURCE_CATEGORY_PAGE",
		3: "RESOURCE_CATEGORY_FUNCTION",
	}
	Admin_ResourceCategory_value = map[string]int32{
		"RESOURCE_CATEGORY_INVALID":  0,
		"RESOURCE_CATEGORY_MENU":     1,
		"RESOURCE_CATEGORY_PAGE":     2,
		"RESOURCE_CATEGORY_FUNCTION": 3,
	}
)

func (x Admin_ResourceCategory) Enum() *Admin_ResourceCategory {
	p := new(Admin_ResourceCategory)
	*p = x
	return p
}

func (x Admin_ResourceCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Admin_ResourceCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_rockim_shared_enums_admin_proto_enumTypes[0].Descriptor()
}

func (Admin_ResourceCategory) Type() protoreflect.EnumType {
	return &file_rockim_shared_enums_admin_proto_enumTypes[0]
}

func (x Admin_ResourceCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Admin_ResourceCategory.Descriptor instead.
func (Admin_ResourceCategory) EnumDescriptor() ([]byte, []int) {
	return file_rockim_shared_enums_admin_proto_rawDescGZIP(), []int{0, 0}
}

type Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Admin) Reset() {
	*x = Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_shared_enums_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_shared_enums_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admin.ProtoReflect.Descriptor instead.
func (*Admin) Descriptor() ([]byte, []int) {
	return file_rockim_shared_enums_admin_proto_rawDescGZIP(), []int{0}
}

var File_rockim_shared_enums_admin_proto protoreflect.FileDescriptor

var file_rockim_shared_enums_admin_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x22, 0x89, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4d, 0x45, 0x4e, 0x55, 0x10, 0x01,
	0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x43, 0x41, 0x54,
	0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52,
	0x59, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x42, 0x57, 0x0a, 0x16,
	0x63, 0x6e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x65, 0x6e, 0x6d, 0x75, 0x73, 0x50, 0x01, 0x5a, 0x25, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xaa,
	0x02, 0x13, 0x52, 0x6f, 0x63, 0x6b, 0x49, 0x4d, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_shared_enums_admin_proto_rawDescOnce sync.Once
	file_rockim_shared_enums_admin_proto_rawDescData = file_rockim_shared_enums_admin_proto_rawDesc
)

func file_rockim_shared_enums_admin_proto_rawDescGZIP() []byte {
	file_rockim_shared_enums_admin_proto_rawDescOnce.Do(func() {
		file_rockim_shared_enums_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_shared_enums_admin_proto_rawDescData)
	})
	return file_rockim_shared_enums_admin_proto_rawDescData
}

var file_rockim_shared_enums_admin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rockim_shared_enums_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rockim_shared_enums_admin_proto_goTypes = []interface{}{
	(Admin_ResourceCategory)(0), // 0: rockim.shared.enums.Admin.ResourceCategory
	(*Admin)(nil),               // 1: rockim.shared.enums.Admin
}
var file_rockim_shared_enums_admin_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rockim_shared_enums_admin_proto_init() }
func file_rockim_shared_enums_admin_proto_init() {
	if File_rockim_shared_enums_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_shared_enums_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Admin); i {
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
			RawDescriptor: file_rockim_shared_enums_admin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rockim_shared_enums_admin_proto_goTypes,
		DependencyIndexes: file_rockim_shared_enums_admin_proto_depIdxs,
		EnumInfos:         file_rockim_shared_enums_admin_proto_enumTypes,
		MessageInfos:      file_rockim_shared_enums_admin_proto_msgTypes,
	}.Build()
	File_rockim_shared_enums_admin_proto = out.File
	file_rockim_shared_enums_admin_proto_rawDesc = nil
	file_rockim_shared_enums_admin_proto_goTypes = nil
	file_rockim_shared_enums_admin_proto_depIdxs = nil
}
