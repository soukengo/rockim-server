// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/api/client/v1/protocol/http/auth.proto

package http

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
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基本参数
	Base *APIRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 授权码
	AuthCode string `protobuf:"bytes,2,opt,name=auth_code,json=authCode,proto3" json:"auth_code,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_protocol_http_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_protocol_http_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_protocol_http_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetBase() *APIRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *LoginRequest) GetAuthCode() string {
	if x != nil {
		return x.AuthCode
	}
	return ""
}

// LoginResponse 登录响应
type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 访问令牌
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// 失效时间（时间戳）
	ExpireTime int64 `protobuf:"varint,2,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_protocol_http_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_protocol_http_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_protocol_http_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetExpireTime() int64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

var File_rockim_api_client_v1_protocol_http_auth_proto protoreflect.FileDescriptor

var file_rockim_api_client_v1_protocol_http_auth_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x41, 0x50,
	0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x46,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x9a, 0x01, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x41,
	0x50, 0x49, 0x12, 0x8e, 0x01, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x30, 0x2e, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31,
	0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x3a, 0x01, 0x2a, 0x42, 0x60, 0x5a, 0x39, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x3b, 0x68, 0x74, 0x74, 0x70,
	0xaa, 0x02, 0x22, 0x52, 0x6f, 0x63, 0x6b, 0x49, 0x4d, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_api_client_v1_protocol_http_auth_proto_rawDescOnce sync.Once
	file_rockim_api_client_v1_protocol_http_auth_proto_rawDescData = file_rockim_api_client_v1_protocol_http_auth_proto_rawDesc
)

func file_rockim_api_client_v1_protocol_http_auth_proto_rawDescGZIP() []byte {
	file_rockim_api_client_v1_protocol_http_auth_proto_rawDescOnce.Do(func() {
		file_rockim_api_client_v1_protocol_http_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_api_client_v1_protocol_http_auth_proto_rawDescData)
	})
	return file_rockim_api_client_v1_protocol_http_auth_proto_rawDescData
}

var file_rockim_api_client_v1_protocol_http_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rockim_api_client_v1_protocol_http_auth_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),  // 0: rockim.api.client.v1.protocol.http.LoginRequest
	(*LoginResponse)(nil), // 1: rockim.api.client.v1.protocol.http.LoginResponse
	(*APIRequest)(nil),    // 2: rockim.api.client.v1.protocol.http.APIRequest
}
var file_rockim_api_client_v1_protocol_http_auth_proto_depIdxs = []int32{
	2, // 0: rockim.api.client.v1.protocol.http.LoginRequest.base:type_name -> rockim.api.client.v1.protocol.http.APIRequest
	0, // 1: rockim.api.client.v1.protocol.http.AuthAPI.Login:input_type -> rockim.api.client.v1.protocol.http.LoginRequest
	1, // 2: rockim.api.client.v1.protocol.http.AuthAPI.Login:output_type -> rockim.api.client.v1.protocol.http.LoginResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rockim_api_client_v1_protocol_http_auth_proto_init() }
func file_rockim_api_client_v1_protocol_http_auth_proto_init() {
	if File_rockim_api_client_v1_protocol_http_auth_proto != nil {
		return
	}
	file_rockim_api_client_v1_protocol_http_http_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rockim_api_client_v1_protocol_http_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_rockim_api_client_v1_protocol_http_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
			RawDescriptor: file_rockim_api_client_v1_protocol_http_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rockim_api_client_v1_protocol_http_auth_proto_goTypes,
		DependencyIndexes: file_rockim_api_client_v1_protocol_http_auth_proto_depIdxs,
		MessageInfos:      file_rockim_api_client_v1_protocol_http_auth_proto_msgTypes,
	}.Build()
	File_rockim_api_client_v1_protocol_http_auth_proto = out.File
	file_rockim_api_client_v1_protocol_http_auth_proto_rawDesc = nil
	file_rockim_api_client_v1_protocol_http_auth_proto_goTypes = nil
	file_rockim_api_client_v1_protocol_http_auth_proto_depIdxs = nil
}
func (x *LoginRequest) SetBase(base *APIRequest) {
	if x == nil {
		return
	}
	x.Base = base
}

func (x *LoginRequest) SetAuthCode(authCode string) {
	if x == nil {
		return
	}
	x.AuthCode = authCode
}

func (x *LoginResponse) SetToken(token string) {
	if x == nil {
		return
	}
	x.Token = token
}

func (x *LoginResponse) SetExpireTime(expireTime int64) {
	if x == nil {
		return
	}
	x.ExpireTime = expireTime
}
