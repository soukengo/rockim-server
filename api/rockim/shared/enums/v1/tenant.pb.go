// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: api/rockim/shared/enums/v1/tenant.proto

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

// TenantStatus 租户状态
type TenantStatus int32

const (
	// 无效
	TenantStatus_TENANT_STATUS_INVALID TenantStatus = 0
	// 待审核
	TenantStatus_TENANT_STATUS_AUDITING TenantStatus = 1
	// 启用
	TenantStatus_TENANT_STATUS_ACTIVE TenantStatus = 2
)

// Enum value maps for TenantStatus.
var (
	TenantStatus_name = map[int32]string{
		0: "TENANT_STATUS_INVALID",
		1: "TENANT_STATUS_AUDITING",
		2: "TENANT_STATUS_ACTIVE",
	}
	TenantStatus_value = map[string]int32{
		"TENANT_STATUS_INVALID":  0,
		"TENANT_STATUS_AUDITING": 1,
		"TENANT_STATUS_ACTIVE":   2,
	}
)

func (x TenantStatus) Enum() *TenantStatus {
	p := new(TenantStatus)
	*p = x
	return p
}

func (x TenantStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TenantStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_rockim_shared_enums_v1_tenant_proto_enumTypes[0].Descriptor()
}

func (TenantStatus) Type() protoreflect.EnumType {
	return &file_api_rockim_shared_enums_v1_tenant_proto_enumTypes[0]
}

func (x TenantStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TenantStatus.Descriptor instead.
func (TenantStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_rockim_shared_enums_v1_tenant_proto_rawDescGZIP(), []int{0}
}

var File_api_rockim_shared_enums_v1_tenant_proto protoreflect.FileDescriptor

var file_api_rockim_shared_enums_v1_tenant_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2a, 0x5f, 0x0a, 0x0c, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16,
	0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x55,
	0x44, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x45, 0x4e, 0x41,
	0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x02, 0x42, 0x23, 0x5a, 0x21, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_rockim_shared_enums_v1_tenant_proto_rawDescOnce sync.Once
	file_api_rockim_shared_enums_v1_tenant_proto_rawDescData = file_api_rockim_shared_enums_v1_tenant_proto_rawDesc
)

func file_api_rockim_shared_enums_v1_tenant_proto_rawDescGZIP() []byte {
	file_api_rockim_shared_enums_v1_tenant_proto_rawDescOnce.Do(func() {
		file_api_rockim_shared_enums_v1_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rockim_shared_enums_v1_tenant_proto_rawDescData)
	})
	return file_api_rockim_shared_enums_v1_tenant_proto_rawDescData
}

var file_api_rockim_shared_enums_v1_tenant_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_rockim_shared_enums_v1_tenant_proto_goTypes = []interface{}{
	(TenantStatus)(0), // 0: rockim.shared.enums.v1.TenantStatus
}
var file_api_rockim_shared_enums_v1_tenant_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_rockim_shared_enums_v1_tenant_proto_init() }
func file_api_rockim_shared_enums_v1_tenant_proto_init() {
	if File_api_rockim_shared_enums_v1_tenant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_rockim_shared_enums_v1_tenant_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_rockim_shared_enums_v1_tenant_proto_goTypes,
		DependencyIndexes: file_api_rockim_shared_enums_v1_tenant_proto_depIdxs,
		EnumInfos:         file_api_rockim_shared_enums_v1_tenant_proto_enumTypes,
	}.Build()
	File_api_rockim_shared_enums_v1_tenant_proto = out.File
	file_api_rockim_shared_enums_v1_tenant_proto_rawDesc = nil
	file_api_rockim_shared_enums_v1_tenant_proto_goTypes = nil
	file_api_rockim_shared_enums_v1_tenant_proto_depIdxs = nil
}
