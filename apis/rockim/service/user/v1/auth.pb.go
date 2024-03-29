// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/service/user/v1/auth.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	service "rockimserver/apis/rockim/service"
	types "rockimserver/apis/rockim/service/user/v1/types"
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

	// 基础参数
	Base *service.ServiceRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 用户账户名，由接入方指定
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *AuthCodeRequest) Reset() {
	*x = AuthCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_user_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthCodeRequest) ProtoMessage() {}

func (x *AuthCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_user_v1_auth_proto_msgTypes[0]
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
	return file_rockim_service_user_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthCodeRequest) GetBase() *service.ServiceRequest {
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

	// 授权码
	AuthCode string `protobuf:"bytes,1,opt,name=auth_code,json=authCode,proto3" json:"auth_code,omitempty"`
	// 失效时间（时间戳）
	ExpireTime int64 `protobuf:"varint,2,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *AuthCodeResponse) Reset() {
	*x = AuthCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_user_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthCodeResponse) ProtoMessage() {}

func (x *AuthCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_user_v1_auth_proto_msgTypes[1]
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
	return file_rockim_service_user_v1_auth_proto_rawDescGZIP(), []int{1}
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

// LoginRequest 登录请求
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基础参数
	Base *service.ServiceRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 授权码（只能使用一次）
	AuthCode string `protobuf:"bytes,2,opt,name=auth_code,json=authCode,proto3" json:"auth_code,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_user_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_user_v1_auth_proto_msgTypes[2]
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
	return file_rockim_service_user_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetBase() *service.ServiceRequest {
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

	// 接口访问TOKEN
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// 失效时间（时间戳）
	ExpireTime int64 `protobuf:"varint,2,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	// 用户信息
	User *types.User `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_user_v1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_user_v1_auth_proto_msgTypes[3]
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
	return file_rockim_service_user_v1_auth_proto_rawDescGZIP(), []int{3}
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

func (x *LoginResponse) GetUser() *types.User {
	if x != nil {
		return x.User
	}
	return nil
}

// TokenCheckRequest 令牌验证请求
type TokenCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基础参数
	Base *service.ServiceRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 访问令牌
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TokenCheckRequest) Reset() {
	*x = TokenCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_user_v1_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenCheckRequest) ProtoMessage() {}

func (x *TokenCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_user_v1_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenCheckRequest.ProtoReflect.Descriptor instead.
func (*TokenCheckRequest) Descriptor() ([]byte, []int) {
	return file_rockim_service_user_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *TokenCheckRequest) GetBase() *service.ServiceRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *TokenCheckRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// TokenCheckResponse 令牌验证响应
type TokenCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前用户ID
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *TokenCheckResponse) Reset() {
	*x = TokenCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_service_user_v1_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenCheckResponse) ProtoMessage() {}

func (x *TokenCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_service_user_v1_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenCheckResponse.ProtoReflect.Descriptor instead.
func (*TokenCheckResponse) Descriptor() ([]byte, []int) {
	return file_rockim_service_user_v1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *TokenCheckResponse) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

var File_rockim_service_user_v1_auth_proto protoreflect.FileDescriptor

var file_rockim_service_user_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x27, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x0f, 0x41,
	0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c,
	0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x50, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x72, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3c, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x7e, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x70, 0x0a, 0x11, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x26, 0x0a, 0x12, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x32,
	0xa9, 0x02, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x41, 0x50, 0x49, 0x12, 0x63, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x2e,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x54, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x24, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x29, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_rockim_service_user_v1_auth_proto_rawDescOnce sync.Once
	file_rockim_service_user_v1_auth_proto_rawDescData = file_rockim_service_user_v1_auth_proto_rawDesc
)

func file_rockim_service_user_v1_auth_proto_rawDescGZIP() []byte {
	file_rockim_service_user_v1_auth_proto_rawDescOnce.Do(func() {
		file_rockim_service_user_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_service_user_v1_auth_proto_rawDescData)
	})
	return file_rockim_service_user_v1_auth_proto_rawDescData
}

var file_rockim_service_user_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rockim_service_user_v1_auth_proto_goTypes = []interface{}{
	(*AuthCodeRequest)(nil),        // 0: rockim.service.user.v1.AuthCodeRequest
	(*AuthCodeResponse)(nil),       // 1: rockim.service.user.v1.AuthCodeResponse
	(*LoginRequest)(nil),           // 2: rockim.service.user.v1.LoginRequest
	(*LoginResponse)(nil),          // 3: rockim.service.user.v1.LoginResponse
	(*TokenCheckRequest)(nil),      // 4: rockim.service.user.v1.TokenCheckRequest
	(*TokenCheckResponse)(nil),     // 5: rockim.service.user.v1.TokenCheckResponse
	(*service.ServiceRequest)(nil), // 6: rockim.service.ServiceRequest
	(*types.User)(nil),             // 7: rockim.service.user.v1.types.User
}
var file_rockim_service_user_v1_auth_proto_depIdxs = []int32{
	6, // 0: rockim.service.user.v1.AuthCodeRequest.base:type_name -> rockim.service.ServiceRequest
	6, // 1: rockim.service.user.v1.LoginRequest.base:type_name -> rockim.service.ServiceRequest
	7, // 2: rockim.service.user.v1.LoginResponse.user:type_name -> rockim.service.user.v1.types.User
	6, // 3: rockim.service.user.v1.TokenCheckRequest.base:type_name -> rockim.service.ServiceRequest
	0, // 4: rockim.service.user.v1.AuthAPI.CreateAuthCode:input_type -> rockim.service.user.v1.AuthCodeRequest
	2, // 5: rockim.service.user.v1.AuthAPI.Login:input_type -> rockim.service.user.v1.LoginRequest
	4, // 6: rockim.service.user.v1.AuthAPI.CheckToken:input_type -> rockim.service.user.v1.TokenCheckRequest
	1, // 7: rockim.service.user.v1.AuthAPI.CreateAuthCode:output_type -> rockim.service.user.v1.AuthCodeResponse
	3, // 8: rockim.service.user.v1.AuthAPI.Login:output_type -> rockim.service.user.v1.LoginResponse
	5, // 9: rockim.service.user.v1.AuthAPI.CheckToken:output_type -> rockim.service.user.v1.TokenCheckResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rockim_service_user_v1_auth_proto_init() }
func file_rockim_service_user_v1_auth_proto_init() {
	if File_rockim_service_user_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_service_user_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_rockim_service_user_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_rockim_service_user_v1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_rockim_service_user_v1_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_rockim_service_user_v1_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenCheckRequest); i {
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
		file_rockim_service_user_v1_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenCheckResponse); i {
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
			RawDescriptor: file_rockim_service_user_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rockim_service_user_v1_auth_proto_goTypes,
		DependencyIndexes: file_rockim_service_user_v1_auth_proto_depIdxs,
		MessageInfos:      file_rockim_service_user_v1_auth_proto_msgTypes,
	}.Build()
	File_rockim_service_user_v1_auth_proto = out.File
	file_rockim_service_user_v1_auth_proto_rawDesc = nil
	file_rockim_service_user_v1_auth_proto_goTypes = nil
	file_rockim_service_user_v1_auth_proto_depIdxs = nil
}
