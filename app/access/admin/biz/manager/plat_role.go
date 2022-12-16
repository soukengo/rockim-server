package manager

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
)

type PlatRoleRepo interface {
	Create(context.Context, *v1.PlatRoleCreateRequest) error
	Update(context.Context, *v1.PlatRoleUpdateRequest) error
	Delete(context.Context, *v1.PlatRoleDeleteRequest) error

	Paging(context.Context, *v1.PlatRolePagingRequest) (*v1.PlatRolePagingResponse, error)
	ListResourceId(ctx context.Context, roleIds []string) ([]string, error)
}

type PlatRoleUseCase struct {
	repo PlatRoleRepo
}

func NewPlatRoleUseCase(repo PlatRoleRepo) *PlatRoleUseCase {
	return &PlatRoleUseCase{repo: repo}
}

func (uc *PlatRoleUseCase) Create(ctx context.Context, req *v1.PlatRoleCreateRequest) (err error) {
	return uc.repo.Create(ctx, req)
}
func (uc *PlatRoleUseCase) Update(ctx context.Context, req *v1.PlatRoleUpdateRequest) (err error) {
	return uc.repo.Update(ctx, req)
}
func (uc *PlatRoleUseCase) Delete(ctx context.Context, req *v1.PlatRoleDeleteRequest) (err error) {
	return uc.repo.Delete(ctx, req)
}
func (uc *PlatRoleUseCase) Paging(ctx context.Context, req *v1.PlatRolePagingRequest) (*v1.PlatRolePagingResponse, error) {
	return uc.repo.Paging(ctx, req)
}
func (uc *PlatRoleUseCase) ListResourceId(ctx context.Context, roleId string) ([]string, error) {
	return uc.repo.ListResourceId(ctx, []string{roleId})
}
