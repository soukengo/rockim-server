// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/api/openapi/v1/user.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// UserRegisterRequest 注册用户请求
type UserRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基本参数
	Base *APIRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 用户账户名，由接入方指定
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	// 用户名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 头像地址
	AvatarUrl string `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	// 用户自定义字段
	Fields map[string]string `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UserRegisterRequest) Reset() {
	*x = UserRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_openapi_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterRequest) ProtoMessage() {}

func (x *UserRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_openapi_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterRequest.ProtoReflect.Descriptor instead.
func (*UserRegisterRequest) Descriptor() ([]byte, []int) {
	return file_rockim_api_openapi_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegisterRequest) GetBase() *APIRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UserRegisterRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserRegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserRegisterRequest) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *UserRegisterRequest) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

// UserRegisterResponse 注册用户响应
type UserRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *UserRegisterResponse) Reset() {
	*x = UserRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_openapi_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterResponse) ProtoMessage() {}

func (x *UserRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_openapi_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterResponse.ProtoReflect.Descriptor instead.
func (*UserRegisterResponse) Descriptor() ([]byte, []int) {
	return file_rockim_api_openapi_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserRegisterResponse) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

var File_rockim_api_openapi_v1_user_proto protoreflect.FileDescriptor

var file_rockim_api_openapi_v1_user_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x03, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1d, 0xfa, 0x42, 0x1a, 0x72, 0x18, 0x32, 0x16, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a,
	0x30, 0x2d, 0x39, 0x5f, 0x2e, 0x2d, 0x5d, 0x7b, 0x31, 0x2c, 0x36, 0x34, 0x7d, 0x24, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x18, 0x40, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x10, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x02, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x88, 0x01,
	0x01, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x71, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x21, 0xfa, 0x42, 0x05, 0x9a, 0x01, 0x02, 0x10, 0x64, 0xfa, 0x42,
	0x09, 0x9a, 0x01, 0x06, 0x22, 0x04, 0x72, 0x02, 0x18, 0x40, 0xfa, 0x42, 0x0a, 0x9a, 0x01, 0x07,
	0x2a, 0x05, 0x72, 0x03, 0x18, 0x80, 0x04, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a,
	0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x28, 0x0a, 0x14, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x32, 0x95, 0x01, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x41, 0x50, 0x49,
	0x12, 0x89, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x2e,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x42, 0x2c, 0x5a, 0x2a,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_rockim_api_openapi_v1_user_proto_rawDescOnce sync.Once
	file_rockim_api_openapi_v1_user_proto_rawDescData = file_rockim_api_openapi_v1_user_proto_rawDesc
)

func file_rockim_api_openapi_v1_user_proto_rawDescGZIP() []byte {
	file_rockim_api_openapi_v1_user_proto_rawDescOnce.Do(func() {
		file_rockim_api_openapi_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_api_openapi_v1_user_proto_rawDescData)
	})
	return file_rockim_api_openapi_v1_user_proto_rawDescData
}

var file_rockim_api_openapi_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rockim_api_openapi_v1_user_proto_goTypes = []interface{}{
	(*UserRegisterRequest)(nil),  // 0: rockim.api.openapi.v1.UserRegisterRequest
	(*UserRegisterResponse)(nil), // 1: rockim.api.openapi.v1.UserRegisterResponse
	nil,                          // 2: rockim.api.openapi.v1.UserRegisterRequest.FieldsEntry
	(*APIRequest)(nil),           // 3: rockim.api.openapi.v1.APIRequest
}
var file_rockim_api_openapi_v1_user_proto_depIdxs = []int32{
	3, // 0: rockim.api.openapi.v1.UserRegisterRequest.base:type_name -> rockim.api.openapi.v1.APIRequest
	2, // 1: rockim.api.openapi.v1.UserRegisterRequest.fields:type_name -> rockim.api.openapi.v1.UserRegisterRequest.FieldsEntry
	0, // 2: rockim.api.openapi.v1.UserAPI.Register:input_type -> rockim.api.openapi.v1.UserRegisterRequest
	1, // 3: rockim.api.openapi.v1.UserAPI.Register:output_type -> rockim.api.openapi.v1.UserRegisterResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rockim_api_openapi_v1_user_proto_init() }
func file_rockim_api_openapi_v1_user_proto_init() {
	if File_rockim_api_openapi_v1_user_proto != nil {
		return
	}
	file_rockim_api_openapi_v1_openapi_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rockim_api_openapi_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterRequest); i {
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
		file_rockim_api_openapi_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterResponse); i {
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
			RawDescriptor: file_rockim_api_openapi_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rockim_api_openapi_v1_user_proto_goTypes,
		DependencyIndexes: file_rockim_api_openapi_v1_user_proto_depIdxs,
		MessageInfos:      file_rockim_api_openapi_v1_user_proto_msgTypes,
	}.Build()
	File_rockim_api_openapi_v1_user_proto = out.File
	file_rockim_api_openapi_v1_user_proto_rawDesc = nil
	file_rockim_api_openapi_v1_user_proto_goTypes = nil
	file_rockim_api_openapi_v1_user_proto_depIdxs = nil
}
func (x *UserRegisterRequest) SetBase(base *APIRequest) {
	if x == nil {
		return
	}
	x.Base = base
}

func (x *UserRegisterRequest) SetAccount(account string) {
	if x == nil {
		return
	}
	x.Account = account
}

func (x *UserRegisterRequest) SetName(name string) {
	if x == nil {
		return
	}
	x.Name = name
}

func (x *UserRegisterRequest) SetAvatarUrl(avatarUrl string) {
	if x == nil {
		return
	}
	x.AvatarUrl = avatarUrl
}

func (x *UserRegisterRequest) SetFields(fields map[string]string) {
	if x == nil {
		return
	}
	x.Fields = fields
}

func (x *UserRegisterResponse) SetUid(uid string) {
	if x == nil {
		return
	}
	x.Uid = uid
}
