// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/shared/reasons/comet.proto

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

type Comet_ErrorReason int32

const (
	// 不支持的控制类型
	Comet_CONTROL_NOT_SUPPORTED Comet_ErrorReason = 0
	// 控制数据无效
	Comet_CONTROL_DATA_INVALID Comet_ErrorReason = 1
)

// Enum value maps for Comet_ErrorReason.
var (
	Comet_ErrorReason_name = map[int32]string{
		0: "CONTROL_NOT_SUPPORTED",
		1: "CONTROL_DATA_INVALID",
	}
	Comet_ErrorReason_value = map[string]int32{
		"CONTROL_NOT_SUPPORTED": 0,
		"CONTROL_DATA_INVALID":  1,
	}
)

func (x Comet_ErrorReason) Enum() *Comet_ErrorReason {
	p := new(Comet_ErrorReason)
	*p = x
	return p
}

func (x Comet_ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Comet_ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_rockim_shared_reasons_comet_proto_enumTypes[0].Descriptor()
}

func (Comet_ErrorReason) Type() protoreflect.EnumType {
	return &file_rockim_shared_reasons_comet_proto_enumTypes[0]
}

func (x Comet_ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Comet_ErrorReason.Descriptor instead.
func (Comet_ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_rockim_shared_reasons_comet_proto_rawDescGZIP(), []int{0, 0}
}

type Comet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Comet) Reset() {
	*x = Comet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_shared_reasons_comet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comet) ProtoMessage() {}

func (x *Comet) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_shared_reasons_comet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comet.ProtoReflect.Descriptor instead.
func (*Comet) Descriptor() ([]byte, []int) {
	return file_rockim_shared_reasons_comet_proto_rawDescGZIP(), []int{0}
}

var File_rockim_shared_reasons_comet_proto protoreflect.FileDescriptor

var file_rockim_shared_reasons_comet_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x4b, 0x0a, 0x05, 0x43, 0x6f,
	0x6d, 0x65, 0x74, 0x22, 0x42, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a,
	0x14, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x01, 0x42, 0x5d, 0x0a, 0x18, 0x63, 0x6e, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x73, 0x50, 0x01, 0x5a, 0x27, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0xaa, 0x02,
	0x15, 0x52, 0x6f, 0x63, 0x6b, 0x49, 0x4d, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_shared_reasons_comet_proto_rawDescOnce sync.Once
	file_rockim_shared_reasons_comet_proto_rawDescData = file_rockim_shared_reasons_comet_proto_rawDesc
)

func file_rockim_shared_reasons_comet_proto_rawDescGZIP() []byte {
	file_rockim_shared_reasons_comet_proto_rawDescOnce.Do(func() {
		file_rockim_shared_reasons_comet_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_shared_reasons_comet_proto_rawDescData)
	})
	return file_rockim_shared_reasons_comet_proto_rawDescData
}

var file_rockim_shared_reasons_comet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rockim_shared_reasons_comet_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rockim_shared_reasons_comet_proto_goTypes = []interface{}{
	(Comet_ErrorReason)(0), // 0: rockim.shared.reasons.Comet.ErrorReason
	(*Comet)(nil),          // 1: rockim.shared.reasons.Comet
}
var file_rockim_shared_reasons_comet_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rockim_shared_reasons_comet_proto_init() }
func file_rockim_shared_reasons_comet_proto_init() {
	if File_rockim_shared_reasons_comet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_shared_reasons_comet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comet); i {
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
			RawDescriptor: file_rockim_shared_reasons_comet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rockim_shared_reasons_comet_proto_goTypes,
		DependencyIndexes: file_rockim_shared_reasons_comet_proto_depIdxs,
		EnumInfos:         file_rockim_shared_reasons_comet_proto_enumTypes,
		MessageInfos:      file_rockim_shared_reasons_comet_proto_msgTypes,
	}.Build()
	File_rockim_shared_reasons_comet_proto = out.File
	file_rockim_shared_reasons_comet_proto_rawDesc = nil
	file_rockim_shared_reasons_comet_proto_goTypes = nil
	file_rockim_shared_reasons_comet_proto_depIdxs = nil
}
