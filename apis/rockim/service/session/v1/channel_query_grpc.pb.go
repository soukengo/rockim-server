// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: rockim/service/session/v1/channel_query.proto

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

// ChannelQueryAPIClient is the client API for ChannelQueryAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelQueryAPIClient interface {
	// CheckOnline 检查在线状态
	CheckOnline(ctx context.Context, in *OnlineCheckRequest, opts ...grpc.CallOption) (*OnlineCheckResponse, error)
	// BatchCheckOnline 批量检查在线状态
	BatchCheckOnline(ctx context.Context, in *OnlineBatchCheckRequest, opts ...grpc.CallOption) (*OnlineBatchCheckResponse, error)
	// ListUser 获取指定用户的连接ID
	ListUser(ctx context.Context, in *UserChannelListRequest, opts ...grpc.CallOption) (*UserChannelListResponse, error)
}

type channelQueryAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelQueryAPIClient(cc grpc.ClientConnInterface) ChannelQueryAPIClient {
	return &channelQueryAPIClient{cc}
}

func (c *channelQueryAPIClient) CheckOnline(ctx context.Context, in *OnlineCheckRequest, opts ...grpc.CallOption) (*OnlineCheckResponse, error) {
	out := new(OnlineCheckResponse)
	err := c.cc.Invoke(ctx, "/rockim.service.user.v1.ChannelQueryAPI/CheckOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelQueryAPIClient) BatchCheckOnline(ctx context.Context, in *OnlineBatchCheckRequest, opts ...grpc.CallOption) (*OnlineBatchCheckResponse, error) {
	out := new(OnlineBatchCheckResponse)
	err := c.cc.Invoke(ctx, "/rockim.service.user.v1.ChannelQueryAPI/BatchCheckOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelQueryAPIClient) ListUser(ctx context.Context, in *UserChannelListRequest, opts ...grpc.CallOption) (*UserChannelListResponse, error) {
	out := new(UserChannelListResponse)
	err := c.cc.Invoke(ctx, "/rockim.service.user.v1.ChannelQueryAPI/ListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelQueryAPIServer is the server API for ChannelQueryAPI service.
// All implementations must embed UnimplementedChannelQueryAPIServer
// for forward compatibility
type ChannelQueryAPIServer interface {
	// CheckOnline 检查在线状态
	CheckOnline(context.Context, *OnlineCheckRequest) (*OnlineCheckResponse, error)
	// BatchCheckOnline 批量检查在线状态
	BatchCheckOnline(context.Context, *OnlineBatchCheckRequest) (*OnlineBatchCheckResponse, error)
	// ListUser 获取指定用户的连接ID
	ListUser(context.Context, *UserChannelListRequest) (*UserChannelListResponse, error)
	mustEmbedUnimplementedChannelQueryAPIServer()
}

// UnimplementedChannelQueryAPIServer must be embedded to have forward compatible implementations.
type UnimplementedChannelQueryAPIServer struct {
}

func (UnimplementedChannelQueryAPIServer) CheckOnline(context.Context, *OnlineCheckRequest) (*OnlineCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckOnline not implemented")
}
func (UnimplementedChannelQueryAPIServer) BatchCheckOnline(context.Context, *OnlineBatchCheckRequest) (*OnlineBatchCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCheckOnline not implemented")
}
func (UnimplementedChannelQueryAPIServer) ListUser(context.Context, *UserChannelListRequest) (*UserChannelListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedChannelQueryAPIServer) mustEmbedUnimplementedChannelQueryAPIServer() {}

// UnsafeChannelQueryAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelQueryAPIServer will
// result in compilation errors.
type UnsafeChannelQueryAPIServer interface {
	mustEmbedUnimplementedChannelQueryAPIServer()
}

func RegisterChannelQueryAPIServer(s grpc.ServiceRegistrar, srv ChannelQueryAPIServer) {
	s.RegisterService(&ChannelQueryAPI_ServiceDesc, srv)
}

func _ChannelQueryAPI_CheckOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelQueryAPIServer).CheckOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rockim.service.user.v1.ChannelQueryAPI/CheckOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelQueryAPIServer).CheckOnline(ctx, req.(*OnlineCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelQueryAPI_BatchCheckOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineBatchCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelQueryAPIServer).BatchCheckOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rockim.service.user.v1.ChannelQueryAPI/BatchCheckOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelQueryAPIServer).BatchCheckOnline(ctx, req.(*OnlineBatchCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelQueryAPI_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelQueryAPIServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rockim.service.user.v1.ChannelQueryAPI/ListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelQueryAPIServer).ListUser(ctx, req.(*UserChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelQueryAPI_ServiceDesc is the grpc.ServiceDesc for ChannelQueryAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelQueryAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rockim.service.user.v1.ChannelQueryAPI",
	HandlerType: (*ChannelQueryAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckOnline",
			Handler:    _ChannelQueryAPI_CheckOnline_Handler,
		},
		{
			MethodName: "BatchCheckOnline",
			Handler:    _ChannelQueryAPI_BatchCheckOnline_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _ChannelQueryAPI_ListUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rockim/service/session/v1/channel_query.proto",
}