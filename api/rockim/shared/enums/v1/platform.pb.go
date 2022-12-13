// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: api/rockim/shared/enums/v1/platform.proto

package v1

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
type PlatResourceCategory int32

const (
	// 菜单
	PlatResourceCategory_ResourceCategoryMenu PlatResourceCategory = 0
	// 页面
	PlatResourceCategory_ResourceCategoryPage PlatResourceCategory = 1
	// 功能
	PlatResourceCategory_ResourceCategoryFunction PlatResourceCategory = 2
)

// Enum value maps for PlatResourceCategory.
var (
	PlatResourceCategory_name = map[int32]string{
		0: "ResourceCategoryMenu",
		1: "ResourceCategoryPage",
		2: "ResourceCategoryFunction",
	}
	PlatResourceCategory_value = map[string]int32{
		"ResourceCategoryMenu":     0,
		"ResourceCategoryPage":     1,
		"ResourceCategoryFunction": 2,
	}
)

func (x PlatResourceCategory) Enum() *PlatResourceCategory {
	p := new(PlatResourceCategory)
	*p = x
	return p
}

func (x PlatResourceCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlatResourceCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_api_rockim_shared_enums_v1_platform_proto_enumTypes[0].Descriptor()
}

func (PlatResourceCategory) Type() protoreflect.EnumType {
	return &file_api_rockim_shared_enums_v1_platform_proto_enumTypes[0]
}

func (x PlatResourceCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlatResourceCategory.Descriptor instead.
func (PlatResourceCategory) EnumDescriptor() ([]byte, []int) {
	return file_api_rockim_shared_enums_v1_platform_proto_rawDescGZIP(), []int{0}
}

var File_api_rockim_shared_enums_v1_platform_proto protoreflect.FileDescriptor

var file_api_rockim_shared_enums_v1_platform_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x2a, 0x68, 0x0a, 0x14, 0x50, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d,
	0x65, 0x6e, 0x75, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x10, 0x01, 0x12,
	0x1c, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x02, 0x42, 0x23, 0x5a,
	0x21, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_rockim_shared_enums_v1_platform_proto_rawDescOnce sync.Once
	file_api_rockim_shared_enums_v1_platform_proto_rawDescData = file_api_rockim_shared_enums_v1_platform_proto_rawDesc
)

func file_api_rockim_shared_enums_v1_platform_proto_rawDescGZIP() []byte {
	file_api_rockim_shared_enums_v1_platform_proto_rawDescOnce.Do(func() {
		file_api_rockim_shared_enums_v1_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rockim_shared_enums_v1_platform_proto_rawDescData)
	})
	return file_api_rockim_shared_enums_v1_platform_proto_rawDescData
}

var file_api_rockim_shared_enums_v1_platform_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_rockim_shared_enums_v1_platform_proto_goTypes = []interface{}{
	(PlatResourceCategory)(0), // 0: rockim.shared.enums.v1.PlatResourceCategory
}
var file_api_rockim_shared_enums_v1_platform_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_rockim_shared_enums_v1_platform_proto_init() }
func file_api_rockim_shared_enums_v1_platform_proto_init() {
	if File_api_rockim_shared_enums_v1_platform_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_rockim_shared_enums_v1_platform_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_rockim_shared_enums_v1_platform_proto_goTypes,
		DependencyIndexes: file_api_rockim_shared_enums_v1_platform_proto_depIdxs,
		EnumInfos:         file_api_rockim_shared_enums_v1_platform_proto_enumTypes,
	}.Build()
	File_api_rockim_shared_enums_v1_platform_proto = out.File
	file_api_rockim_shared_enums_v1_platform_proto_rawDesc = nil
	file_api_rockim_shared_enums_v1_platform_proto_goTypes = nil
	file_api_rockim_shared_enums_v1_platform_proto_depIdxs = nil
}