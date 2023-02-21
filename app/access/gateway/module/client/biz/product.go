package biz

import (
	"context"
	"rockimserver/apis/rockim/service/platform/v1/types"
)

type ProductRepo interface {
	FindById(ctx context.Context, productId string) (*types.Product, error)
}

type ProductUseCase struct {
	repo ProductRepo
}

func NewProductUseCase(repo ProductRepo) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) FindById(ctx context.Context, productId string) (*types.Product, error) {
	return uc.repo.FindById(ctx, productId)
}
