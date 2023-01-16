package data

import (
	"context"
	"rockim/api/rockim/shared"
	"rockim/app/access/admin/module/manager/biz"
	"rockim/app/access/admin/module/manager/biz/options"
	"rockim/app/access/admin/module/manager/biz/types"
	"rockim/app/access/admin/module/manager/data/database"
)

type sysRoleRepo struct {
	db *database.SysRoleData
}

func NewSysRoleRepo(db *database.SysRoleData) biz.SysRoleRepo {
	return &sysRoleRepo{db: db}
}

func (r *sysRoleRepo) Create(ctx context.Context, record *types.SysRole) (err error) {
	err = r.db.Create(ctx, record)
	return
}

func (r *sysRoleRepo) Update(ctx context.Context, record *types.SysRole) (err error) {
	err = r.db.Update(ctx, record)
	return
}

func (r *sysRoleRepo) Delete(ctx context.Context, id string) (err error) {
	err = r.db.Delete(ctx, id)
	return
}

func (r *sysRoleRepo) FindById(ctx context.Context, id string) (*types.SysRole, error) {
	return r.db.FindByID(ctx, id)
}

func (r *sysRoleRepo) Paging(ctx context.Context, request *options.SysRolePagingOptions) (list []*types.SysRole, p *shared.Paginated, err error) {
	return r.db.Paging(ctx, request)
}

func (r *sysRoleRepo) ListResourceId(ctx context.Context, roleIds []string) ([]string, error) {
	return r.db.ListResourceId(ctx, roleIds)
}

func (r *sysRoleRepo) SaveResourceId(ctx context.Context, roleId string, resourceIds []string) (err error) {
	err = r.db.SaveResourceId(ctx, roleId, resourceIds)
	return
}
