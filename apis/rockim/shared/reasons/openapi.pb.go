// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/shared/reasons/openapi.proto

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

type OpenAPI_ErrorReason int32

const (
	// 接口地址不存在
	OpenAPI_API_NOT_FOUND OpenAPI_ErrorReason = 0
	// 无效的签名
	OpenAPI_SIGN_INVALID OpenAPI_ErrorReason = 1
)

// Enum value maps for OpenAPI_ErrorReason.
var (
	OpenAPI_ErrorReason_name = map[int32]string{
		0: "API_NOT_FOUND",
		1: "SIGN_INVALID",
	}
	OpenAPI_ErrorReason_value = map[string]int32{
		"API_NOT_FOUND": 0,
		"SIGN_INVALID":  1,
	}
)

func (x OpenAPI_ErrorReason) Enum() *OpenAPI_ErrorReason {
	p := new(OpenAPI_ErrorReason)
	*p = x
	return p
}

func (x OpenAPI_ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpenAPI_ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_rockim_shared_reasons_openapi_proto_enumTypes[0].Descriptor()
}

func (OpenAPI_ErrorReason) Type() protoreflect.EnumType {
	return &file_rockim_shared_reasons_openapi_proto_enumTypes[0]
}

func (x OpenAPI_ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpenAPI_ErrorReason.Descriptor instead.
func (OpenAPI_ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_rockim_shared_reasons_openapi_proto_rawDescGZIP(), []int{0, 0}
}

type OpenAPI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OpenAPI) Reset() {
	*x = OpenAPI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_shared_reasons_openapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAPI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAPI) ProtoMessage() {}

func (x *OpenAPI) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_shared_reasons_openapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAPI.ProtoReflect.Descriptor instead.
func (*OpenAPI) Descriptor() ([]byte, []int) {
	return file_rockim_shared_reasons_openapi_proto_rawDescGZIP(), []int{0}
}

var File_rockim_shared_reasons_openapi_proto protoreflect.FileDescriptor

var file_rockim_shared_reasons_openapi_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x3d, 0x0a, 0x07,
	0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x22, 0x32, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x50, 0x49, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x49, 0x47,
	0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x01, 0x42, 0x45, 0x0a, 0x18, 0x63,
	0x6e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x50, 0x01, 0x5a, 0x27, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_shared_reasons_openapi_proto_rawDescOnce sync.Once
	file_rockim_shared_reasons_openapi_proto_rawDescData = file_rockim_shared_reasons_openapi_proto_rawDesc
)

func file_rockim_shared_reasons_openapi_proto_rawDescGZIP() []byte {
	file_rockim_shared_reasons_openapi_proto_rawDescOnce.Do(func() {
		file_rockim_shared_reasons_openapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_shared_reasons_openapi_proto_rawDescData)
	})
	return file_rockim_shared_reasons_openapi_proto_rawDescData
}

var file_rockim_shared_reasons_openapi_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rockim_shared_reasons_openapi_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rockim_shared_reasons_openapi_proto_goTypes = []interface{}{
	(OpenAPI_ErrorReason)(0), // 0: rockim.shared.reasons.OpenAPI.ErrorReason
	(*OpenAPI)(nil),          // 1: rockim.shared.reasons.OpenAPI
}
var file_rockim_shared_reasons_openapi_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rockim_shared_reasons_openapi_proto_init() }
func file_rockim_shared_reasons_openapi_proto_init() {
	if File_rockim_shared_reasons_openapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_shared_reasons_openapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAPI); i {
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
			RawDescriptor: file_rockim_shared_reasons_openapi_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rockim_shared_reasons_openapi_proto_goTypes,
		DependencyIndexes: file_rockim_shared_reasons_openapi_proto_depIdxs,
		EnumInfos:         file_rockim_shared_reasons_openapi_proto_enumTypes,
		MessageInfos:      file_rockim_shared_reasons_openapi_proto_msgTypes,
	}.Build()
	File_rockim_shared_reasons_openapi_proto = out.File
	file_rockim_shared_reasons_openapi_proto_rawDesc = nil
	file_rockim_shared_reasons_openapi_proto_goTypes = nil
	file_rockim_shared_reasons_openapi_proto_depIdxs = nil
}
