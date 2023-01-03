package data

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/access/admin/biz/manager"
	"rockim/app/access/admin/biz/tenant"
)

type managerTenantResourceRepo struct {
	ac v1.TenantResourceAPIClient
}

func NewManagerTenantResourceRepo(ac v1.TenantResourceAPIClient) manager.TenantResourceRepo {
	return &managerTenantResourceRepo{ac: ac}
}

func (r *managerTenantResourceRepo) Create(ctx context.Context, request *v1.TenantResourceCreateRequest) (err error) {
	_, err = r.ac.Create(ctx, request)
	return
}

func (r *managerTenantResourceRepo) Update(ctx context.Context, request *v1.TenantResourceUpdateRequest) (err error) {
	_, err = r.ac.Update(ctx, request)
	return
}

func (r *managerTenantResourceRepo) Delete(ctx context.Context, request *v1.TenantResourceDeleteRequest) (err error) {
	_, err = r.ac.Delete(ctx, request)
	return
}

func (r *managerTenantResourceRepo) List(ctx context.Context) ([]*types.TenantResource, error) {
	ret, err := r.ac.List(ctx, &v1.TenantResourceListRequest{})
	if err != nil {
		return nil, err
	}
	return ret.List, nil
}
func (r *managerTenantResourceRepo) ListByIds(ctx context.Context, ids []string) ([]*types.TenantResource, error) {
	ret, err := r.ac.ListByIds(ctx, &v1.TenantResourceListByIdsRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	return ret.List, nil
}

type tenantResourceRepo struct {
	ac v1.TenantResourceAPIClient
}

func NewTenantResourceRepo(ac v1.TenantResourceAPIClient) tenant.TenantResourceRepo {
	return &tenantResourceRepo{ac: ac}
}

func (r *tenantResourceRepo) List(ctx context.Context) ([]*types.TenantResource, error) {
	ret, err := r.ac.List(ctx, &v1.TenantResourceListRequest{})
	if err != nil {
		return nil, err
	}
	return ret.List, nil
}
