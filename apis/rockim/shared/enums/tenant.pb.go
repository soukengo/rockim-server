// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/shared/enums/tenant.proto

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

// Status 商户状态
type Tenant_Status int32

const (
	// 无效
	Tenant_TENANT_STATUS_INVALID Tenant_Status = 0
	// 待审核
	Tenant_TENANT_STATUS_AUDITING Tenant_Status = 1
	// 启用
	Tenant_TENANT_STATUS_ACTIVE Tenant_Status = 2
)

// Enum value maps for Tenant_Status.
var (
	Tenant_Status_name = map[int32]string{
		0: "TENANT_STATUS_INVALID",
		1: "TENANT_STATUS_AUDITING",
		2: "TENANT_STATUS_ACTIVE",
	}
	Tenant_Status_value = map[string]int32{
		"TENANT_STATUS_INVALID":  0,
		"TENANT_STATUS_AUDITING": 1,
		"TENANT_STATUS_ACTIVE":   2,
	}
)

func (x Tenant_Status) Enum() *Tenant_Status {
	p := new(Tenant_Status)
	*p = x
	return p
}

func (x Tenant_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Tenant_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_rockim_shared_enums_tenant_proto_enumTypes[0].Descriptor()
}

func (Tenant_Status) Type() protoreflect.EnumType {
	return &file_rockim_shared_enums_tenant_proto_enumTypes[0]
}

func (x Tenant_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Tenant_Status.Descriptor instead.
func (Tenant_Status) EnumDescriptor() ([]byte, []int) {
	return file_rockim_shared_enums_tenant_proto_rawDescGZIP(), []int{0, 0}
}

type Tenant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_shared_enums_tenant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_shared_enums_tenant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_rockim_shared_enums_tenant_proto_rawDescGZIP(), []int{0}
}

var File_rockim_shared_enums_tenant_proto protoreflect.FileDescriptor

var file_rockim_shared_enums_tenant_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x63, 0x0a, 0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x22, 0x59, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x54,
	0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x42, 0x57, 0x0a, 0x16,
	0x63, 0x6e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x65, 0x6e, 0x6d, 0x75, 0x73, 0x50, 0x01, 0x5a, 0x25, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xaa,
	0x02, 0x13, 0x52, 0x6f, 0x63, 0x6b, 0x49, 0x4d, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_shared_enums_tenant_proto_rawDescOnce sync.Once
	file_rockim_shared_enums_tenant_proto_rawDescData = file_rockim_shared_enums_tenant_proto_rawDesc
)

func file_rockim_shared_enums_tenant_proto_rawDescGZIP() []byte {
	file_rockim_shared_enums_tenant_proto_rawDescOnce.Do(func() {
		file_rockim_shared_enums_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_shared_enums_tenant_proto_rawDescData)
	})
	return file_rockim_shared_enums_tenant_proto_rawDescData
}

var file_rockim_shared_enums_tenant_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rockim_shared_enums_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rockim_shared_enums_tenant_proto_goTypes = []interface{}{
	(Tenant_Status)(0), // 0: rockim.shared.enums.Tenant.Status
	(*Tenant)(nil),     // 1: rockim.shared.enums.Tenant
}
var file_rockim_shared_enums_tenant_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rockim_shared_enums_tenant_proto_init() }
func file_rockim_shared_enums_tenant_proto_init() {
	if File_rockim_shared_enums_tenant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_shared_enums_tenant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tenant); i {
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
			RawDescriptor: file_rockim_shared_enums_tenant_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rockim_shared_enums_tenant_proto_goTypes,
		DependencyIndexes: file_rockim_shared_enums_tenant_proto_depIdxs,
		EnumInfos:         file_rockim_shared_enums_tenant_proto_enumTypes,
		MessageInfos:      file_rockim_shared_enums_tenant_proto_msgTypes,
	}.Build()
	File_rockim_shared_enums_tenant_proto = out.File
	file_rockim_shared_enums_tenant_proto_rawDesc = nil
	file_rockim_shared_enums_tenant_proto_goTypes = nil
	file_rockim_shared_enums_tenant_proto_depIdxs = nil
}
