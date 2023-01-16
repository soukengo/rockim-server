// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: api/rockim/admin/manager/v1/sys_resource.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	types "rockim/api/rockim/admin/manager/v1/types"
	enums "rockim/api/rockim/shared/enums"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SysResourceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category enums.AdminResourceCategory `protobuf:"varint,1,opt,name=category,proto3,enum=rockim.shared.enums.AdminResourceCategory" json:"category,omitempty"`
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

func (x *SysResourceOptions) Reset() {
	*x = SysResourceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysResourceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysResourceOptions) ProtoMessage() {}

func (x *SysResourceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysResourceOptions.ProtoReflect.Descriptor instead.
func (*SysResourceOptions) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescGZIP(), []int{0}
}

func (x *SysResourceOptions) GetCategory() enums.AdminResourceCategory {
	if x != nil {
		return x.Category
	}
	return enums.AdminResourceCategory_ADMIN_RESOURCE_CATEGORY_INVALID
}

func (x *SysResourceOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SysResourceOptions) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *SysResourceOptions) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SysResourceOptions) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *SysResourceOptions) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type SysResourceCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options *SysResourceOptions `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *SysResourceCreateRequest) Reset() {
	*x = SysResourceCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysResourceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysResourceCreateRequest) ProtoMessage() {}

func (x *SysResourceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysResourceCreateRequest.ProtoReflect.Descriptor instead.
func (*SysResourceCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescGZIP(), []int{1}
}

func (x *SysResourceCreateRequest) GetOptions() *SysResourceOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type SysResourceCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysResourceCreateResponse) Reset() {
	*x = SysResourceCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysResourceCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysResourceCreateResponse) ProtoMessage() {}

func (x *SysResourceCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysResourceCreateResponse.ProtoReflect.Descriptor instead.
func (*SysResourceCreateResponse) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescGZIP(), []int{2}
}

type SysResourceUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Options *SysResourceOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *SysResourceUpdateRequest) Reset() {
	*x = SysResourceUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysResourceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysResourceUpdateRequest) ProtoMessage() {}

func (x *SysResourceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysResourceUpdateRequest.ProtoReflect.Descriptor instead.
func (*SysResourceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescGZIP(), []int{3}
}

func (x *SysResourceUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SysResourceUpdateRequest) GetOptions() *SysResourceOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type SysResourceUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysResourceUpdateResponse) Reset() {
	*x = SysResourceUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysResourceUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysResourceUpdateResponse) ProtoMessage() {}

func (x *SysResourceUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysResourceUpdateResponse.ProtoReflect.Descriptor instead.
func (*SysResourceUpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescGZIP(), []int{4}
}

type SysResourceDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SysResourceDeleteRequest) Reset() {
	*x = SysResourceDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysResourceDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysResourceDeleteRequest) ProtoMessage() {}

func (x *SysResourceDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysResourceDeleteRequest.ProtoReflect.Descriptor instead.
func (*SysResourceDeleteRequest) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescGZIP(), []int{5}
}

func (x *SysResourceDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SysResourceDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysResourceDeleteResponse) Reset() {
	*x = SysResourceDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysResourceDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysResourceDeleteResponse) ProtoMessage() {}

func (x *SysResourceDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysResourceDeleteResponse.ProtoReflect.Descriptor instead.
func (*SysResourceDeleteResponse) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescGZIP(), []int{6}
}

type SysResourceTreeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysResourceTreeRequest) Reset() {
	*x = SysResourceTreeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysResourceTreeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysResourceTreeRequest) ProtoMessage() {}

func (x *SysResourceTreeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysResourceTreeRequest.ProtoReflect.Descriptor instead.
func (*SysResourceTreeRequest) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescGZIP(), []int{7}
}

type SysResourceTreeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*types.SysResourceTree `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *SysResourceTreeResponse) Reset() {
	*x = SysResourceTreeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysResourceTreeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysResourceTreeResponse) ProtoMessage() {}

func (x *SysResourceTreeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysResourceTreeResponse.ProtoReflect.Descriptor instead.
func (*SysResourceTreeResponse) Descriptor() ([]byte, []int) {
	return file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescGZIP(), []int{8}
}

func (x *SysResourceTreeResponse) GetList() []*types.SysResourceTree {
	if x != nil {
		return x.List
	}
	return nil
}

var File_api_rockim_admin_manager_v1_sys_resource_proto protoreflect.FileDescriptor

var file_api_rockim_admin_manager_v1_sys_resource_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79,
	0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x73, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x12, 0x53,
	0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x50, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x22, 0x6b, 0x0a, 0x18, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4f,
	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x1b, 0x0a, 0x19, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x84, 0x01, 0x0a,
	0x18, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x4f, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79,
	0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x33, 0x0a, 0x18, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x17,
	0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x94, 0x05, 0x0a, 0x0e,
	0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x50, 0x49, 0x12, 0x9f,
	0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x9f, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32,
	0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0x9f, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x31, 0x2e,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x32, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x9b, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x65,
	0x65, 0x12, 0x2f, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x30, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x22, 0x21, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x3a,
	0x01, 0x2a, 0x42, 0x27, 0x5a, 0x25, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescOnce sync.Once
	file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescData = file_api_rockim_admin_manager_v1_sys_resource_proto_rawDesc
)

func file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescGZIP() []byte {
	file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescOnce.Do(func() {
		file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescData)
	})
	return file_api_rockim_admin_manager_v1_sys_resource_proto_rawDescData
}

var file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_rockim_admin_manager_v1_sys_resource_proto_goTypes = []interface{}{
	(*SysResourceOptions)(nil),        // 0: rockim.admin.manager.v1.SysResourceOptions
	(*SysResourceCreateRequest)(nil),  // 1: rockim.admin.manager.v1.SysResourceCreateRequest
	(*SysResourceCreateResponse)(nil), // 2: rockim.admin.manager.v1.SysResourceCreateResponse
	(*SysResourceUpdateRequest)(nil),  // 3: rockim.admin.manager.v1.SysResourceUpdateRequest
	(*SysResourceUpdateResponse)(nil), // 4: rockim.admin.manager.v1.SysResourceUpdateResponse
	(*SysResourceDeleteRequest)(nil),  // 5: rockim.admin.manager.v1.SysResourceDeleteRequest
	(*SysResourceDeleteResponse)(nil), // 6: rockim.admin.manager.v1.SysResourceDeleteResponse
	(*SysResourceTreeRequest)(nil),    // 7: rockim.admin.manager.v1.SysResourceTreeRequest
	(*SysResourceTreeResponse)(nil),   // 8: rockim.admin.manager.v1.SysResourceTreeResponse
	(enums.AdminResourceCategory)(0),  // 9: rockim.shared.enums.AdminResourceCategory
	(*types.SysResourceTree)(nil),     // 10: rockim.admin.manager.v1.types.SysResourceTree
}
var file_api_rockim_admin_manager_v1_sys_resource_proto_depIdxs = []int32{
	9,  // 0: rockim.admin.manager.v1.SysResourceOptions.category:type_name -> rockim.shared.enums.AdminResourceCategory
	0,  // 1: rockim.admin.manager.v1.SysResourceCreateRequest.options:type_name -> rockim.admin.manager.v1.SysResourceOptions
	0,  // 2: rockim.admin.manager.v1.SysResourceUpdateRequest.options:type_name -> rockim.admin.manager.v1.SysResourceOptions
	10, // 3: rockim.admin.manager.v1.SysResourceTreeResponse.list:type_name -> rockim.admin.manager.v1.types.SysResourceTree
	1,  // 4: rockim.admin.manager.v1.SysResourceAPI.Create:input_type -> rockim.admin.manager.v1.SysResourceCreateRequest
	3,  // 5: rockim.admin.manager.v1.SysResourceAPI.Update:input_type -> rockim.admin.manager.v1.SysResourceUpdateRequest
	5,  // 6: rockim.admin.manager.v1.SysResourceAPI.Delete:input_type -> rockim.admin.manager.v1.SysResourceDeleteRequest
	7,  // 7: rockim.admin.manager.v1.SysResourceAPI.ListTree:input_type -> rockim.admin.manager.v1.SysResourceTreeRequest
	2,  // 8: rockim.admin.manager.v1.SysResourceAPI.Create:output_type -> rockim.admin.manager.v1.SysResourceCreateResponse
	4,  // 9: rockim.admin.manager.v1.SysResourceAPI.Update:output_type -> rockim.admin.manager.v1.SysResourceUpdateResponse
	6,  // 10: rockim.admin.manager.v1.SysResourceAPI.Delete:output_type -> rockim.admin.manager.v1.SysResourceDeleteResponse
	8,  // 11: rockim.admin.manager.v1.SysResourceAPI.ListTree:output_type -> rockim.admin.manager.v1.SysResourceTreeResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_rockim_admin_manager_v1_sys_resource_proto_init() }
func file_api_rockim_admin_manager_v1_sys_resource_proto_init() {
	if File_api_rockim_admin_manager_v1_sys_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysResourceOptions); i {
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
		file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysResourceCreateRequest); i {
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
		file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysResourceCreateResponse); i {
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
		file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysResourceUpdateRequest); i {
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
		file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysResourceUpdateResponse); i {
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
		file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysResourceDeleteRequest); i {
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
		file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysResourceDeleteResponse); i {
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
		file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysResourceTreeRequest); i {
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
		file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysResourceTreeResponse); i {
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
			RawDescriptor: file_api_rockim_admin_manager_v1_sys_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_rockim_admin_manager_v1_sys_resource_proto_goTypes,
		DependencyIndexes: file_api_rockim_admin_manager_v1_sys_resource_proto_depIdxs,
		MessageInfos:      file_api_rockim_admin_manager_v1_sys_resource_proto_msgTypes,
	}.Build()
	File_api_rockim_admin_manager_v1_sys_resource_proto = out.File
	file_api_rockim_admin_manager_v1_sys_resource_proto_rawDesc = nil
	file_api_rockim_admin_manager_v1_sys_resource_proto_goTypes = nil
	file_api_rockim_admin_manager_v1_sys_resource_proto_depIdxs = nil
}
