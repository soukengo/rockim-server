package data

import (
	"context"
	v1 "rockimserver/apis/rockim/service/platform/v1"
	"rockimserver/apis/rockim/service/platform/v1/types"
	"rockimserver/app/access/admin/module/tenant/biz"
)

type tenantRepo struct {
	ac v1.TenantAPIClient
}

func NewTenantRepo(ac v1.TenantAPIClient) biz.TenantRepo {
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
