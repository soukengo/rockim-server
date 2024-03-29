// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.0
// source: rockim/api/client/v1/protocol/http/product.proto

package http

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

const OperationProductAPIFetchConfig = "/rockim.api.client.v1.protocol.http.ProductAPI/FetchConfig"

type ProductAPIHTTPServer interface {
	// FetchConfig FetchConfig 获取全局配置
	FetchConfig(context.Context, *ConfigFetchRequest) (*ConfigFetchResponse, error)
}

func RegisterProductAPIHTTPServer(s *http.Server, srv ProductAPIHTTPServer) {
	r := s.Route("/")
	r.POST("/client/v1/product/config/fetch", _ProductAPI_FetchConfig0_HTTP_Handler(srv))
}

func _ProductAPI_FetchConfig0_HTTP_Handler(srv ProductAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ConfigFetchRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductAPIFetchConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FetchConfig(ctx, req.(*ConfigFetchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ConfigFetchResponse)
		return ctx.Result(200, reply)
	}
}

type ProductAPIHTTPClient interface {
	FetchConfig(ctx context.Context, req *ConfigFetchRequest, opts ...http.CallOption) (rsp *ConfigFetchResponse, err error)
}

type ProductAPIHTTPClientImpl struct {
	cc *http.Client
}

func NewProductAPIHTTPClient(client *http.Client) ProductAPIHTTPClient {
	return &ProductAPIHTTPClientImpl{client}
}

func (c *ProductAPIHTTPClientImpl) FetchConfig(ctx context.Context, in *ConfigFetchRequest, opts ...http.CallOption) (*ConfigFetchResponse, error) {
	var out ConfigFetchResponse
	pattern := "/client/v1/product/config/fetch"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProductAPIFetchConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
