package manager

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
)

type PlatResourceRepo interface {
	Create(context.Context, *v1.PlatResourceCreateRequest) error
	Update(context.Context, *v1.PlatResourceUpdateRequest) error
	Delete(context.Context, *v1.PlatResourceDeleteRequest) error

	List(context.Context) ([]*types.PlatResource, error)
	ListByIds(context.Context, []string) ([]*types.PlatResource, error)
}

type PlatResourceUseCase struct {
	repo PlatResourceRepo
}

func NewPlatResourceUseCase(repo PlatResourceRepo) *PlatResourceUseCase {
	return &PlatResourceUseCase{repo: repo}
}

func (uc *PlatResourceUseCase) Create(ctx context.Context, req *v1.PlatResourceCreateRequest) (err error) {
	return uc.repo.Create(ctx, req)
}
func (uc *PlatResourceUseCase) Update(ctx context.Context, req *v1.PlatResourceUpdateRequest) (err error) {
	return uc.repo.Update(ctx, req)
}
func (uc *PlatResourceUseCase) Delete(ctx context.Context, req *v1.PlatResourceDeleteRequest) (err error) {
	return uc.repo.Delete(ctx, req)
}
func (uc *PlatResourceUseCase) List(ctx context.Context) (res []*types.PlatResource, err error) {
	return uc.repo.List(ctx)
}
