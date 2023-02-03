package biz

import (
	"context"
	v1 "rockimserver/apis/rockim/service/platform/v1"
)

type TenantRepo interface {
	Create(context.Context, *v1.TenantCreateRequest) error
	Update(context.Context, *v1.TenantUpdateRequest) error

	Paging(context.Context, *v1.TenantPagingRequest) (*v1.TenantPagingResponse, error)
}

type TenantUseCase struct {
	repo TenantRepo
}

func NewTenantUseCase(repo TenantRepo) *TenantUseCase {
	return &TenantUseCase{repo: repo}
}

func (uc *TenantUseCase) Create(ctx context.Context, req *v1.TenantCreateRequest) (err error) {
	return uc.repo.Create(ctx, req)
}
func (uc *TenantUseCase) Update(ctx context.Context, req *v1.TenantUpdateRequest) (err error) {
	return uc.repo.Update(ctx, req)
}

func (uc *TenantUseCase) Paging(ctx context.Context, req *v1.TenantPagingRequest) (*v1.TenantPagingResponse, error) {
	return uc.repo.Paging(ctx, req)
}
