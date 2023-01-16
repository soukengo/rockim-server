// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.0
// source: api/rockim/admin/manager/v1/tenant.proto

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

const OperationTenantAPICreate = "/rockim.admin.manager.v1.TenantAPI/Create"
const OperationTenantAPIPaging = "/rockim.admin.manager.v1.TenantAPI/Paging"
const OperationTenantAPIUpdate = "/rockim.admin.manager.v1.TenantAPI/Update"

type TenantAPIHTTPServer interface {
	Create(context.Context, *TenantCreateRequest) (*TenantCreateResponse, error)
	Paging(context.Context, *TenantPagingRequest) (*TenantPagingResponse, error)
	Update(context.Context, *TenantUpdateRequest) (*TenantUpdateResponse, error)
}

func RegisterTenantAPIHTTPServer(s *http.Server, srv TenantAPIHTTPServer) {
	r := s.Route("/")
	r.POST("/api/manager/v1/tenant/tenant/create", _TenantAPI_Create4_HTTP_Handler(srv))
	r.POST("/api/manager/v1/tenant/tenant/update", _TenantAPI_Update4_HTTP_Handler(srv))
	r.POST("/api/manager/v1/tenant/tenant/paging", _TenantAPI_Paging2_HTTP_Handler(srv))
}

func _TenantAPI_Create4_HTTP_Handler(srv TenantAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TenantCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantAPICreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*TenantCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TenantCreateResponse)
		return ctx.Result(200, reply)
	}
}

func _TenantAPI_Update4_HTTP_Handler(srv TenantAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TenantUpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantAPIUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*TenantUpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TenantUpdateResponse)
		return ctx.Result(200, reply)
	}
}

func _TenantAPI_Paging2_HTTP_Handler(srv TenantAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TenantPagingRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantAPIPaging)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Paging(ctx, req.(*TenantPagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TenantPagingResponse)
		return ctx.Result(200, reply)
	}
}

type TenantAPIHTTPClient interface {
	Create(ctx context.Context, req *TenantCreateRequest, opts ...http.CallOption) (rsp *TenantCreateResponse, err error)
	Paging(ctx context.Context, req *TenantPagingRequest, opts ...http.CallOption) (rsp *TenantPagingResponse, err error)
	Update(ctx context.Context, req *TenantUpdateRequest, opts ...http.CallOption) (rsp *TenantUpdateResponse, err error)
}

type TenantAPIHTTPClientImpl struct {
	cc *http.Client
}

func NewTenantAPIHTTPClient(client *http.Client) TenantAPIHTTPClient {
	return &TenantAPIHTTPClientImpl{client}
}

func (c *TenantAPIHTTPClientImpl) Create(ctx context.Context, in *TenantCreateRequest, opts ...http.CallOption) (*TenantCreateResponse, error) {
	var out TenantCreateResponse
	pattern := "/api/manager/v1/tenant/tenant/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTenantAPICreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TenantAPIHTTPClientImpl) Paging(ctx context.Context, in *TenantPagingRequest, opts ...http.CallOption) (*TenantPagingResponse, error) {
	var out TenantPagingResponse
	pattern := "/api/manager/v1/tenant/tenant/paging"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTenantAPIPaging))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TenantAPIHTTPClientImpl) Update(ctx context.Context, in *TenantUpdateRequest, opts ...http.CallOption) (*TenantUpdateResponse, error) {
	var out TenantUpdateResponse
	pattern := "/api/manager/v1/tenant/tenant/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTenantAPIUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
