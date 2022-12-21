package manager

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
)

type PlatUserRepo interface {
	Create(context.Context, *v1.PlatUserCreateRequest) error
	Update(context.Context, *v1.PlatUserUpdateRequest) error
	Delete(context.Context, *v1.PlatUserDeleteRequest) error

	Paging(context.Context, *v1.PlatUserPagingRequest) (*v1.PlatUserPagingResponse, error)
	FindByAccount(ctx context.Context, account string) (*types.PlatUser, error)

	ListRoleId(ctx context.Context, userId string) ([]string, error)
	SaveRoleId(ctx context.Context, userId string, roleIds []string) error
}

type PlatUserUseCase struct {
	repo PlatUserRepo
}

func NewPlatUserUseCase(repo PlatUserRepo) *PlatUserUseCase {
	return &PlatUserUseCase{repo: repo}
}

func (uc *PlatUserUseCase) Create(ctx context.Context, req *v1.PlatUserCreateRequest) (err error) {
	return uc.repo.Create(ctx, req)
}
func (uc *PlatUserUseCase) Update(ctx context.Context, req *v1.PlatUserUpdateRequest) (err error) {
	return uc.repo.Update(ctx, req)
}
func (uc *PlatUserUseCase) Delete(ctx context.Context, req *v1.PlatUserDeleteRequest) (err error) {
	return uc.repo.Delete(ctx, req)
}
func (uc *PlatUserUseCase) Paging(ctx context.Context, req *v1.PlatUserPagingRequest) (*v1.PlatUserPagingResponse, error) {
	return uc.repo.Paging(ctx, req)
}
func (uc *PlatUserUseCase) ListRoleId(ctx context.Context, userId string) ([]string, error) {
	return uc.repo.ListRoleId(ctx, userId)
}

func (uc *PlatUserUseCase) SaveRoleId(ctx context.Context, userId string, roleIds []string) error {
	return uc.repo.SaveRoleId(ctx, userId, roleIds)
}
