// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: api/rockim/service/platform/v1/tenant.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	types1 "rockim/api/rockim/service/platform/v1/types"
	types "rockim/api/rockim/shared/types"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TenantCreateRequest 创建商户请求
type TenantCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 账号
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// 密码
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// 名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TenantCreateRequest) Reset() {
	*x = TenantCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantCreateRequest) ProtoMessage() {}

func (x *TenantCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantCreateRequest.ProtoReflect.Descriptor instead.
func (*TenantCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_rockim_service_platform_v1_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *TenantCreateRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *TenantCreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *TenantCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// TenantCreateResponse 创建商户响应
type TenantCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TenantCreateResponse) Reset() {
	*x = TenantCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantCreateResponse) ProtoMessage() {}

func (x *TenantCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantCreateResponse.ProtoReflect.Descriptor instead.
func (*TenantCreateResponse) Descriptor() ([]byte, []int) {
	return file_api_rockim_service_platform_v1_tenant_proto_rawDescGZIP(), []int{1}
}

// TenantUpdateRequest 更新商户请求
type TenantUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 账号
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TenantUpdateRequest) Reset() {
	*x = TenantUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantUpdateRequest) ProtoMessage() {}

func (x *TenantUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantUpdateRequest.ProtoReflect.Descriptor instead.
func (*TenantUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_rockim_service_platform_v1_tenant_proto_rawDescGZIP(), []int{2}
}

func (x *TenantUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TenantUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// TenantUpdateResponse 更新商户响应
type TenantUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TenantUpdateResponse) Reset() {
	*x = TenantUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantUpdateResponse) ProtoMessage() {}

func (x *TenantUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantUpdateResponse.ProtoReflect.Descriptor instead.
func (*TenantUpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_rockim_service_platform_v1_tenant_proto_rawDescGZIP(), []int{3}
}

// TenantPagingRequest 分页获取商户请求
type TenantPagingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginate *types.Paginating `protobuf:"bytes,1,opt,name=paginate,proto3" json:"paginate,omitempty"`
	Keyword  string            `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *TenantPagingRequest) Reset() {
	*x = TenantPagingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantPagingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantPagingRequest) ProtoMessage() {}

func (x *TenantPagingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantPagingRequest.ProtoReflect.Descriptor instead.
func (*TenantPagingRequest) Descriptor() ([]byte, []int) {
	return file_api_rockim_service_platform_v1_tenant_proto_rawDescGZIP(), []int{4}
}

func (x *TenantPagingRequest) GetPaginate() *types.Paginating {
	if x != nil {
		return x.Paginate
	}
	return nil
}

func (x *TenantPagingRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

// TenantPagingResponse 分页获取商户响应
type TenantPagingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 商户列表
	List []*types1.Tenant `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	// 分页数据
	Paginate *types.Paginated `protobuf:"bytes,2,opt,name=paginate,proto3" json:"paginate,omitempty"`
}

func (x *TenantPagingResponse) Reset() {
	*x = TenantPagingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantPagingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantPagingResponse) ProtoMessage() {}

func (x *TenantPagingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantPagingResponse.ProtoReflect.Descriptor instead.
func (*TenantPagingResponse) Descriptor() ([]byte, []int) {
	return file_api_rockim_service_platform_v1_tenant_proto_rawDescGZIP(), []int{5}
}

func (x *TenantPagingResponse) GetList() []*types1.Tenant {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *TenantPagingResponse) GetPaginate() *types.Paginated {
	if x != nil {
		return x.Paginate
	}
	return nil
}

// TenantFindRequest 查找商户请求
type TenantFindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *TenantFindRequest) Reset() {
	*x = TenantFindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantFindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantFindRequest) ProtoMessage() {}

func (x *TenantFindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantFindRequest.ProtoReflect.Descriptor instead.
func (*TenantFindRequest) Descriptor() ([]byte, []int) {
	return file_api_rockim_service_platform_v1_tenant_proto_rawDescGZIP(), []int{6}
}

func (x *TenantFindRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// TenantFindResponse 查找商户响应
type TenantFindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant *types1.Tenant `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *TenantFindResponse) Reset() {
	*x = TenantFindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantFindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantFindResponse) ProtoMessage() {}

func (x *TenantFindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_service_platform_v1_tenant_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantFindResponse.ProtoReflect.Descriptor instead.
func (*TenantFindResponse) Descriptor() ([]byte, []int) {
	return file_api_rockim_service_platform_v1_tenant_proto_rawDescGZIP(), []int{7}
}

func (x *TenantFindResponse) GetTenant() *types1.Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

var File_api_rockim_service_platform_v1_tenant_proto protoreflect.FileDescriptor

var file_api_rockim_service_platform_v1_tenant_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2d, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x13, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x23, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x16, 0x0a, 0x14, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x0a, 0x13, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6c,
	0x0a, 0x13, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x90, 0x01, 0x0a,
	0x14, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x64, 0x52, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x22,
	0x2d, 0x0a, 0x11, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x56,
	0x0a, 0x12, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x32, 0xb9, 0x03, 0x0a, 0x09, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x41, 0x50, 0x49, 0x12, 0x6b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2f,
	0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b,
	0x0a, 0x06, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x04, 0x46,
	0x69, 0x6e, 0x64, 0x12, 0x2d, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_rockim_service_platform_v1_tenant_proto_rawDescOnce sync.Once
	file_api_rockim_service_platform_v1_tenant_proto_rawDescData = file_api_rockim_service_platform_v1_tenant_proto_rawDesc
)

func file_api_rockim_service_platform_v1_tenant_proto_rawDescGZIP() []byte {
	file_api_rockim_service_platform_v1_tenant_proto_rawDescOnce.Do(func() {
		file_api_rockim_service_platform_v1_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rockim_service_platform_v1_tenant_proto_rawDescData)
	})
	return file_api_rockim_service_platform_v1_tenant_proto_rawDescData
}

var file_api_rockim_service_platform_v1_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_rockim_service_platform_v1_tenant_proto_goTypes = []interface{}{
	(*TenantCreateRequest)(nil),  // 0: rockim.service.platform.v1.TenantCreateRequest
	(*TenantCreateResponse)(nil), // 1: rockim.service.platform.v1.TenantCreateResponse
	(*TenantUpdateRequest)(nil),  // 2: rockim.service.platform.v1.TenantUpdateRequest
	(*TenantUpdateResponse)(nil), // 3: rockim.service.platform.v1.TenantUpdateResponse
	(*TenantPagingRequest)(nil),  // 4: rockim.service.platform.v1.TenantPagingRequest
	(*TenantPagingResponse)(nil), // 5: rockim.service.platform.v1.TenantPagingResponse
	(*TenantFindRequest)(nil),    // 6: rockim.service.platform.v1.TenantFindRequest
	(*TenantFindResponse)(nil),   // 7: rockim.service.platform.v1.TenantFindResponse
	(*types.Paginating)(nil),     // 8: rockim.shared.types.Paginating
	(*types1.Tenant)(nil),        // 9: rockim.service.platform.v1.types.Tenant
	(*types.Paginated)(nil),      // 10: rockim.shared.types.Paginated
}
var file_api_rockim_service_platform_v1_tenant_proto_depIdxs = []int32{
	8,  // 0: rockim.service.platform.v1.TenantPagingRequest.paginate:type_name -> rockim.shared.types.Paginating
	9,  // 1: rockim.service.platform.v1.TenantPagingResponse.list:type_name -> rockim.service.platform.v1.types.Tenant
	10, // 2: rockim.service.platform.v1.TenantPagingResponse.paginate:type_name -> rockim.shared.types.Paginated
	9,  // 3: rockim.service.platform.v1.TenantFindResponse.tenant:type_name -> rockim.service.platform.v1.types.Tenant
	0,  // 4: rockim.service.platform.v1.TenantAPI.Create:input_type -> rockim.service.platform.v1.TenantCreateRequest
	2,  // 5: rockim.service.platform.v1.TenantAPI.Update:input_type -> rockim.service.platform.v1.TenantUpdateRequest
	4,  // 6: rockim.service.platform.v1.TenantAPI.Paging:input_type -> rockim.service.platform.v1.TenantPagingRequest
	6,  // 7: rockim.service.platform.v1.TenantAPI.Find:input_type -> rockim.service.platform.v1.TenantFindRequest
	1,  // 8: rockim.service.platform.v1.TenantAPI.Create:output_type -> rockim.service.platform.v1.TenantCreateResponse
	3,  // 9: rockim.service.platform.v1.TenantAPI.Update:output_type -> rockim.service.platform.v1.TenantUpdateResponse
	5,  // 10: rockim.service.platform.v1.TenantAPI.Paging:output_type -> rockim.service.platform.v1.TenantPagingResponse
	7,  // 11: rockim.service.platform.v1.TenantAPI.Find:output_type -> rockim.service.platform.v1.TenantFindResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_rockim_service_platform_v1_tenant_proto_init() }
func file_api_rockim_service_platform_v1_tenant_proto_init() {
	if File_api_rockim_service_platform_v1_tenant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_rockim_service_platform_v1_tenant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantCreateRequest); i {
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
		file_api_rockim_service_platform_v1_tenant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantCreateResponse); i {
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
		file_api_rockim_service_platform_v1_tenant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantUpdateRequest); i {
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
		file_api_rockim_service_platform_v1_tenant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantUpdateResponse); i {
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
		file_api_rockim_service_platform_v1_tenant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantPagingRequest); i {
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
		file_api_rockim_service_platform_v1_tenant_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantPagingResponse); i {
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
		file_api_rockim_service_platform_v1_tenant_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantFindRequest); i {
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
		file_api_rockim_service_platform_v1_tenant_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantFindResponse); i {
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
			RawDescriptor: file_api_rockim_service_platform_v1_tenant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_rockim_service_platform_v1_tenant_proto_goTypes,
		DependencyIndexes: file_api_rockim_service_platform_v1_tenant_proto_depIdxs,
		MessageInfos:      file_api_rockim_service_platform_v1_tenant_proto_msgTypes,
	}.Build()
	File_api_rockim_service_platform_v1_tenant_proto = out.File
	file_api_rockim_service_platform_v1_tenant_proto_rawDesc = nil
	file_api_rockim_service_platform_v1_tenant_proto_goTypes = nil
	file_api_rockim_service_platform_v1_tenant_proto_depIdxs = nil
}
