package manager

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
)

type TenantResourceRepo interface {
	Create(context.Context, *v1.TenantResourceCreateRequest) error
	Update(context.Context, *v1.TenantResourceUpdateRequest) error
	Delete(context.Context, *v1.TenantResourceDeleteRequest) error

	List(context.Context) ([]*types.TenantResource, error)
	ListByIds(context.Context, []string) ([]*types.TenantResource, error)
}

type TenantResourceUseCase struct {
	repo TenantResourceRepo
}

func NewTenantResourceUseCase(repo TenantResourceRepo) *TenantResourceUseCase {
	return &TenantResourceUseCase{repo: repo}
}

func (uc *TenantResourceUseCase) Create(ctx context.Context, req *v1.TenantResourceCreateRequest) (err error) {
	return uc.repo.Create(ctx, req)
}
func (uc *TenantResourceUseCase) Update(ctx context.Context, req *v1.TenantResourceUpdateRequest) (err error) {
	return uc.repo.Update(ctx, req)
}
func (uc *TenantResourceUseCase) Delete(ctx context.Context, req *v1.TenantResourceDeleteRequest) (err error) {
	return uc.repo.Delete(ctx, req)
}
func (uc *TenantResourceUseCase) List(ctx context.Context) (res []*types.TenantResource, err error) {
	return uc.repo.List(ctx)
}
