// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.0
// source: rockim/api/client/v1/user.proto

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

const OperationUserAPIFind = "/rockim.api.client.v1.UserAPI/Find"

type UserAPIHTTPServer interface {
	Find(context.Context, *UserFindRequest) (*UserFindResponse, error)
}

func RegisterUserAPIHTTPServer(s *http.Server, srv UserAPIHTTPServer) {
	r := s.Route("/")
	r.POST("/client/v1/user/find", _UserAPI_Find0_HTTP_Handler(srv))
}

func _UserAPI_Find0_HTTP_Handler(srv UserAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserFindRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPIFind)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Find(ctx, req.(*UserFindRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserFindResponse)
		return ctx.Result(200, reply)
	}
}

type UserAPIHTTPClient interface {
	Find(ctx context.Context, req *UserFindRequest, opts ...http.CallOption) (rsp *UserFindResponse, err error)
}

type UserAPIHTTPClientImpl struct {
	cc *http.Client
}

func NewUserAPIHTTPClient(client *http.Client) UserAPIHTTPClient {
	return &UserAPIHTTPClientImpl{client}
}

func (c *UserAPIHTTPClientImpl) Find(ctx context.Context, in *UserFindRequest, opts ...http.CallOption) (*UserFindResponse, error) {
	var out UserFindResponse
	pattern := "/client/v1/user/find"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserAPIFind))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
