// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: rockim/service/group/v1/chatroom.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatRoomAPIClient is the client API for ChatRoomAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatRoomAPIClient interface {
	// Create 创建聊天室
	Create(ctx context.Context, in *ChatRoomCreateRequest, opts ...grpc.CallOption) (*ChatRoomCreateResponse, error)
	// Dismiss 解散聊天室
	Dismiss(ctx context.Context, in *ChatRoomDismissRequest, opts ...grpc.CallOption) (*ChatRoomDismissResponse, error)
}

type chatRoomAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewChatRoomAPIClient(cc grpc.ClientConnInterface) ChatRoomAPIClient {
	return &chatRoomAPIClient{cc}
}

func (c *chatRoomAPIClient) Create(ctx context.Context, in *ChatRoomCreateRequest, opts ...grpc.CallOption) (*ChatRoomCreateResponse, error) {
	out := new(ChatRoomCreateResponse)
	err := c.cc.Invoke(ctx, "/rockim.service.group.v1.ChatRoomAPI/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatRoomAPIClient) Dismiss(ctx context.Context, in *ChatRoomDismissRequest, opts ...grpc.CallOption) (*ChatRoomDismissResponse, error) {
	out := new(ChatRoomDismissResponse)
	err := c.cc.Invoke(ctx, "/rockim.service.group.v1.ChatRoomAPI/Dismiss", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatRoomAPIServer is the server API for ChatRoomAPI service.
// All implementations must embed UnimplementedChatRoomAPIServer
// for forward compatibility
type ChatRoomAPIServer interface {
	// Create 创建聊天室
	Create(context.Context, *ChatRoomCreateRequest) (*ChatRoomCreateResponse, error)
	// Dismiss 解散聊天室
	Dismiss(context.Context, *ChatRoomDismissRequest) (*ChatRoomDismissResponse, error)
	mustEmbedUnimplementedChatRoomAPIServer()
}

// UnimplementedChatRoomAPIServer must be embedded to have forward compatible implementations.
type UnimplementedChatRoomAPIServer struct {
}

func (UnimplementedChatRoomAPIServer) Create(context.Context, *ChatRoomCreateRequest) (*ChatRoomCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedChatRoomAPIServer) Dismiss(context.Context, *ChatRoomDismissRequest) (*ChatRoomDismissResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dismiss not implemented")
}
func (UnimplementedChatRoomAPIServer) mustEmbedUnimplementedChatRoomAPIServer() {}

// UnsafeChatRoomAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatRoomAPIServer will
// result in compilation errors.
type UnsafeChatRoomAPIServer interface {
	mustEmbedUnimplementedChatRoomAPIServer()
}

func RegisterChatRoomAPIServer(s grpc.ServiceRegistrar, srv ChatRoomAPIServer) {
	s.RegisterService(&ChatRoomAPI_ServiceDesc, srv)
}

func _ChatRoomAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRoomCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRoomAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rockim.service.group.v1.ChatRoomAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRoomAPIServer).Create(ctx, req.(*ChatRoomCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatRoomAPI_Dismiss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRoomDismissRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRoomAPIServer).Dismiss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rockim.service.group.v1.ChatRoomAPI/Dismiss",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRoomAPIServer).Dismiss(ctx, req.(*ChatRoomDismissRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatRoomAPI_ServiceDesc is the grpc.ServiceDesc for ChatRoomAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatRoomAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rockim.service.group.v1.ChatRoomAPI",
	HandlerType: (*ChatRoomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ChatRoomAPI_Create_Handler,
		},
		{
			MethodName: "Dismiss",
			Handler:    _ChatRoomAPI_Dismiss_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rockim/service/group/v1/chatroom.proto",
}
