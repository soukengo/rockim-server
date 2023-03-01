// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/api/openapi/v1/auth.proto

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

// AuthCodeRequest 获取授权码请求
type AuthCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基本参数
	Base *APIRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 账号
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *AuthCodeRequest) Reset() {
	*x = AuthCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_openapi_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthCodeRequest) ProtoMessage() {}

func (x *AuthCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_openapi_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthCodeRequest.ProtoReflect.Descriptor instead.
func (*AuthCodeRequest) Descriptor() ([]byte, []int) {
	return file_rockim_api_openapi_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthCodeRequest) GetBase() *APIRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *AuthCodeRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// AuthCodeResponse 获取授权码响应
type AuthCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 授权码，用于客户端登录
	AuthCode string `protobuf:"bytes,1,opt,name=auth_code,json=authCode,proto3" json:"auth_code,omitempty"`
	// 失效时间（时间戳）
	ExpireTime int64 `protobuf:"varint,2,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *AuthCodeResponse) Reset() {
	*x = AuthCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_openapi_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthCodeResponse) ProtoMessage() {}

func (x *AuthCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_openapi_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthCodeResponse.ProtoReflect.Descriptor instead.
func (*AuthCodeResponse) Descriptor() ([]byte, []int) {
	return file_rockim_api_openapi_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthCodeResponse) GetAuthCode() string {
	if x != nil {
		return x.AuthCode
	}
	return ""
}

func (x *AuthCodeResponse) GetExpireTime() int64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

var File_rockim_api_openapi_v1_auth_proto protoreflect.FileDescriptor

var file_rockim_api_openapi_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x10,
	0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x8f,
	0x01, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x41, 0x50, 0x49, 0x12, 0x83, 0x01, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x2e,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x01, 0x2a,
	0x42, 0x2c, 0x5a, 0x2a, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_api_openapi_v1_auth_proto_rawDescOnce sync.Once
	file_rockim_api_openapi_v1_auth_proto_rawDescData = file_rockim_api_openapi_v1_auth_proto_rawDesc
)

func file_rockim_api_openapi_v1_auth_proto_rawDescGZIP() []byte {
	file_rockim_api_openapi_v1_auth_proto_rawDescOnce.Do(func() {
		file_rockim_api_openapi_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_api_openapi_v1_auth_proto_rawDescData)
	})
	return file_rockim_api_openapi_v1_auth_proto_rawDescData
}

var file_rockim_api_openapi_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rockim_api_openapi_v1_auth_proto_goTypes = []interface{}{
	(*AuthCodeRequest)(nil),  // 0: rockim.api.openapi.v1.AuthCodeRequest
	(*AuthCodeResponse)(nil), // 1: rockim.api.openapi.v1.AuthCodeResponse
	(*APIRequest)(nil),       // 2: rockim.api.openapi.v1.APIRequest
}
var file_rockim_api_openapi_v1_auth_proto_depIdxs = []int32{
	2, // 0: rockim.api.openapi.v1.AuthCodeRequest.base:type_name -> rockim.api.openapi.v1.APIRequest
	0, // 1: rockim.api.openapi.v1.AuthAPI.CreateAuthCode:input_type -> rockim.api.openapi.v1.AuthCodeRequest
	1, // 2: rockim.api.openapi.v1.AuthAPI.CreateAuthCode:output_type -> rockim.api.openapi.v1.AuthCodeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rockim_api_openapi_v1_auth_proto_init() }
func file_rockim_api_openapi_v1_auth_proto_init() {
	if File_rockim_api_openapi_v1_auth_proto != nil {
		return
	}
	file_rockim_api_openapi_v1_openapi_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rockim_api_openapi_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthCodeRequest); i {
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
		file_rockim_api_openapi_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthCodeResponse); i {
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
			RawDescriptor: file_rockim_api_openapi_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rockim_api_openapi_v1_auth_proto_goTypes,
		DependencyIndexes: file_rockim_api_openapi_v1_auth_proto_depIdxs,
		MessageInfos:      file_rockim_api_openapi_v1_auth_proto_msgTypes,
	}.Build()
	File_rockim_api_openapi_v1_auth_proto = out.File
	file_rockim_api_openapi_v1_auth_proto_rawDesc = nil
	file_rockim_api_openapi_v1_auth_proto_goTypes = nil
	file_rockim_api_openapi_v1_auth_proto_depIdxs = nil
}
func (x *AuthCodeRequest) SetBase(base *APIRequest) {
	if x == nil {
		return
	}
	x.Base = base
}

func (x *AuthCodeRequest) SetAccount(account string) {
	if x == nil {
		return
	}
	x.Account = account
}

func (x *AuthCodeResponse) SetAuthCode(authCode string) {
	if x == nil {
		return
	}
	x.AuthCode = authCode
}

func (x *AuthCodeResponse) SetExpireTime(expireTime int64) {
	if x == nil {
		return
	}
	x.ExpireTime = expireTime
}