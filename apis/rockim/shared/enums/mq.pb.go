// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/shared/enums/mq.proto

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

// Topic topic
type MQ_Topic int32

const (
	// 消息投递
	MQ_MESSAGE_DELIVERY MQ_Topic = 0
)

// Enum value maps for MQ_Topic.
var (
	MQ_Topic_name = map[int32]string{
		0: "MESSAGE_DELIVERY",
	}
	MQ_Topic_value = map[string]int32{
		"MESSAGE_DELIVERY": 0,
	}
)

func (x MQ_Topic) Enum() *MQ_Topic {
	p := new(MQ_Topic)
	*p = x
	return p
}

func (x MQ_Topic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MQ_Topic) Descriptor() protoreflect.EnumDescriptor {
	return file_rockim_shared_enums_mq_proto_enumTypes[0].Descriptor()
}

func (MQ_Topic) Type() protoreflect.EnumType {
	return &file_rockim_shared_enums_mq_proto_enumTypes[0]
}

func (x MQ_Topic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MQ_Topic.Descriptor instead.
func (MQ_Topic) EnumDescriptor() ([]byte, []int) {
	return file_rockim_shared_enums_mq_proto_rawDescGZIP(), []int{0, 0}
}

// MQ 消息中间件
type MQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MQ) Reset() {
	*x = MQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_shared_enums_mq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MQ) ProtoMessage() {}

func (x *MQ) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_shared_enums_mq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MQ.ProtoReflect.Descriptor instead.
func (*MQ) Descriptor() ([]byte, []int) {
	return file_rockim_shared_enums_mq_proto_rawDescGZIP(), []int{0}
}

var File_rockim_shared_enums_mq_proto protoreflect.FileDescriptor

var file_rockim_shared_enums_mq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x22, 0x23, 0x0a, 0x02, 0x4d, 0x51, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x44, 0x45,
	0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x10, 0x00, 0x42, 0x41, 0x0a, 0x16, 0x63, 0x6e, 0x2e, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x65, 0x6e, 0x6d,
	0x75, 0x73, 0x50, 0x01, 0x5a, 0x25, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rockim_shared_enums_mq_proto_rawDescOnce sync.Once
	file_rockim_shared_enums_mq_proto_rawDescData = file_rockim_shared_enums_mq_proto_rawDesc
)

func file_rockim_shared_enums_mq_proto_rawDescGZIP() []byte {
	file_rockim_shared_enums_mq_proto_rawDescOnce.Do(func() {
		file_rockim_shared_enums_mq_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_shared_enums_mq_proto_rawDescData)
	})
	return file_rockim_shared_enums_mq_proto_rawDescData
}

var file_rockim_shared_enums_mq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rockim_shared_enums_mq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rockim_shared_enums_mq_proto_goTypes = []interface{}{
	(MQ_Topic)(0), // 0: rockim.shared.enums.MQ.Topic
	(*MQ)(nil),    // 1: rockim.shared.enums.MQ
}
var file_rockim_shared_enums_mq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rockim_shared_enums_mq_proto_init() }
func file_rockim_shared_enums_mq_proto_init() {
	if File_rockim_shared_enums_mq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_shared_enums_mq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MQ); i {
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
			RawDescriptor: file_rockim_shared_enums_mq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rockim_shared_enums_mq_proto_goTypes,
		DependencyIndexes: file_rockim_shared_enums_mq_proto_depIdxs,
		EnumInfos:         file_rockim_shared_enums_mq_proto_enumTypes,
		MessageInfos:      file_rockim_shared_enums_mq_proto_msgTypes,
	}.Build()
	File_rockim_shared_enums_mq_proto = out.File
	file_rockim_shared_enums_mq_proto_rawDesc = nil
	file_rockim_shared_enums_mq_proto_goTypes = nil
	file_rockim_shared_enums_mq_proto_depIdxs = nil
}