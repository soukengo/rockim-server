// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.0
// source: api/rockim/admin/manager/v1/plat_role.proto

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

const OperationPlatRoleAPICreate = "/rockim.admin.manager.v1.PlatRoleAPI/Create"
const OperationPlatRoleAPIDelete = "/rockim.admin.manager.v1.PlatRoleAPI/Delete"
const OperationPlatRoleAPIListResourceId = "/rockim.admin.manager.v1.PlatRoleAPI/ListResourceId"
const OperationPlatRoleAPIPaging = "/rockim.admin.manager.v1.PlatRoleAPI/Paging"
const OperationPlatRoleAPISaveResourceId = "/rockim.admin.manager.v1.PlatRoleAPI/SaveResourceId"
const OperationPlatRoleAPIUpdate = "/rockim.admin.manager.v1.PlatRoleAPI/Update"

type PlatRoleAPIHTTPServer interface {
	Create(context.Context, *PlatRoleCreateRequest) (*PlatRoleCreateResponse, error)
	Delete(context.Context, *PlatRoleDeleteRequest) (*PlatRoleDeleteResponse, error)
	ListResourceId(context.Context, *PlatRoleResourceIdListRequest) (*PlatRoleResourceIdListResponse, error)
	Paging(context.Context, *PlatRolePagingRequest) (*PlatRolePagingResponse, error)
	SaveResourceId(context.Context, *PlatRoleResourceIdSaveRequest) (*PlatRoleResourceIdSaveResponse, error)
	Update(context.Context, *PlatRoleUpdateRequest) (*PlatRoleUpdateResponse, error)
}

func RegisterPlatRoleAPIHTTPServer(s *http.Server, srv PlatRoleAPIHTTPServer) {
	r := s.Route("/")
	r.POST("/api/manager/v1/platform/role/create", _PlatRoleAPI_Create1_HTTP_Handler(srv))
	r.POST("/api/manager/v1/platform/role/update", _PlatRoleAPI_Update1_HTTP_Handler(srv))
	r.POST("/api/manager/v1/platform/role/delete", _PlatRoleAPI_Delete1_HTTP_Handler(srv))
	r.POST("/api/manager/v1/platform/role/paging", _PlatRoleAPI_Paging0_HTTP_Handler(srv))
	r.POST("/api/manager/v1/platform/role/resource_id/list", _PlatRoleAPI_ListResourceId0_HTTP_Handler(srv))
	r.POST("/api/manager/v1/platform/role/resource_id/save", _PlatRoleAPI_SaveResourceId0_HTTP_Handler(srv))
}

func _PlatRoleAPI_Create1_HTTP_Handler(srv PlatRoleAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlatRoleCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPlatRoleAPICreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*PlatRoleCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlatRoleCreateResponse)
		return ctx.Result(200, reply)
	}
}

func _PlatRoleAPI_Update1_HTTP_Handler(srv PlatRoleAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlatRoleUpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPlatRoleAPIUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*PlatRoleUpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlatRoleUpdateResponse)
		return ctx.Result(200, reply)
	}
}

func _PlatRoleAPI_Delete1_HTTP_Handler(srv PlatRoleAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlatRoleDeleteRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPlatRoleAPIDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*PlatRoleDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlatRoleDeleteResponse)
		return ctx.Result(200, reply)
	}
}

func _PlatRoleAPI_Paging0_HTTP_Handler(srv PlatRoleAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlatRolePagingRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPlatRoleAPIPaging)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Paging(ctx, req.(*PlatRolePagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlatRolePagingResponse)
		return ctx.Result(200, reply)
	}
}

func _PlatRoleAPI_ListResourceId0_HTTP_Handler(srv PlatRoleAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlatRoleResourceIdListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPlatRoleAPIListResourceId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListResourceId(ctx, req.(*PlatRoleResourceIdListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlatRoleResourceIdListResponse)
		return ctx.Result(200, reply)
	}
}

func _PlatRoleAPI_SaveResourceId0_HTTP_Handler(srv PlatRoleAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlatRoleResourceIdSaveRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPlatRoleAPISaveResourceId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveResourceId(ctx, req.(*PlatRoleResourceIdSaveRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlatRoleResourceIdSaveResponse)
		return ctx.Result(200, reply)
	}
}

type PlatRoleAPIHTTPClient interface {
	Create(ctx context.Context, req *PlatRoleCreateRequest, opts ...http.CallOption) (rsp *PlatRoleCreateResponse, err error)
	Delete(ctx context.Context, req *PlatRoleDeleteRequest, opts ...http.CallOption) (rsp *PlatRoleDeleteResponse, err error)
	ListResourceId(ctx context.Context, req *PlatRoleResourceIdListRequest, opts ...http.CallOption) (rsp *PlatRoleResourceIdListResponse, err error)
	Paging(ctx context.Context, req *PlatRolePagingRequest, opts ...http.CallOption) (rsp *PlatRolePagingResponse, err error)
	SaveResourceId(ctx context.Context, req *PlatRoleResourceIdSaveRequest, opts ...http.CallOption) (rsp *PlatRoleResourceIdSaveResponse, err error)
	Update(ctx context.Context, req *PlatRoleUpdateRequest, opts ...http.CallOption) (rsp *PlatRoleUpdateResponse, err error)
}

type PlatRoleAPIHTTPClientImpl struct {
	cc *http.Client
}

func NewPlatRoleAPIHTTPClient(client *http.Client) PlatRoleAPIHTTPClient {
	return &PlatRoleAPIHTTPClientImpl{client}
}

func (c *PlatRoleAPIHTTPClientImpl) Create(ctx context.Context, in *PlatRoleCreateRequest, opts ...http.CallOption) (*PlatRoleCreateResponse, error) {
	var out PlatRoleCreateResponse
	pattern := "/api/manager/v1/platform/role/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPlatRoleAPICreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PlatRoleAPIHTTPClientImpl) Delete(ctx context.Context, in *PlatRoleDeleteRequest, opts ...http.CallOption) (*PlatRoleDeleteResponse, error) {
	var out PlatRoleDeleteResponse
	pattern := "/api/manager/v1/platform/role/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPlatRoleAPIDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PlatRoleAPIHTTPClientImpl) ListResourceId(ctx context.Context, in *PlatRoleResourceIdListRequest, opts ...http.CallOption) (*PlatRoleResourceIdListResponse, error) {
	var out PlatRoleResourceIdListResponse
	pattern := "/api/manager/v1/platform/role/resource_id/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPlatRoleAPIListResourceId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PlatRoleAPIHTTPClientImpl) Paging(ctx context.Context, in *PlatRolePagingRequest, opts ...http.CallOption) (*PlatRolePagingResponse, error) {
	var out PlatRolePagingResponse
	pattern := "/api/manager/v1/platform/role/paging"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPlatRoleAPIPaging))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PlatRoleAPIHTTPClientImpl) SaveResourceId(ctx context.Context, in *PlatRoleResourceIdSaveRequest, opts ...http.CallOption) (*PlatRoleResourceIdSaveResponse, error) {
	var out PlatRoleResourceIdSaveResponse
	pattern := "/api/manager/v1/platform/role/resource_id/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPlatRoleAPISaveResourceId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PlatRoleAPIHTTPClientImpl) Update(ctx context.Context, in *PlatRoleUpdateRequest, opts ...http.CallOption) (*PlatRoleUpdateResponse, error) {
	var out PlatRoleUpdateResponse
	pattern := "/api/manager/v1/platform/role/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPlatRoleAPIUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}