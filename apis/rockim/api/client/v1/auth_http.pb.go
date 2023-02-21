// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.0
// source: rockim/api/client/v1/auth.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAuthAPILogin = "/rockim.api.client.v1.AuthAPI/Login"

type AuthAPIHTTPServer interface {
	// Login Login 登录
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

func RegisterAuthAPIHTTPServer(s *http.Server, srv AuthAPIHTTPServer) {
	r := s.Route("/")
	r.POST("/client/v1/auth/login", _AuthAPI_Login0_HTTP_Handler(srv))
}

func _AuthAPI_Login0_HTTP_Handler(srv AuthAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthAPILogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResponse)
		return ctx.Result(200, reply)
	}
}

type AuthAPIHTTPClient interface {
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginResponse, err error)
}

type AuthAPIHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthAPIHTTPClient(client *http.Client) AuthAPIHTTPClient {
	return &AuthAPIHTTPClientImpl{client}
}

func (c *AuthAPIHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginResponse, error) {
	var out LoginResponse
	pattern := "/client/v1/auth/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthAPILogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
