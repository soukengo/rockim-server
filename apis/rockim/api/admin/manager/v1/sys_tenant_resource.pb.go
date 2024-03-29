// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/api/admin/manager/v1/sys_tenant_resource.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	types "rockimserver/apis/rockim/api/admin/manager/v1/types"
	enums "rockimserver/apis/rockim/shared/enums"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SysTenantResourceOptions
type SysTenantResourceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category enums.Admin_ResourceCategory `protobuf:"varint,1,opt,name=category,proto3,enum=rockim.shared.enums.Admin_ResourceCategory" json:"category,omitempty"`
	// 菜单名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 上级ID
	ParentId string `protobuf:"bytes,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// URL
	Url string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	// ICON
	Icon string `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	// 排序号
	Order int32 `protobuf:"varint,6,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *SysTenantResourceOptions) Reset() {
	*x = SysTenantResourceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysTenantResourceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysTenantResourceOptions) ProtoMessage() {}

func (x *SysTenantResourceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysTenantResourceOptions.ProtoReflect.Descriptor instead.
func (*SysTenantResourceOptions) Descriptor() ([]byte, []int) {
	return file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescGZIP(), []int{0}
}

func (x *SysTenantResourceOptions) GetCategory() enums.Admin_ResourceCategory {
	if x != nil {
		return x.Category
	}
	return enums.Admin_RESOURCE_CATEGORY_INVALID
}

func (x *SysTenantResourceOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SysTenantResourceOptions) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *SysTenantResourceOptions) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SysTenantResourceOptions) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *SysTenantResourceOptions) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type TenantResourceCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options *SysTenantResourceOptions `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *TenantResourceCreateRequest) Reset() {
	*x = TenantResourceCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantResourceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantResourceCreateRequest) ProtoMessage() {}

func (x *TenantResourceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantResourceCreateRequest.ProtoReflect.Descriptor instead.
func (*TenantResourceCreateRequest) Descriptor() ([]byte, []int) {
	return file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescGZIP(), []int{1}
}

func (x *TenantResourceCreateRequest) GetOptions() *SysTenantResourceOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type TenantResourceCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TenantResourceCreateResponse) Reset() {
	*x = TenantResourceCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantResourceCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantResourceCreateResponse) ProtoMessage() {}

func (x *TenantResourceCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantResourceCreateResponse.ProtoReflect.Descriptor instead.
func (*TenantResourceCreateResponse) Descriptor() ([]byte, []int) {
	return file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescGZIP(), []int{2}
}

type TenantResourceUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Options *SysTenantResourceOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *TenantResourceUpdateRequest) Reset() {
	*x = TenantResourceUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantResourceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantResourceUpdateRequest) ProtoMessage() {}

func (x *TenantResourceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantResourceUpdateRequest.ProtoReflect.Descriptor instead.
func (*TenantResourceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescGZIP(), []int{3}
}

func (x *TenantResourceUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TenantResourceUpdateRequest) GetOptions() *SysTenantResourceOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type TenantResourceUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TenantResourceUpdateResponse) Reset() {
	*x = TenantResourceUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantResourceUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantResourceUpdateResponse) ProtoMessage() {}

func (x *TenantResourceUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantResourceUpdateResponse.ProtoReflect.Descriptor instead.
func (*TenantResourceUpdateResponse) Descriptor() ([]byte, []int) {
	return file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescGZIP(), []int{4}
}

type TenantResourceDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TenantResourceDeleteRequest) Reset() {
	*x = TenantResourceDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantResourceDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantResourceDeleteRequest) ProtoMessage() {}

func (x *TenantResourceDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantResourceDeleteRequest.ProtoReflect.Descriptor instead.
func (*TenantResourceDeleteRequest) Descriptor() ([]byte, []int) {
	return file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescGZIP(), []int{5}
}

func (x *TenantResourceDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TenantResourceDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TenantResourceDeleteResponse) Reset() {
	*x = TenantResourceDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantResourceDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantResourceDeleteResponse) ProtoMessage() {}

func (x *TenantResourceDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantResourceDeleteResponse.ProtoReflect.Descriptor instead.
func (*TenantResourceDeleteResponse) Descriptor() ([]byte, []int) {
	return file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescGZIP(), []int{6}
}

type TenantResourceTreeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TenantResourceTreeRequest) Reset() {
	*x = TenantResourceTreeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantResourceTreeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantResourceTreeRequest) ProtoMessage() {}

func (x *TenantResourceTreeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantResourceTreeRequest.ProtoReflect.Descriptor instead.
func (*TenantResourceTreeRequest) Descriptor() ([]byte, []int) {
	return file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescGZIP(), []int{7}
}

type TenantResourceTreeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*types.SysTenantResourceTree `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *TenantResourceTreeResponse) Reset() {
	*x = TenantResourceTreeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantResourceTreeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantResourceTreeResponse) ProtoMessage() {}

func (x *TenantResourceTreeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantResourceTreeResponse.ProtoReflect.Descriptor instead.
func (*TenantResourceTreeResponse) Descriptor() ([]byte, []int) {
	return file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescGZIP(), []int{8}
}

func (x *TenantResourceTreeResponse) GetList() []*types.SysTenantResourceTree {
	if x != nil {
		return x.List
	}
	return nil
}

var File_rockim_api_admin_manager_v1_sys_tenant_resource_proto protoreflect.FileDescriptor

var file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDesc = []byte{
	0x0a, 0x35, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79,
	0x73, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2b, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x18, 0x53, 0x79,
	0x73, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x51, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x78, 0x0a, 0x1b, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x59, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x1e, 0x0a, 0x1c, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x1b, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x59, 0x0a, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1e, 0x0a, 0x1c, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x1b, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e,
	0x0a, 0x1c, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x0a, 0x19, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6a, 0x0a, 0x1a, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x65,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x73, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x65,
	0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xf6, 0x05, 0x0a, 0x14, 0x53, 0x79, 0x73, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x50, 0x49,
	0x12, 0xb6, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x38, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x22, 0x2c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x2f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0xb6, 0x01, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x38, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39,
	0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x31, 0x22, 0x2c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0xb6, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x38, 0x2e,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x22, 0x2c, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79,
	0x73, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0xb2, 0x01, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x65, 0x65, 0x12, 0x36, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x37, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x65,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2f, 0x22, 0x2a, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x3a, 0x01, 0x2a,
	0x42, 0x32, 0x5a, 0x30, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescOnce sync.Once
	file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescData = file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDesc
)

func file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescGZIP() []byte {
	file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescOnce.Do(func() {
		file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescData)
	})
	return file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDescData
}

var file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_goTypes = []interface{}{
	(*SysTenantResourceOptions)(nil),     // 0: rockim.api.admin.manager.v1.SysTenantResourceOptions
	(*TenantResourceCreateRequest)(nil),  // 1: rockim.api.admin.manager.v1.TenantResourceCreateRequest
	(*TenantResourceCreateResponse)(nil), // 2: rockim.api.admin.manager.v1.TenantResourceCreateResponse
	(*TenantResourceUpdateRequest)(nil),  // 3: rockim.api.admin.manager.v1.TenantResourceUpdateRequest
	(*TenantResourceUpdateResponse)(nil), // 4: rockim.api.admin.manager.v1.TenantResourceUpdateResponse
	(*TenantResourceDeleteRequest)(nil),  // 5: rockim.api.admin.manager.v1.TenantResourceDeleteRequest
	(*TenantResourceDeleteResponse)(nil), // 6: rockim.api.admin.manager.v1.TenantResourceDeleteResponse
	(*TenantResourceTreeRequest)(nil),    // 7: rockim.api.admin.manager.v1.TenantResourceTreeRequest
	(*TenantResourceTreeResponse)(nil),   // 8: rockim.api.admin.manager.v1.TenantResourceTreeResponse
	(enums.Admin_ResourceCategory)(0),    // 9: rockim.shared.enums.Admin.ResourceCategory
	(*types.SysTenantResourceTree)(nil),  // 10: rockim.api.admin.manager.v1.types.SysTenantResourceTree
}
var file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_depIdxs = []int32{
	9,  // 0: rockim.api.admin.manager.v1.SysTenantResourceOptions.category:type_name -> rockim.shared.enums.Admin.ResourceCategory
	0,  // 1: rockim.api.admin.manager.v1.TenantResourceCreateRequest.options:type_name -> rockim.api.admin.manager.v1.SysTenantResourceOptions
	0,  // 2: rockim.api.admin.manager.v1.TenantResourceUpdateRequest.options:type_name -> rockim.api.admin.manager.v1.SysTenantResourceOptions
	10, // 3: rockim.api.admin.manager.v1.TenantResourceTreeResponse.list:type_name -> rockim.api.admin.manager.v1.types.SysTenantResourceTree
	1,  // 4: rockim.api.admin.manager.v1.SysTenantResourceAPI.Create:input_type -> rockim.api.admin.manager.v1.TenantResourceCreateRequest
	3,  // 5: rockim.api.admin.manager.v1.SysTenantResourceAPI.Update:input_type -> rockim.api.admin.manager.v1.TenantResourceUpdateRequest
	5,  // 6: rockim.api.admin.manager.v1.SysTenantResourceAPI.Delete:input_type -> rockim.api.admin.manager.v1.TenantResourceDeleteRequest
	7,  // 7: rockim.api.admin.manager.v1.SysTenantResourceAPI.ListTree:input_type -> rockim.api.admin.manager.v1.TenantResourceTreeRequest
	2,  // 8: rockim.api.admin.manager.v1.SysTenantResourceAPI.Create:output_type -> rockim.api.admin.manager.v1.TenantResourceCreateResponse
	4,  // 9: rockim.api.admin.manager.v1.SysTenantResourceAPI.Update:output_type -> rockim.api.admin.manager.v1.TenantResourceUpdateResponse
	6,  // 10: rockim.api.admin.manager.v1.SysTenantResourceAPI.Delete:output_type -> rockim.api.admin.manager.v1.TenantResourceDeleteResponse
	8,  // 11: rockim.api.admin.manager.v1.SysTenantResourceAPI.ListTree:output_type -> rockim.api.admin.manager.v1.TenantResourceTreeResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_init() }
func file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_init() {
	if File_rockim_api_admin_manager_v1_sys_tenant_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysTenantResourceOptions); i {
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
		file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantResourceCreateRequest); i {
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
		file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantResourceCreateResponse); i {
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
		file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantResourceUpdateRequest); i {
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
		file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantResourceUpdateResponse); i {
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
		file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantResourceDeleteRequest); i {
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
		file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantResourceDeleteResponse); i {
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
		file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantResourceTreeRequest); i {
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
		file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantResourceTreeResponse); i {
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
			RawDescriptor: file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_goTypes,
		DependencyIndexes: file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_depIdxs,
		MessageInfos:      file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_msgTypes,
	}.Build()
	File_rockim_api_admin_manager_v1_sys_tenant_resource_proto = out.File
	file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_rawDesc = nil
	file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_goTypes = nil
	file_rockim_api_admin_manager_v1_sys_tenant_resource_proto_depIdxs = nil
}
