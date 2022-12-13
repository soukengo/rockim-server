package biz

import (
	"context"
	"rockim/api/rockim/service/platform/v1/types"
)

var ()

// PlatRoleRepo is a PlatRole repo.
type PlatRoleRepo interface {
	ListByIds(ctx context.Context, ids []string) ([]*types.PlatRole, error)
	ListResourceId(ctx context.Context, ids []string) ([]string, error)
}

// PlatRoleUseCase is a PlatRole use case.
type PlatRoleUseCase struct {
	repo PlatRoleRepo
}

// NewPlatRoleUseCase new a PlatRole use case.
func NewPlatRoleUseCase(repo PlatRoleRepo) *PlatRoleUseCase {
	return &PlatRoleUseCase{repo: repo}
}

func (uc *PlatRoleUseCase) ListByIds(ctx context.Context, ids []string) (res []*types.PlatRole, err error) {
	return uc.repo.ListByIds(ctx, ids)
}
func (uc *PlatRoleUseCase) ListResourceId(ctx context.Context, ids []string) ([]string, error) {
	return uc.repo.ListResourceId(ctx, ids)
}
