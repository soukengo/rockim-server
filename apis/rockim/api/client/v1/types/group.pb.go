// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/api/client/v1/types/group.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	enums "rockimserver/apis/rockim/shared/enums"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Group 群组
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分类
	Category enums.Group_Category `protobuf:"varint,1,opt,name=category,proto3,enum=rockim.shared.enums.Group_Category" json:"category,omitempty"`
	// 所属应用
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// 自定义群组ID，由接入方指定（不填的话SDK服务端自动生成）
	CustomGroupId string `protobuf:"bytes,3,opt,name=custom_group_id,json=customGroupId,proto3" json:"custom_group_id,omitempty"`
	// 群组名称
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// 图标
	IconUrl string `protobuf:"bytes,5,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
	// 自定义字段
	Fields map[string]string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 状态
	Status enums.Group_Status `protobuf:"varint,7,opt,name=status,proto3,enum=rockim.shared.enums.Group_Status" json:"status,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_types_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_types_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_types_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetCategory() enums.Group_Category {
	if x != nil {
		return x.Category
	}
	return enums.Group_CHATROOM
}

func (x *Group) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Group) GetCustomGroupId() string {
	if x != nil {
		return x.CustomGroupId
	}
	return ""
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *Group) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Group) GetStatus() enums.Group_Status {
	if x != nil {
		return x.Status
	}
	return enums.Group_NORMAL
}

// GroupMember 群组成员
type GroupMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 加入时间
	CreateTime int64 `protobuf:"varint,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// 所属应用
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// 群组id
	GroupId string `protobuf:"bytes,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// 用户账号
	Account string `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	// 角色
	Role enums.Group_MemberRole `protobuf:"varint,5,opt,name=role,proto3,enum=rockim.shared.enums.Group_MemberRole" json:"role,omitempty"`
	// 自定义字段
	Fields map[string]string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GroupMember) Reset() {
	*x = GroupMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_client_v1_types_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMember) ProtoMessage() {}

func (x *GroupMember) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_client_v1_types_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMember.ProtoReflect.Descriptor instead.
func (*GroupMember) Descriptor() ([]byte, []int) {
	return file_rockim_api_client_v1_types_group_proto_rawDescGZIP(), []int{1}
}

func (x *GroupMember) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *GroupMember) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *GroupMember) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GroupMember) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *GroupMember) GetRole() enums.Group_MemberRole {
	if x != nil {
		return x.Role
	}
	return enums.Group_ORDINARY
}

func (x *GroupMember) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

var File_rockim_api_client_v1_types_group_proto protoreflect.FileDescriptor

var file_rockim_api_client_v1_types_group_proto_rawDesc = []byte{
	0x0a, 0x26, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x1f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x02, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x3f, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x23, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x45, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x39, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xc5, 0x02, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x51, 0x5a, 0x32, 0x72,
	0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65,
	0x73, 0xaa, 0x02, 0x1a, 0x52, 0x6f, 0x63, 0x6b, 0x49, 0x4d, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_api_client_v1_types_group_proto_rawDescOnce sync.Once
	file_rockim_api_client_v1_types_group_proto_rawDescData = file_rockim_api_client_v1_types_group_proto_rawDesc
)

func file_rockim_api_client_v1_types_group_proto_rawDescGZIP() []byte {
	file_rockim_api_client_v1_types_group_proto_rawDescOnce.Do(func() {
		file_rockim_api_client_v1_types_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_api_client_v1_types_group_proto_rawDescData)
	})
	return file_rockim_api_client_v1_types_group_proto_rawDescData
}

var file_rockim_api_client_v1_types_group_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rockim_api_client_v1_types_group_proto_goTypes = []interface{}{
	(*Group)(nil),               // 0: rockim.api.client.v1.types.Group
	(*GroupMember)(nil),         // 1: rockim.api.client.v1.types.GroupMember
	nil,                         // 2: rockim.api.client.v1.types.Group.FieldsEntry
	nil,                         // 3: rockim.api.client.v1.types.GroupMember.FieldsEntry
	(enums.Group_Category)(0),   // 4: rockim.shared.enums.Group.Category
	(enums.Group_Status)(0),     // 5: rockim.shared.enums.Group.Status
	(enums.Group_MemberRole)(0), // 6: rockim.shared.enums.Group.MemberRole
}
var file_rockim_api_client_v1_types_group_proto_depIdxs = []int32{
	4, // 0: rockim.api.client.v1.types.Group.category:type_name -> rockim.shared.enums.Group.Category
	2, // 1: rockim.api.client.v1.types.Group.fields:type_name -> rockim.api.client.v1.types.Group.FieldsEntry
	5, // 2: rockim.api.client.v1.types.Group.status:type_name -> rockim.shared.enums.Group.Status
	6, // 3: rockim.api.client.v1.types.GroupMember.role:type_name -> rockim.shared.enums.Group.MemberRole
	3, // 4: rockim.api.client.v1.types.GroupMember.fields:type_name -> rockim.api.client.v1.types.GroupMember.FieldsEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_rockim_api_client_v1_types_group_proto_init() }
func file_rockim_api_client_v1_types_group_proto_init() {
	if File_rockim_api_client_v1_types_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_api_client_v1_types_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_rockim_api_client_v1_types_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMember); i {
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
			RawDescriptor: file_rockim_api_client_v1_types_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rockim_api_client_v1_types_group_proto_goTypes,
		DependencyIndexes: file_rockim_api_client_v1_types_group_proto_depIdxs,
		MessageInfos:      file_rockim_api_client_v1_types_group_proto_msgTypes,
	}.Build()
	File_rockim_api_client_v1_types_group_proto = out.File
	file_rockim_api_client_v1_types_group_proto_rawDesc = nil
	file_rockim_api_client_v1_types_group_proto_goTypes = nil
	file_rockim_api_client_v1_types_group_proto_depIdxs = nil
}
