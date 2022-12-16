// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.0
// source: api/rockim/admin/manager/v1/plat_resource.proto

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

const OperationPlatResourceAPICreate = "/rockim.admin.manager.v1.PlatResourceAPI/Create"
const OperationPlatResourceAPIDelete = "/rockim.admin.manager.v1.PlatResourceAPI/Delete"
const OperationPlatResourceAPIListTree = "/rockim.admin.manager.v1.PlatResourceAPI/ListTree"
const OperationPlatResourceAPIUpdate = "/rockim.admin.manager.v1.PlatResourceAPI/Update"

type PlatResourceAPIHTTPServer interface {
	Create(context.Context, *PlatResourceCreateRequest) (*PlatResourceCreateResponse, error)
	Delete(context.Context, *PlatResourceDeleteRequest) (*PlatResourceDeleteResponse, error)
	ListTree(context.Context, *PlatResourceTreeRequest) (*PlatResourceTreeResponse, error)
	Update(context.Context, *PlatResourceUpdateRequest) (*PlatResourceUpdateResponse, error)
}

func RegisterPlatResourceAPIHTTPServer(s *http.Server, srv PlatResourceAPIHTTPServer) {
	r := s.Route("/")
	r.POST("/api/manager/v1/platform/resource/create", _PlatResourceAPI_Create0_HTTP_Handler(srv))
	r.POST("/api/manager/v1/platform/resource/update", _PlatResourceAPI_Update0_HTTP_Handler(srv))
	r.POST("/api/manager/v1/platform/resource/delete", _PlatResourceAPI_Delete0_HTTP_Handler(srv))
	r.POST("/api/manager/v1/platform/resource/tree", _PlatResourceAPI_ListTree0_HTTP_Handler(srv))
}

func _PlatResourceAPI_Create0_HTTP_Handler(srv PlatResourceAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlatResourceCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPlatResourceAPICreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*PlatResourceCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlatResourceCreateResponse)
		return ctx.Result(200, reply)
	}
}

func _PlatResourceAPI_Update0_HTTP_Handler(srv PlatResourceAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlatResourceUpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPlatResourceAPIUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*PlatResourceUpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlatResourceUpdateResponse)
		return ctx.Result(200, reply)
	}
}

func _PlatResourceAPI_Delete0_HTTP_Handler(srv PlatResourceAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlatResourceDeleteRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPlatResourceAPIDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*PlatResourceDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlatResourceDeleteResponse)
		return ctx.Result(200, reply)
	}
}

func _PlatResourceAPI_ListTree0_HTTP_Handler(srv PlatResourceAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlatResourceTreeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPlatResourceAPIListTree)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTree(ctx, req.(*PlatResourceTreeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlatResourceTreeResponse)
		return ctx.Result(200, reply)
	}
}

type PlatResourceAPIHTTPClient interface {
	Create(ctx context.Context, req *PlatResourceCreateRequest, opts ...http.CallOption) (rsp *PlatResourceCreateResponse, err error)
	Delete(ctx context.Context, req *PlatResourceDeleteRequest, opts ...http.CallOption) (rsp *PlatResourceDeleteResponse, err error)
	ListTree(ctx context.Context, req *PlatResourceTreeRequest, opts ...http.CallOption) (rsp *PlatResourceTreeResponse, err error)
	Update(ctx context.Context, req *PlatResourceUpdateRequest, opts ...http.CallOption) (rsp *PlatResourceUpdateResponse, err error)
}

type PlatResourceAPIHTTPClientImpl struct {
	cc *http.Client
}

func NewPlatResourceAPIHTTPClient(client *http.Client) PlatResourceAPIHTTPClient {
	return &PlatResourceAPIHTTPClientImpl{client}
}

func (c *PlatResourceAPIHTTPClientImpl) Create(ctx context.Context, in *PlatResourceCreateRequest, opts ...http.CallOption) (*PlatResourceCreateResponse, error) {
	var out PlatResourceCreateResponse
	pattern := "/api/manager/v1/platform/resource/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPlatResourceAPICreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PlatResourceAPIHTTPClientImpl) Delete(ctx context.Context, in *PlatResourceDeleteRequest, opts ...http.CallOption) (*PlatResourceDeleteResponse, error) {
	var out PlatResourceDeleteResponse
	pattern := "/api/manager/v1/platform/resource/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPlatResourceAPIDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PlatResourceAPIHTTPClientImpl) ListTree(ctx context.Context, in *PlatResourceTreeRequest, opts ...http.CallOption) (*PlatResourceTreeResponse, error) {
	var out PlatResourceTreeResponse
	pattern := "/api/manager/v1/platform/resource/tree"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPlatResourceAPIListTree))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PlatResourceAPIHTTPClientImpl) Update(ctx context.Context, in *PlatResourceUpdateRequest, opts ...http.CallOption) (*PlatResourceUpdateResponse, error) {
	var out PlatResourceUpdateResponse
	pattern := "/api/manager/v1/platform/resource/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPlatResourceAPIUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
