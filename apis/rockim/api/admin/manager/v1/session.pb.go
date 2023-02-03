// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: rockim/api/admin/manager/v1/session.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	types "rockimserver/apis/rockim/api/admin/manager/v1/types"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SessionCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SessionCheckRequest) Reset() {
	*x = SessionCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_admin_manager_v1_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionCheckRequest) ProtoMessage() {}

func (x *SessionCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_admin_manager_v1_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionCheckRequest.ProtoReflect.Descriptor instead.
func (*SessionCheckRequest) Descriptor() ([]byte, []int) {
	return file_rockim_api_admin_manager_v1_session_proto_rawDescGZIP(), []int{0}
}

type SessionCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *types.SessionUser `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *SessionCheckResponse) Reset() {
	*x = SessionCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_admin_manager_v1_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionCheckResponse) ProtoMessage() {}

func (x *SessionCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_admin_manager_v1_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionCheckResponse.ProtoReflect.Descriptor instead.
func (*SessionCheckResponse) Descriptor() ([]byte, []int) {
	return file_rockim_api_admin_manager_v1_session_proto_rawDescGZIP(), []int{1}
}

func (x *SessionCheckResponse) GetSession() *types.SessionUser {
	if x != nil {
		return x.Session
	}
	return nil
}

type SessionListResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SessionListResourceRequest) Reset() {
	*x = SessionListResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_admin_manager_v1_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionListResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionListResourceRequest) ProtoMessage() {}

func (x *SessionListResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_admin_manager_v1_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionListResourceRequest.ProtoReflect.Descriptor instead.
func (*SessionListResourceRequest) Descriptor() ([]byte, []int) {
	return file_rockim_api_admin_manager_v1_session_proto_rawDescGZIP(), []int{2}
}

type SessionListResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*types.SysResourceTree `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *SessionListResourceResponse) Reset() {
	*x = SessionListResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rockim_api_admin_manager_v1_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionListResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionListResourceResponse) ProtoMessage() {}

func (x *SessionListResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rockim_api_admin_manager_v1_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionListResourceResponse.ProtoReflect.Descriptor instead.
func (*SessionListResourceResponse) Descriptor() ([]byte, []int) {
	return file_rockim_api_admin_manager_v1_session_proto_rawDescGZIP(), []int{3}
}

func (x *SessionListResourceResponse) GetList() []*types.SysResourceTree {
	if x != nil {
		return x.List
	}
	return nil
}

var File_rockim_api_admin_manager_v1_session_proto protoreflect.FileDescriptor

var file_rockim_api_admin_manager_v1_session_proto_rawDesc = []byte{
	0x0a, 0x29, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x72, 0x6f, 0x63,
	0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x60, 0x0a, 0x14, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x1c, 0x0a,
	0x1a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x65, 0x0a, 0x1b, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x73,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x32, 0xd7, 0x02, 0x0a, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x50,
	0x49, 0x12, 0x96, 0x01, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x30, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0xaf, 0x01, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x37, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x22, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x32, 0x5a, 0x30,
	0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rockim_api_admin_manager_v1_session_proto_rawDescOnce sync.Once
	file_rockim_api_admin_manager_v1_session_proto_rawDescData = file_rockim_api_admin_manager_v1_session_proto_rawDesc
)

func file_rockim_api_admin_manager_v1_session_proto_rawDescGZIP() []byte {
	file_rockim_api_admin_manager_v1_session_proto_rawDescOnce.Do(func() {
		file_rockim_api_admin_manager_v1_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_rockim_api_admin_manager_v1_session_proto_rawDescData)
	})
	return file_rockim_api_admin_manager_v1_session_proto_rawDescData
}

var file_rockim_api_admin_manager_v1_session_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rockim_api_admin_manager_v1_session_proto_goTypes = []interface{}{
	(*SessionCheckRequest)(nil),         // 0: rockim.api.admin.manager.v1.SessionCheckRequest
	(*SessionCheckResponse)(nil),        // 1: rockim.api.admin.manager.v1.SessionCheckResponse
	(*SessionListResourceRequest)(nil),  // 2: rockim.api.admin.manager.v1.SessionListResourceRequest
	(*SessionListResourceResponse)(nil), // 3: rockim.api.admin.manager.v1.SessionListResourceResponse
	(*types.SessionUser)(nil),           // 4: rockim.api.admin.manager.v1.types.SessionUser
	(*types.SysResourceTree)(nil),       // 5: rockim.api.admin.manager.v1.types.SysResourceTree
}
var file_rockim_api_admin_manager_v1_session_proto_depIdxs = []int32{
	4, // 0: rockim.api.admin.manager.v1.SessionCheckResponse.session:type_name -> rockim.api.admin.manager.v1.types.SessionUser
	5, // 1: rockim.api.admin.manager.v1.SessionListResourceResponse.list:type_name -> rockim.api.admin.manager.v1.types.SysResourceTree
	0, // 2: rockim.api.admin.manager.v1.SessionAPI.Check:input_type -> rockim.api.admin.manager.v1.SessionCheckRequest
	2, // 3: rockim.api.admin.manager.v1.SessionAPI.ListResource:input_type -> rockim.api.admin.manager.v1.SessionListResourceRequest
	1, // 4: rockim.api.admin.manager.v1.SessionAPI.Check:output_type -> rockim.api.admin.manager.v1.SessionCheckResponse
	3, // 5: rockim.api.admin.manager.v1.SessionAPI.ListResource:output_type -> rockim.api.admin.manager.v1.SessionListResourceResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rockim_api_admin_manager_v1_session_proto_init() }
func file_rockim_api_admin_manager_v1_session_proto_init() {
	if File_rockim_api_admin_manager_v1_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rockim_api_admin_manager_v1_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionCheckRequest); i {
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
		file_rockim_api_admin_manager_v1_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionCheckResponse); i {
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
		file_rockim_api_admin_manager_v1_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionListResourceRequest); i {
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
		file_rockim_api_admin_manager_v1_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionListResourceResponse); i {
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
			RawDescriptor: file_rockim_api_admin_manager_v1_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rockim_api_admin_manager_v1_session_proto_goTypes,
		DependencyIndexes: file_rockim_api_admin_manager_v1_session_proto_depIdxs,
		MessageInfos:      file_rockim_api_admin_manager_v1_session_proto_msgTypes,
	}.Build()
	File_rockim_api_admin_manager_v1_session_proto = out.File
	file_rockim_api_admin_manager_v1_session_proto_rawDesc = nil
	file_rockim_api_admin_manager_v1_session_proto_goTypes = nil
	file_rockim_api_admin_manager_v1_session_proto_depIdxs = nil
}
