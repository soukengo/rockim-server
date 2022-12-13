package biz

import (
	"context"
	"rockim/api/rockim/service/platform/v1/types"
)

var ()

// PlatResourceRepo is a PlatResource repo.
type PlatResourceRepo interface {
	ListByIds(ctx context.Context, ids []string) ([]*types.PlatResource, error)
}

// PlatResourceUseCase is a PlatResource use case.
type PlatResourceUseCase struct {
	repo PlatResourceRepo
}

// NewPlatResourceUseCase new a PlatResource use case.
func NewPlatResourceUseCase(repo PlatResourceRepo) *PlatResourceUseCase {
	return &PlatResourceUseCase{repo: repo}
}

func (uc *PlatResourceUseCase) ListByIds(ctx context.Context, ids []string) (res []*types.PlatResource, err error) {
	return uc.repo.ListByIds(ctx, ids)
}
