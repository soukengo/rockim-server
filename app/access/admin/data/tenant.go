package data

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/access/admin/biz/manager"
	"rockim/app/access/admin/biz/tenant"
)

type managerTenantRepo struct {
	ac v1.TenantAPIClient
}

func NewManagerTenantRepo(uac v1.TenantAPIClient) manager.TenantRepo {
	return &managerTenantRepo{ac: uac}
}

func (r *managerTenantRepo) Create(ctx context.Context, request *v1.TenantCreateRequest) (err error) {
	_, err = r.ac.Create(ctx, request)
	return
}

func (r *managerTenantRepo) Update(ctx context.Context, request *v1.TenantUpdateRequest) (err error) {
	_, err = r.ac.Update(ctx, request)
	return
}

func (r *managerTenantRepo) Paging(ctx context.Context, request *v1.TenantPagingRequest) (reply *v1.TenantPagingResponse, err error) {
	return r.ac.Paging(ctx, request)
}

type tenantRepo struct {
	ac v1.TenantAPIClient
}

func NewTenantRepo(ac v1.TenantAPIClient) tenant.TenantRepo {
	return &tenantRepo{ac: ac}
}

func (r *tenantRepo) FindByEmail(ctx context.Context, email string) (res *types.Tenant, err error) {
	ret, err := r.ac.Find(ctx, &v1.TenantFindRequest{
		Email: email,
	})
	if err != nil {
		return
	}
	return ret.Tenant, nil
}
