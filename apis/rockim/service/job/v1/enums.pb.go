// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/service/job/v1/enums.proto

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

// TaskType 任务类型
type TaskType int32

const (
	// 长连接消息投递
	TaskType_COMET_DISPATCH TaskType = 0
)

// Enum value maps for TaskType.
var (
	TaskType_name = map[int32]string{
		0: "COMET_DISPATCH",
	}
	TaskType_value = map[string]int32{
		"COMET_DISPATCH": 0,
	}
)

func (x TaskType) Enum() *TaskType {
	p := new(TaskType)
	*p = x
	return p
}

func (x TaskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskType) Descriptor() protoreflect.EnumDescriptor {
	return file_rockim_service_job_v1_enums_proto_enumTypes[0].Descriptor()
}

func (TaskType) Type() protoreflect.EnumType {
	return &file_rockim_service_job_v1_enums_proto_enumTypes[0]
}

func (x TaskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskType.Descriptor instead.
func (TaskType) EnumDescriptor() ([]byte, []int) {
	return file_rockim_service_job_v1_enums_proto_rawDescGZIP(), []int{0}
}

var File_rockim_service_job_v1_enums_proto protoreflect.FileDescriptor

var file_rockim_service_job_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2a, 0x1e, 0x0a, 0x08, 0x54, 0x61,
	0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4d, 0x45, 0x54, 0x5f,
	0x44, 0x49, 0x53, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x72, 0x6f,
	0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6a,
	0x6f, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_service_job_v1_enums_proto_rawDescOnce sync.Once
	file_rockim_service_job_v1_enums_proto_rawDescData = file_rockim_service_job_v1_enums_proto_rawDesc
)

func file_rockim_service_job_v1_enums_proto_rawDescGZIP() []byte {
	file_rockim_service_job_v1_enums_proto_rawDescOnce.Do(func() {
		file_rockim_service_job_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_service_job_v1_enums_proto_rawDescData)
	})
	return file_rockim_service_job_v1_enums_proto_rawDescData
}

var file_rockim_service_job_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rockim_service_job_v1_enums_proto_goTypes = []interface{}{
	(TaskType)(0), // 0: rockim.service.job.v1.TaskType
}
var file_rockim_service_job_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rockim_service_job_v1_enums_proto_init() }
func file_rockim_service_job_v1_enums_proto_init() {
	if File_rockim_service_job_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rockim_service_job_v1_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rockim_service_job_v1_enums_proto_goTypes,
		DependencyIndexes: file_rockim_service_job_v1_enums_proto_depIdxs,
		EnumInfos:         file_rockim_service_job_v1_enums_proto_enumTypes,
	}.Build()
	File_rockim_service_job_v1_enums_proto = out.File
	file_rockim_service_job_v1_enums_proto_rawDesc = nil
	file_rockim_service_job_v1_enums_proto_goTypes = nil
	file_rockim_service_job_v1_enums_proto_depIdxs = nil
}
