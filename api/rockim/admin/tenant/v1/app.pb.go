// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: api/rockim/admin/tenant/v1/app.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	types "rockim/api/rockim/admin/tenant/v1/types"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AppCreateRequest 创建请求
type AppCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 所属租户
	TenantId string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// 应用名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AppCreateRequest) Reset() {
	*x = AppCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppCreateRequest) ProtoMessage() {}

func (x *AppCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppCreateRequest.ProtoReflect.Descriptor instead.
func (*AppCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_tenant_v1_app_proto_rawDescGZIP(), []int{0}
}

func (x *AppCreateRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *AppCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// AppCreateRequest 创建响应
type AppCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AppCreateResponse) Reset() {
	*x = AppCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppCreateResponse) ProtoMessage() {}

func (x *AppCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppCreateResponse.ProtoReflect.Descriptor instead.
func (*AppCreateResponse) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_tenant_v1_app_proto_rawDescGZIP(), []int{1}
}

func (x *AppCreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// AppUpdateRequest 更新请求
type AppUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// app id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 应用名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 应用密钥
	AppKey string `protobuf:"bytes,3,opt,name=app_key,json=appKey,proto3" json:"app_key,omitempty"`
}

func (x *AppUpdateRequest) Reset() {
	*x = AppUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppUpdateRequest) ProtoMessage() {}

func (x *AppUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppUpdateRequest.ProtoReflect.Descriptor instead.
func (*AppUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_tenant_v1_app_proto_rawDescGZIP(), []int{2}
}

func (x *AppUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppUpdateRequest) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

// AppUpdateResponse 更新响应
type AppUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppUpdateResponse) Reset() {
	*x = AppUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppUpdateResponse) ProtoMessage() {}

func (x *AppUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppUpdateResponse.ProtoReflect.Descriptor instead.
func (*AppUpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_tenant_v1_app_proto_rawDescGZIP(), []int{3}
}

// AppFindRequest 查找应用请求
type AppFindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// app id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AppFindRequest) Reset() {
	*x = AppFindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppFindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppFindRequest) ProtoMessage() {}

func (x *AppFindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppFindRequest.ProtoReflect.Descriptor instead.
func (*AppFindRequest) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_tenant_v1_app_proto_rawDescGZIP(), []int{4}
}

func (x *AppFindRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// AppFindResponse 查找应用响应
type AppFindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App *types.App `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *AppFindResponse) Reset() {
	*x = AppFindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppFindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppFindResponse) ProtoMessage() {}

func (x *AppFindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppFindResponse.ProtoReflect.Descriptor instead.
func (*AppFindResponse) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_tenant_v1_app_proto_rawDescGZIP(), []int{5}
}

func (x *AppFindResponse) GetApp() *types.App {
	if x != nil {
		return x.App
	}
	return nil
}

// AppListRequest 获取商户应用列表请求
type AppListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppListRequest) Reset() {
	*x = AppListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppListRequest) ProtoMessage() {}

func (x *AppListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppListRequest.ProtoReflect.Descriptor instead.
func (*AppListRequest) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_tenant_v1_app_proto_rawDescGZIP(), []int{6}
}

// AppListResponse 获取商户应用列表响应
type AppListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*types.App `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *AppListResponse) Reset() {
	*x = AppListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppListResponse) ProtoMessage() {}

func (x *AppListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_tenant_v1_app_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppListResponse.ProtoReflect.Descriptor instead.
func (*AppListResponse) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_tenant_v1_app_proto_rawDescGZIP(), []int{7}
}

func (x *AppListResponse) GetList() []*types.App {
	if x != nil {
		return x.List
	}
	return nil
}

var File_api_rockim_admin_tenant_v1_app_proto protoreflect.FileDescriptor

var file_api_rockim_admin_tenant_v1_app_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a,
	0x10, 0x41, 0x70, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x10, 0x41, 0x70, 0x70,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70,
	0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70,
	0x4b, 0x65, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x46,
	0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x0f, 0x41, 0x70,
	0x70, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61,
	0x70, 0x70, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x90,
	0x04, 0x0a, 0x06, 0x41, 0x70, 0x70, 0x41, 0x50, 0x49, 0x12, 0x83, 0x01, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70,
	0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x83, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70,
	0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x7b, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x26, 0x2e,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x70, 0x70, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x3a,
	0x01, 0x2a, 0x12, 0x7d, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x2e, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01,
	0x2a, 0x42, 0x26, 0x5a, 0x24, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_rockim_admin_tenant_v1_app_proto_rawDescOnce sync.Once
	file_api_rockim_admin_tenant_v1_app_proto_rawDescData = file_api_rockim_admin_tenant_v1_app_proto_rawDesc
)

func file_api_rockim_admin_tenant_v1_app_proto_rawDescGZIP() []byte {
	file_api_rockim_admin_tenant_v1_app_proto_rawDescOnce.Do(func() {
		file_api_rockim_admin_tenant_v1_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rockim_admin_tenant_v1_app_proto_rawDescData)
	})
	return file_api_rockim_admin_tenant_v1_app_proto_rawDescData
}

var file_api_rockim_admin_tenant_v1_app_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_rockim_admin_tenant_v1_app_proto_goTypes = []interface{}{
	(*AppCreateRequest)(nil),  // 0: rockim.admin.tenant.v1.AppCreateRequest
	(*AppCreateResponse)(nil), // 1: rockim.admin.tenant.v1.AppCreateResponse
	(*AppUpdateRequest)(nil),  // 2: rockim.admin.tenant.v1.AppUpdateRequest
	(*AppUpdateResponse)(nil), // 3: rockim.admin.tenant.v1.AppUpdateResponse
	(*AppFindRequest)(nil),    // 4: rockim.admin.tenant.v1.AppFindRequest
	(*AppFindResponse)(nil),   // 5: rockim.admin.tenant.v1.AppFindResponse
	(*AppListRequest)(nil),    // 6: rockim.admin.tenant.v1.AppListRequest
	(*AppListResponse)(nil),   // 7: rockim.admin.tenant.v1.AppListResponse
	(*types.App)(nil),         // 8: rockim.admin.tenant.v1.types.App
}
var file_api_rockim_admin_tenant_v1_app_proto_depIdxs = []int32{
	8, // 0: rockim.admin.tenant.v1.AppFindResponse.app:type_name -> rockim.admin.tenant.v1.types.App
	8, // 1: rockim.admin.tenant.v1.AppListResponse.list:type_name -> rockim.admin.tenant.v1.types.App
	0, // 2: rockim.admin.tenant.v1.AppAPI.Create:input_type -> rockim.admin.tenant.v1.AppCreateRequest
	2, // 3: rockim.admin.tenant.v1.AppAPI.Update:input_type -> rockim.admin.tenant.v1.AppUpdateRequest
	4, // 4: rockim.admin.tenant.v1.AppAPI.Find:input_type -> rockim.admin.tenant.v1.AppFindRequest
	6, // 5: rockim.admin.tenant.v1.AppAPI.Paging:input_type -> rockim.admin.tenant.v1.AppListRequest
	1, // 6: rockim.admin.tenant.v1.AppAPI.Create:output_type -> rockim.admin.tenant.v1.AppCreateResponse
	3, // 7: rockim.admin.tenant.v1.AppAPI.Update:output_type -> rockim.admin.tenant.v1.AppUpdateResponse
	5, // 8: rockim.admin.tenant.v1.AppAPI.Find:output_type -> rockim.admin.tenant.v1.AppFindResponse
	7, // 9: rockim.admin.tenant.v1.AppAPI.Paging:output_type -> rockim.admin.tenant.v1.AppListResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_rockim_admin_tenant_v1_app_proto_init() }
func file_api_rockim_admin_tenant_v1_app_proto_init() {
	if File_api_rockim_admin_tenant_v1_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_rockim_admin_tenant_v1_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppCreateRequest); i {
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
		file_api_rockim_admin_tenant_v1_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppCreateResponse); i {
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
		file_api_rockim_admin_tenant_v1_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppUpdateRequest); i {
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
		file_api_rockim_admin_tenant_v1_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppUpdateResponse); i {
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
		file_api_rockim_admin_tenant_v1_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppFindRequest); i {
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
		file_api_rockim_admin_tenant_v1_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppFindResponse); i {
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
		file_api_rockim_admin_tenant_v1_app_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppListRequest); i {
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
		file_api_rockim_admin_tenant_v1_app_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppListResponse); i {
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
			RawDescriptor: file_api_rockim_admin_tenant_v1_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_rockim_admin_tenant_v1_app_proto_goTypes,
		DependencyIndexes: file_api_rockim_admin_tenant_v1_app_proto_depIdxs,
		MessageInfos:      file_api_rockim_admin_tenant_v1_app_proto_msgTypes,
	}.Build()
	File_api_rockim_admin_tenant_v1_app_proto = out.File
	file_api_rockim_admin_tenant_v1_app_proto_rawDesc = nil
	file_api_rockim_admin_tenant_v1_app_proto_goTypes = nil
	file_api_rockim_admin_tenant_v1_app_proto_depIdxs = nil
}
