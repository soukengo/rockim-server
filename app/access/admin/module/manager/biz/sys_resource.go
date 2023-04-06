package biz

import (
	"context"
	"github.com/soukengo/gopkg/util/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rockimserver/app/access/admin/module/manager/biz/options"
	"rockimserver/app/access/admin/module/manager/biz/types"
	"time"
)

type SysResourceRepo interface {
	Create(context.Context, *types.SysResource) error
	Update(context.Context, *types.SysResource) error
	Delete(context.Context, string) error

	FindById(context.Context, string) (*types.SysResource, error)
	List(context.Context) ([]*types.SysResource, error)
	ListByIds(context.Context, []string) ([]*types.SysResource, error)
}

type SysResourceUseCase struct {
	repo SysResourceRepo
}

func NewSysResourceUseCase(repo SysResourceRepo) *SysResourceUseCase {
	return &SysResourceUseCase{repo: repo}
}

func (uc *SysResourceUseCase) Create(ctx context.Context, opts *options.SysResourceCreateOptions) (err error) {
	resource := &types.SysResource{}
	err = copier.Copy(resource, opts.Options)
	if err != nil {
		return
	}
	resource.Id = primitive.NewObjectID().Hex()
	resource.CreateTime = time.Now().UnixMilli()
	return uc.repo.Create(ctx, resource)
}
func (uc *SysResourceUseCase) Update(ctx context.Context, opts *options.SysResourceUpdateOptions) (err error) {
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
func (uc *SysResourceUseCase) Delete(ctx context.Context, id string) (err error) {
	return uc.repo.Delete(ctx, id)
}
func (uc *SysResourceUseCase) List(ctx context.Context) (res []*types.SysResource, err error) {
	return uc.repo.List(ctx)
}
