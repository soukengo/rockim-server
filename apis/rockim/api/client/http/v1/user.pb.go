// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/api/client/http/v1/user.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	types "rockimserver/apis/rockim/api/client/http/v1/types"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// UserFindRequest 用户查找请求
type UserFindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基本参数
	Base *APIRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 用户账户名，由接入方指定
	Account string `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *UserFindRequest) Reset() {
	*x = UserFindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_http_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFindRequest) ProtoMessage() {}

func (x *UserFindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_http_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFindRequest.ProtoReflect.Descriptor instead.
func (*UserFindRequest) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_http_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserFindRequest) GetBase() *APIRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UserFindRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// UserFindResponse 用户查找响应
type UserFindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *types.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserFindResponse) Reset() {
	*x = UserFindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_http_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFindResponse) ProtoMessage() {}

func (x *UserFindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_http_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFindResponse.ProtoReflect.Descriptor instead.
func (*UserFindResponse) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_http_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserFindResponse) GetUser() *types.User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_rockim_api_client_http_v1_user_proto protoreflect.FileDescriptor

var file_rockim_api_client_http_v1_user_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x0f, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x32, 0x8c, 0x01, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x41, 0x50, 0x49,
	0x12, 0x80, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x2a, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x66, 0x69, 0x6e, 0x64,
	0x3a, 0x01, 0x2a, 0x42, 0x30, 0x5a, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_api_client_http_v1_user_proto_rawDescOnce sync.Once
	file_rockim_api_client_http_v1_user_proto_rawDescData = file_rockim_api_client_http_v1_user_proto_rawDesc
)

func file_rockim_api_client_http_v1_user_proto_rawDescGZIP() []byte {
	file_rockim_api_client_http_v1_user_proto_rawDescOnce.Do(func() {
		file_rockim_api_client_http_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_api_client_http_v1_user_proto_rawDescData)
	})
	return file_rockim_api_client_http_v1_user_proto_rawDescData
}

var file_rockim_api_client_http_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rockim_api_client_http_v1_user_proto_goTypes = []interface{}{
	(*UserFindRequest)(nil),  // 0: rockim.api.client.http.v1.UserFindRequest
	(*UserFindResponse)(nil), // 1: rockim.api.client.http.v1.UserFindResponse
	(*APIRequest)(nil),       // 2: rockim.api.client.http.v1.APIRequest
	(*types.User)(nil),       // 3: rockim.api.client.http.v1.types.User
}
var file_rockim_api_client_http_v1_user_proto_depIdxs = []int32{
	2, // 0: rockim.api.client.http.v1.UserFindRequest.base:type_name -> rockim.api.client.http.v1.APIRequest
	3, // 1: rockim.api.client.http.v1.UserFindResponse.user:type_name -> rockim.api.client.http.v1.types.User
	0, // 2: rockim.api.client.http.v1.UserAPI.Find:input_type -> rockim.api.client.http.v1.UserFindRequest
	1, // 3: rockim.api.client.http.v1.UserAPI.Find:output_type -> rockim.api.client.http.v1.UserFindResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rockim_api_client_http_v1_user_proto_init() }
func file_rockim_api_client_http_v1_user_proto_init() {
	if File_rockim_api_client_http_v1_user_proto != nil {
		return
	}
	file_rockim_api_client_http_v1_http_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rockim_api_client_http_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFindRequest); i {
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
		file_rockim_api_client_http_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFindResponse); i {
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
			RawDescriptor: file_rockim_api_client_http_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rockim_api_client_http_v1_user_proto_goTypes,
		DependencyIndexes: file_rockim_api_client_http_v1_user_proto_depIdxs,
		MessageInfos:      file_rockim_api_client_http_v1_user_proto_msgTypes,
	}.Build()
	File_rockim_api_client_http_v1_user_proto = out.File
	file_rockim_api_client_http_v1_user_proto_rawDesc = nil
	file_rockim_api_client_http_v1_user_proto_goTypes = nil
	file_rockim_api_client_http_v1_user_proto_depIdxs = nil
}
func (x *UserFindRequest) SetBase(base *APIRequest) {
	if x == nil {
		return
	}
	x.Base = base
}

func (x *UserFindRequest) SetAccount(account string) {
	if x == nil {
		return
	}
	x.Account = account
}

func (x *UserFindResponse) SetUser(user *types.User) {
	if x == nil {
		return
	}
	x.User = user
}
