package data

import (
	"context"
	v1 "rockimserver/apis/rockim/service/platform/v1"
	"rockimserver/app/access/admin/module/manager/biz"
)

type tenantRepo struct {
	ac v1.TenantAPIClient
}

func NewTenantRepo(uac v1.TenantAPIClient) biz.TenantRepo {
	return &tenantRepo{ac: uac}
}

func (r *tenantRepo) Create(ctx context.Context, request *v1.TenantCreateRequest) (err error) {
	_, err = r.ac.Create(ctx, request)
	return
}

func (r *tenantRepo) Update(ctx context.Context, request *v1.TenantUpdateRequest) (err error) {
	_, err = r.ac.Update(ctx, request)
	return
}

func (r *tenantRepo) Paging(ctx context.Context, request *v1.TenantPagingRequest) (reply *v1.TenantPagingResponse, err error) {
	return r.ac.Paging(ctx, request)
}
