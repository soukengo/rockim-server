package biz

import (
	"context"
	"rockim/api/rockim/service/platform/v1/types"
)

type ProductRepo interface {
	Create(ctx context.Context, record *ProductCreateRequest) error
	Update(ctx context.Context, record *ProductUpdateRequest) error
	ListByTenant(ctx context.Context, tenantId string) (res []*types.Product, err error)
}

// ProductCreateRequest 创建请求
type ProductCreateRequest struct {
	// 所属租户
	TenantId string
	// 应用名称
	Name string
}

// ProductUpdateRequest 更新请求
type ProductUpdateRequest struct {
	// product id
	Id string
	// 应用名称
	Name string
}

type ProductPagingRequest struct {
	TenantId string
}

type ProductUseCase struct {
	repo ProductRepo
}

func NewProductUseCase(repo ProductRepo) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) Create(ctx context.Context, req *ProductCreateRequest) (err error) {
	return uc.repo.Create(ctx, req)
}

func (uc *ProductUseCase) Update(ctx context.Context, req *ProductUpdateRequest) (err error) {
	return uc.repo.Update(ctx, req)
}

func (uc *ProductUseCase) ListByTenant(ctx context.Context, req *ProductPagingRequest) (ret []*types.Product, err error) {
	return uc.repo.ListByTenant(ctx, req.TenantId)
}
