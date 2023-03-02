// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: rockim/service/user/v1/user.proto

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

// UserAPIClient is the client API for UserAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAPIClient interface {
	// Register 注册用户
	Register(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
	// Find 查询用户
	Find(ctx context.Context, in *UserFindRequest, opts ...grpc.CallOption) (*UserFindResponse, error)
	// FindByAccount 根据账号查询用户
	FindByAccount(ctx context.Context, in *UserFindByAccountRequest, opts ...grpc.CallOption) (*UserFindByAccountResponse, error)
	// FindUid 查询用户ID
	FindUid(ctx context.Context, in *UserIdFindRequest, opts ...grpc.CallOption) (*UserIdFindResponse, error)
}

type userAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAPIClient(cc grpc.ClientConnInterface) UserAPIClient {
	return &userAPIClient{cc}
}

func (c *userAPIClient) Register(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, "/rockim.service.user.v1.UserAPI/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) Find(ctx context.Context, in *UserFindRequest, opts ...grpc.CallOption) (*UserFindResponse, error) {
	out := new(UserFindResponse)
	err := c.cc.Invoke(ctx, "/rockim.service.user.v1.UserAPI/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) FindByAccount(ctx context.Context, in *UserFindByAccountRequest, opts ...grpc.CallOption) (*UserFindByAccountResponse, error) {
	out := new(UserFindByAccountResponse)
	err := c.cc.Invoke(ctx, "/rockim.service.user.v1.UserAPI/FindByAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) FindUid(ctx context.Context, in *UserIdFindRequest, opts ...grpc.CallOption) (*UserIdFindResponse, error) {
	out := new(UserIdFindResponse)
	err := c.cc.Invoke(ctx, "/rockim.service.user.v1.UserAPI/FindUid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAPIServer is the server API for UserAPI service.
// All implementations must embed UnimplementedUserAPIServer
// for forward compatibility
type UserAPIServer interface {
	// Register 注册用户
	Register(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
	// Find 查询用户
	Find(context.Context, *UserFindRequest) (*UserFindResponse, error)
	// FindByAccount 根据账号查询用户
	FindByAccount(context.Context, *UserFindByAccountRequest) (*UserFindByAccountResponse, error)
	// FindUid 查询用户ID
	FindUid(context.Context, *UserIdFindRequest) (*UserIdFindResponse, error)
	mustEmbedUnimplementedUserAPIServer()
}

// UnimplementedUserAPIServer must be embedded to have forward compatible implementations.
type UnimplementedUserAPIServer struct {
}

func (UnimplementedUserAPIServer) Register(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserAPIServer) Find(context.Context, *UserFindRequest) (*UserFindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedUserAPIServer) FindByAccount(context.Context, *UserFindByAccountRequest) (*UserFindByAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByAccount not implemented")
}
func (UnimplementedUserAPIServer) FindUid(context.Context, *UserIdFindRequest) (*UserIdFindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUid not implemented")
}
func (UnimplementedUserAPIServer) mustEmbedUnimplementedUserAPIServer() {}

// UnsafeUserAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAPIServer will
// result in compilation errors.
type UnsafeUserAPIServer interface {
	mustEmbedUnimplementedUserAPIServer()
}

func RegisterUserAPIServer(s grpc.ServiceRegistrar, srv UserAPIServer) {
	s.RegisterService(&UserAPI_ServiceDesc, srv)
}

func _UserAPI_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rockim.service.user.v1.UserAPI/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).Register(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rockim.service.user.v1.UserAPI/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).Find(ctx, req.(*UserFindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_FindByAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFindByAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).FindByAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rockim.service.user.v1.UserAPI/FindByAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).FindByAccount(ctx, req.(*UserFindByAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_FindUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdFindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).FindUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rockim.service.user.v1.UserAPI/FindUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).FindUid(ctx, req.(*UserIdFindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAPI_ServiceDesc is the grpc.ServiceDesc for UserAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rockim.service.user.v1.UserAPI",
	HandlerType: (*UserAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserAPI_Register_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _UserAPI_Find_Handler,
		},
		{
			MethodName: "FindByAccount",
			Handler:    _UserAPI_FindByAccount_Handler,
		},
		{
			MethodName: "FindUid",
			Handler:    _UserAPI_FindUid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rockim/service/user/v1/user.proto",
}
