// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: api/rockim/service/platform/v1/types/platform.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	v1 "rockim/api/rockim/shared/enums/v1"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PlatUser 平台用户
type PlatUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 平台用户ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 创建时间
	CreateTime int64 `protobuf:"varint,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// 创建时间
	UpdateTime int64 `protobuf:"varint,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// 登录名
	Account string `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	// 密码
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	// 用户名称
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// 是否超级管理员
	IsAdmin bool `protobuf:"varint,7,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
}

func (x *PlatUser) Reset() {
	*x = PlatUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_service_platform_v1_types_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatUser) ProtoMessage() {}

func (x *PlatUser) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_service_platform_v1_types_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatUser.ProtoReflect.Descriptor instead.
func (*PlatUser) Descriptor() ([]byte, []int) {
	return file_api_rockim_service_platform_v1_types_platform_proto_rawDescGZIP(), []int{0}
}

func (x *PlatUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlatUser) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *PlatUser) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *PlatUser) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *PlatUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PlatUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlatUser) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

// PlatRole 平台角色
type PlatRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 创建时间
	CreateTime int64 `protobuf:"varint,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// 创建时间
	UpdateTime int64 `protobuf:"varint,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// 角色名称
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PlatRole) Reset() {
	*x = PlatRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_service_platform_v1_types_platform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatRole) ProtoMessage() {}

func (x *PlatRole) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_service_platform_v1_types_platform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatRole.ProtoReflect.Descriptor instead.
func (*PlatRole) Descriptor() ([]byte, []int) {
	return file_api_rockim_service_platform_v1_types_platform_proto_rawDescGZIP(), []int{1}
}

func (x *PlatRole) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlatRole) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *PlatRole) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *PlatRole) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// PlatResource 平台资源
type PlatResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 创建时间
	CreateTime int64 `protobuf:"varint,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// 创建时间
	UpdateTime int64 `protobuf:"varint,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// 分类
	Category v1.PlatResourceCategory `protobuf:"varint,4,opt,name=category,proto3,enum=rockim.shared.enums.v1.PlatResourceCategory" json:"category,omitempty"`
	// 资源名称
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// 上级ID
	ParentId string `protobuf:"bytes,6,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// URL
	Url string `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	// ICON
	Icon string `protobuf:"bytes,8,opt,name=icon,proto3" json:"icon,omitempty"`
	// 排序号
	Order int32 `protobuf:"varint,9,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *PlatResource) Reset() {
	*x = PlatResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rockim_service_platform_v1_types_platform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatResource) ProtoMessage() {}

func (x *PlatResource) ProtoReflect() protoreflect.Message {
	mi := &file_api_rockim_service_platform_v1_types_platform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatResource.ProtoReflect.Descriptor instead.
func (*PlatResource) Descriptor() ([]byte, []int) {
	return file_api_rockim_service_platform_v1_types_platform_proto_rawDescGZIP(), []int{2}
}

func (x *PlatResource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlatResource) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *PlatResource) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *PlatResource) GetCategory() v1.PlatResourceCategory {
	if x != nil {
		return x.Category
	}
	return v1.PlatResourceCategory_PLAT_RESOURCE_CATEGORY_INVALID
}

func (x *PlatResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlatResource) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *PlatResource) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PlatResource) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *PlatResource) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

var File_api_rockim_service_platform_v1_types_platform_proto protoreflect.FileDescriptor

var file_api_rockim_service_platform_v1_types_platform_proto_rawDesc = []byte{
	0x0a, 0x33, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x25, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1,
	0x01, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x22, 0x70, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x97, 0x02, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x33,
	0x5a, 0x31, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_rockim_service_platform_v1_types_platform_proto_rawDescOnce sync.Once
	file_api_rockim_service_platform_v1_types_platform_proto_rawDescData = file_api_rockim_service_platform_v1_types_platform_proto_rawDesc
)

func file_api_rockim_service_platform_v1_types_platform_proto_rawDescGZIP() []byte {
	file_api_rockim_service_platform_v1_types_platform_proto_rawDescOnce.Do(func() {
		file_api_rockim_service_platform_v1_types_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rockim_service_platform_v1_types_platform_proto_rawDescData)
	})
	return file_api_rockim_service_platform_v1_types_platform_proto_rawDescData
}

var file_api_rockim_service_platform_v1_types_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_rockim_service_platform_v1_types_platform_proto_goTypes = []interface{}{
	(*PlatUser)(nil),             // 0: rockim.service.platform.v1.types.PlatUser
	(*PlatRole)(nil),             // 1: rockim.service.platform.v1.types.PlatRole
	(*PlatResource)(nil),         // 2: rockim.service.platform.v1.types.PlatResource
	(v1.PlatResourceCategory)(0), // 3: rockim.shared.enums.v1.PlatResourceCategory
}
var file_api_rockim_service_platform_v1_types_platform_proto_depIdxs = []int32{
	3, // 0: rockim.service.platform.v1.types.PlatResource.category:type_name -> rockim.shared.enums.v1.PlatResourceCategory
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_rockim_service_platform_v1_types_platform_proto_init() }
func file_api_rockim_service_platform_v1_types_platform_proto_init() {
	if File_api_rockim_service_platform_v1_types_platform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_rockim_service_platform_v1_types_platform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatUser); i {
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
		file_api_rockim_service_platform_v1_types_platform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatRole); i {
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
		file_api_rockim_service_platform_v1_types_platform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatResource); i {
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
			RawDescriptor: file_api_rockim_service_platform_v1_types_platform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_rockim_service_platform_v1_types_platform_proto_goTypes,
		DependencyIndexes: file_api_rockim_service_platform_v1_types_platform_proto_depIdxs,
		MessageInfos:      file_api_rockim_service_platform_v1_types_platform_proto_msgTypes,
	}.Build()
	File_api_rockim_service_platform_v1_types_platform_proto = out.File
	file_api_rockim_service_platform_v1_types_platform_proto_rawDesc = nil
	file_api_rockim_service_platform_v1_types_platform_proto_goTypes = nil
	file_api_rockim_service_platform_v1_types_platform_proto_depIdxs = nil
}
