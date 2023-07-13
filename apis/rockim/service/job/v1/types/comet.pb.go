// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/service/job/v1/types/comet.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	types "rockimserver/apis/rockim/service/comet/v1/types"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 目标类型
type CometMessage_Target_TargetType int32

const (
	CometMessage_Target_CHANNEL CometMessage_Target_TargetType = 0
	CometMessage_Target_ROOM    CometMessage_Target_TargetType = 1
)

// Enum value maps for CometMessage_Target_TargetType.
var (
	CometMessage_Target_TargetType_name = map[int32]string{
		0: "CHANNEL",
		1: "ROOM",
	}
	CometMessage_Target_TargetType_value = map[string]int32{
		"CHANNEL": 0,
		"ROOM":    1,
	}
)

func (x CometMessage_Target_TargetType) Enum() *CometMessage_Target_TargetType {
	p := new(CometMessage_Target_TargetType)
	*p = x
	return p
}

func (x CometMessage_Target_TargetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CometMessage_Target_TargetType) Descriptor() protoreflect.EnumDescriptor {
	return file_rockim_service_job_v1_types_comet_proto_enumTypes[0].Descriptor()
}

func (CometMessage_Target_TargetType) Type() protoreflect.EnumType {
	return &file_rockim_service_job_v1_types_comet_proto_enumTypes[0]
}

func (x CometMessage_Target_TargetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CometMessage_Target_TargetType.Descriptor instead.
func (CometMessage_Target_TargetType) EnumDescriptor() ([]byte, []int) {
	return file_rockim_service_job_v1_types_comet_proto_rawDescGZIP(), []int{0, 0, 0}
}

// CometMessage 长连接消息
type CometMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 应用ID
	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// 目标
	Target *CometMessage_Target `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	// 数据内容
	Message *types.Message `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CometMessage) Reset() {
	*x = CometMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_job_v1_types_comet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CometMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CometMessage) ProtoMessage() {}

func (x *CometMessage) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_job_v1_types_comet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CometMessage.ProtoReflect.Descriptor instead.
func (*CometMessage) Descriptor() ([]byte, []int) {
	return file_rockim_service_job_v1_types_comet_proto_rawDescGZIP(), []int{0}
}

func (x *CometMessage) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CometMessage) GetTarget() *CometMessage_Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *CometMessage) GetMessage() *types.Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type CometMessage_Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 目标类型
	TargetType CometMessage_Target_TargetType `protobuf:"varint,1,opt,name=target_type,json=targetType,proto3,enum=rockim.service.job.v1.types.CometMessage_Target_TargetType" json:"target_type,omitempty"`
	// 房间
	Room *types.Room `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	// 连接所在服务器ID
	ServerId string `protobuf:"bytes,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// 连接ID
	Channels []string `protobuf:"bytes,4,rep,name=channels,proto3" json:"channels,omitempty"`
}

func (x *CometMessage_Target) Reset() {
	*x = CometMessage_Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_job_v1_types_comet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CometMessage_Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CometMessage_Target) ProtoMessage() {}

func (x *CometMessage_Target) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_job_v1_types_comet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CometMessage_Target.ProtoReflect.Descriptor instead.
func (*CometMessage_Target) Descriptor() ([]byte, []int) {
	return file_rockim_service_job_v1_types_comet_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CometMessage_Target) GetTargetType() CometMessage_Target_TargetType {
	if x != nil {
		return x.TargetType
	}
	return CometMessage_Target_CHANNEL
}

func (x *CometMessage_Target) GetRoom() *types.Room {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *CometMessage_Target) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *CometMessage_Target) GetChannels() []string {
	if x != nil {
		return x.Channels
	}
	return nil
}

var File_rockim_service_job_v1_types_comet_proto protoreflect.FileDescriptor

var file_rockim_service_job_v1_types_comet_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x29, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb9, 0x03, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x48, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6d, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x40, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0xfd, 0x01,
	0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x5c, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x65,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x22, 0x23, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45,
	0x4c, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x01, 0x42, 0x35, 0x5a,
	0x33, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_service_job_v1_types_comet_proto_rawDescOnce sync.Once
	file_rockim_service_job_v1_types_comet_proto_rawDescData = file_rockim_service_job_v1_types_comet_proto_rawDesc
)

func file_rockim_service_job_v1_types_comet_proto_rawDescGZIP() []byte {
	file_rockim_service_job_v1_types_comet_proto_rawDescOnce.Do(func() {
		file_rockim_service_job_v1_types_comet_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_service_job_v1_types_comet_proto_rawDescData)
	})
	return file_rockim_service_job_v1_types_comet_proto_rawDescData
}

var file_rockim_service_job_v1_types_comet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rockim_service_job_v1_types_comet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rockim_service_job_v1_types_comet_proto_goTypes = []interface{}{
	(CometMessage_Target_TargetType)(0), // 0: rockim.service.job.v1.types.CometMessage.Target.TargetType
	(*CometMessage)(nil),                // 1: rockim.service.job.v1.types.CometMessage
	(*CometMessage_Target)(nil),         // 2: rockim.service.job.v1.types.CometMessage.Target
	(*types.Message)(nil),               // 3: rockim.service.comet.v1.types.Message
	(*types.Room)(nil),                  // 4: rockim.service.comet.v1.types.Room
}
var file_rockim_service_job_v1_types_comet_proto_depIdxs = []int32{
	2, // 0: rockim.service.job.v1.types.CometMessage.target:type_name -> rockim.service.job.v1.types.CometMessage.Target
	3, // 1: rockim.service.job.v1.types.CometMessage.message:type_name -> rockim.service.comet.v1.types.Message
	0, // 2: rockim.service.job.v1.types.CometMessage.Target.target_type:type_name -> rockim.service.job.v1.types.CometMessage.Target.TargetType
	4, // 3: rockim.service.job.v1.types.CometMessage.Target.room:type_name -> rockim.service.comet.v1.types.Room
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rockim_service_job_v1_types_comet_proto_init() }
func file_rockim_service_job_v1_types_comet_proto_init() {
	if File_rockim_service_job_v1_types_comet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_service_job_v1_types_comet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CometMessage); i {
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
		file_rockim_service_job_v1_types_comet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CometMessage_Target); i {
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
			RawDescriptor: file_rockim_service_job_v1_types_comet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rockim_service_job_v1_types_comet_proto_goTypes,
		DependencyIndexes: file_rockim_service_job_v1_types_comet_proto_depIdxs,
		EnumInfos:         file_rockim_service_job_v1_types_comet_proto_enumTypes,
		MessageInfos:      file_rockim_service_job_v1_types_comet_proto_msgTypes,
	}.Build()
	File_rockim_service_job_v1_types_comet_proto = out.File
	file_rockim_service_job_v1_types_comet_proto_rawDesc = nil
	file_rockim_service_job_v1_types_comet_proto_goTypes = nil
	file_rockim_service_job_v1_types_comet_proto_depIdxs = nil
}
