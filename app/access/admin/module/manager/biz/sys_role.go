package biz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rockimserver/apis/rockim/shared"
	"rockimserver/app/access/admin/module/manager/biz/options"
	"rockimserver/app/access/admin/module/manager/biz/types"
	"rockimserver/pkg/util/copier"
	"time"
)

type SysRoleRepo interface {
	Create(context.Context, *types.SysRole) error
	Update(context.Context, *types.SysRole) error
	Delete(context.Context, string) error

	FindById(context.Context, string) (*types.SysRole, error)
	Paging(context.Context, *options.SysRolePagingOptions) ([]*types.SysRole, *shared.Paginated, error)
	ListResourceId(ctx context.Context, roleIds []string) ([]string, error)
	SaveResourceId(ctx context.Context, roleId string, resourceIds []string) error
}

type SysRoleUseCase struct {
	repo SysRoleRepo
}

func NewSysRoleUseCase(repo SysRoleRepo) *SysRoleUseCase {
	return &SysRoleUseCase{repo: repo}
}

func (uc *SysRoleUseCase) Create(ctx context.Context, req *options.SysRoleCreateOptions) (err error) {
	record := &types.SysRole{Name: req.Options.Name}
	record.Id = primitive.NewObjectID().Hex()
	record.CreateTime = time.Now().UnixMilli()
	return uc.repo.Create(ctx, record)
}
func (uc *SysRoleUseCase) Update(ctx context.Context, req *options.SysRoleUpdateOptions) (err error) {
	record, err := uc.repo.FindById(ctx, req.Id)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(record, req.Options, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return
	}
	record.UpdateTime = time.Now().UnixMilli()
	return uc.repo.Update(ctx, record)
}
func (uc *SysRoleUseCase) Delete(ctx context.Context, id string) (err error) {
	return uc.repo.Delete(ctx, id)
}
func (uc *SysRoleUseCase) Paging(ctx context.Context, req *options.SysRolePagingOptions) ([]*types.SysRole, *shared.Paginated, error) {
	return uc.repo.Paging(ctx, req)
}
func (uc *SysRoleUseCase) ListResourceId(ctx context.Context, roleId string) ([]string, error) {
	return uc.repo.ListResourceId(ctx, []string{roleId})
}
func (uc *SysRoleUseCase) SaveResourceId(ctx context.Context, roleId string, resourceIds []string) error {
	return uc.repo.SaveResourceId(ctx, roleId, resourceIds)
}
