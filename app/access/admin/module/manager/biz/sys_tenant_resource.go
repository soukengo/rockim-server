package biz

import (
	"context"
	"github.com/soukengo/gopkg/util/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rockimserver/app/access/admin/module/manager/biz/options"
	"rockimserver/app/access/admin/module/manager/biz/types"
	"time"
)

type SysTenantResourceRepo interface {
	Create(context.Context, *types.SysTenantResource) error
	Update(context.Context, *types.SysTenantResource) error
	Delete(context.Context, string) error

	FindById(context.Context, string) (*types.SysTenantResource, error)
	List(context.Context) ([]*types.SysTenantResource, error)
	ListByIds(context.Context, []string) ([]*types.SysTenantResource, error)
}

type SysTenantResourceUseCase struct {
	repo SysTenantResourceRepo
}

func NewSysTenantResourceUseCase(repo SysTenantResourceRepo) *SysTenantResourceUseCase {
	return &SysTenantResourceUseCase{repo: repo}
}

func (uc *SysTenantResourceUseCase) Create(ctx context.Context, opts *options.SysTenantResourceCreateOptions) (err error) {
	resource := &types.SysTenantResource{}
	err = copier.Copy(resource, opts.Options)
	if err != nil {
		return
	}
	resource.Id = primitive.NewObjectID().Hex()
	resource.CreateTime = time.Now().UnixMilli()
	return uc.repo.Create(ctx, resource)
}
func (uc *SysTenantResourceUseCase) Update(ctx context.Context, opts *options.SysTenantResourceUpdateOptions) (err error) {
	resource, err := uc.repo.FindById(ctx, opts.Id)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(resource, opts.Options, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return
	}
	resource.UpdateTime = time.Now().UnixMilli()
	return uc.repo.Update(ctx, resource)
}
func (uc *SysTenantResourceUseCase) Delete(ctx context.Context, id string) (err error) {
	return uc.repo.Delete(ctx, id)
}
func (uc *SysTenantResourceUseCase) List(ctx context.Context) (res []*types.SysTenantResource, err error) {
	return uc.repo.List(ctx)
}
