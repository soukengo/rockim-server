package data

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/access/admin/module/tenant/biz"
)

type productRepo struct {
	ac v1.ProductAPIClient
}

func NewProductRepo(ac v1.ProductAPIClient) biz.ProductRepo {
	return &productRepo{ac: ac}
}

func (r *productRepo) Create(ctx context.Context, req *biz.ProductCreateRequest) (err error) {
	_, err = r.ac.Create(ctx, &v1.ProductCreateRequest{
		TenantId: req.TenantId,
		Name:     req.Name,
	})
	return
}

func (r *productRepo) Update(ctx context.Context, req *biz.ProductUpdateRequest) (err error) {
	_, err = r.ac.Update(ctx, &v1.ProductUpdateRequest{
		Id:   req.Id,
		Name: req.Name,
	})
	return
}

func (r *productRepo) ListByTenant(ctx context.Context, tenantId string) (res []*types.Product, err error) {
	ret, err := r.ac.ListByTenant(ctx, &v1.ProductListByTenantRequest{TenantId: tenantId})
	if err != nil {
		return
	}
	return ret.List, nil
}
