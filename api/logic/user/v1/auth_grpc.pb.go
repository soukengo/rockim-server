// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: api/logic/user/v1/auth.proto

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

// AuthAPIClient is the client API for AuthAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthAPIClient interface {
	// 登录接口
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthAPIClient(cc grpc.ClientConnInterface) AuthAPIClient {
	return &authAPIClient{cc}
}

func (c *authAPIClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/logic.user.v1.AuthAPI/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthAPIServer is the server API for AuthAPI service.
// All implementations must embed UnimplementedAuthAPIServer
// for forward compatibility
type AuthAPIServer interface {
	// 登录接口
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedAuthAPIServer()
}

// UnimplementedAuthAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAuthAPIServer struct {
}

func (UnimplementedAuthAPIServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthAPIServer) mustEmbedUnimplementedAuthAPIServer() {}

// UnsafeAuthAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthAPIServer will
// result in compilation errors.
type UnsafeAuthAPIServer interface {
	mustEmbedUnimplementedAuthAPIServer()
}

func RegisterAuthAPIServer(s grpc.ServiceRegistrar, srv AuthAPIServer) {
	s.RegisterService(&AuthAPI_ServiceDesc, srv)
}

func _AuthAPI_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.user.v1.AuthAPI/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthAPI_ServiceDesc is the grpc.ServiceDesc for AuthAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logic.user.v1.AuthAPI",
	HandlerType: (*AuthAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthAPI_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/logic/user/v1/auth.proto",
}
