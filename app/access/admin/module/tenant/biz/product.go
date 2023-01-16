package biz

import (
	"context"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/access/admin/module/tenant/biz/options"
)

type ProductRepo interface {
	Create(ctx context.Context, opts *options.ProductCreateOptions) error
	Update(ctx context.Context, opts *options.ProductUpdateOptions) error
	ListByTenant(ctx context.Context, tenantId string) (res []*types.Product, err error)
}

type ProductUseCase struct {
	repo ProductRepo
}

func NewProductUseCase(repo ProductRepo) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) Create(ctx context.Context, opts *options.ProductCreateOptions) (err error) {
	return uc.repo.Create(ctx, opts)
}

func (uc *ProductUseCase) Update(ctx context.Context, opts *options.ProductUpdateOptions) (err error) {
	return uc.repo.Update(ctx, opts)
}

func (uc *ProductUseCase) ListByTenant(ctx context.Context, req *options.ProductPagingOptions) (ret []*types.Product, err error) {
	return uc.repo.ListByTenant(ctx, req.TenantId)
}
