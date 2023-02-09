// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.0
// source: rockim/api/admin/manager/v1/sys_resource.proto

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

const OperationSysResourceAPICreate = "/rockim.api.admin.manager.v1.SysResourceAPI/Create"
const OperationSysResourceAPIDelete = "/rockim.api.admin.manager.v1.SysResourceAPI/Delete"
const OperationSysResourceAPIListTree = "/rockim.api.admin.manager.v1.SysResourceAPI/ListTree"
const OperationSysResourceAPIUpdate = "/rockim.api.admin.manager.v1.SysResourceAPI/Update"

type SysResourceAPIHTTPServer interface {
	Create(context.Context, *SysResourceCreateRequest) (*SysResourceCreateResponse, error)
	Delete(context.Context, *SysResourceDeleteRequest) (*SysResourceDeleteResponse, error)
	ListTree(context.Context, *SysResourceTreeRequest) (*SysResourceTreeResponse, error)
	Update(context.Context, *SysResourceUpdateRequest) (*SysResourceUpdateResponse, error)
}

func RegisterSysResourceAPIHTTPServer(s *http.Server, srv SysResourceAPIHTTPServer) {
	r := s.Route("/")
	r.POST("/api/manager/v1/sys/resource/create", _SysResourceAPI_Create0_HTTP_Handler(srv))
	r.POST("/api/manager/v1/sys/resource/update", _SysResourceAPI_Update0_HTTP_Handler(srv))
	r.POST("/api/manager/v1/sys/resource/delete", _SysResourceAPI_Delete0_HTTP_Handler(srv))
	r.POST("/api/manager/v1/sys/resource/tree", _SysResourceAPI_ListTree0_HTTP_Handler(srv))
}

func _SysResourceAPI_Create0_HTTP_Handler(srv SysResourceAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysResourceCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysResourceAPICreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*SysResourceCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysResourceCreateResponse)
		return ctx.Result(200, reply)
	}
}

func _SysResourceAPI_Update0_HTTP_Handler(srv SysResourceAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysResourceUpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysResourceAPIUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*SysResourceUpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysResourceUpdateResponse)
		return ctx.Result(200, reply)
	}
}

func _SysResourceAPI_Delete0_HTTP_Handler(srv SysResourceAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysResourceDeleteRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysResourceAPIDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*SysResourceDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysResourceDeleteResponse)
		return ctx.Result(200, reply)
	}
}

func _SysResourceAPI_ListTree0_HTTP_Handler(srv SysResourceAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysResourceTreeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysResourceAPIListTree)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTree(ctx, req.(*SysResourceTreeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysResourceTreeResponse)
		return ctx.Result(200, reply)
	}
}

type SysResourceAPIHTTPClient interface {
	Create(ctx context.Context, req *SysResourceCreateRequest, opts ...http.CallOption) (rsp *SysResourceCreateResponse, err error)
	Delete(ctx context.Context, req *SysResourceDeleteRequest, opts ...http.CallOption) (rsp *SysResourceDeleteResponse, err error)
	ListTree(ctx context.Context, req *SysResourceTreeRequest, opts ...http.CallOption) (rsp *SysResourceTreeResponse, err error)
	Update(ctx context.Context, req *SysResourceUpdateRequest, opts ...http.CallOption) (rsp *SysResourceUpdateResponse, err error)
}

type SysResourceAPIHTTPClientImpl struct {
	cc *http.Client
}

func NewSysResourceAPIHTTPClient(client *http.Client) SysResourceAPIHTTPClient {
	return &SysResourceAPIHTTPClientImpl{client}
}

func (c *SysResourceAPIHTTPClientImpl) Create(ctx context.Context, in *SysResourceCreateRequest, opts ...http.CallOption) (*SysResourceCreateResponse, error) {
	var out SysResourceCreateResponse
	pattern := "/api/manager/v1/sys/resource/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysResourceAPICreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysResourceAPIHTTPClientImpl) Delete(ctx context.Context, in *SysResourceDeleteRequest, opts ...http.CallOption) (*SysResourceDeleteResponse, error) {
	var out SysResourceDeleteResponse
	pattern := "/api/manager/v1/sys/resource/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysResourceAPIDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysResourceAPIHTTPClientImpl) ListTree(ctx context.Context, in *SysResourceTreeRequest, opts ...http.CallOption) (*SysResourceTreeResponse, error) {
	var out SysResourceTreeResponse
	pattern := "/api/manager/v1/sys/resource/tree"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysResourceAPIListTree))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysResourceAPIHTTPClientImpl) Update(ctx context.Context, in *SysResourceUpdateRequest, opts ...http.CallOption) (*SysResourceUpdateResponse, error) {
	var out SysResourceUpdateResponse
	pattern := "/api/manager/v1/sys/resource/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysResourceAPIUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}