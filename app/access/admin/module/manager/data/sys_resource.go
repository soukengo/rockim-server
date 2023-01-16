package data

import (
	"context"
	"rockim/app/access/admin/module/manager/biz"
	"rockim/app/access/admin/module/manager/biz/types"
	"rockim/app/access/admin/module/manager/data/database"
)

type platResourceRepo struct {
	db database.SysResourceData
}

func NewSysResourceRepo(db database.SysResourceData) biz.SysResourceRepo {
	return &platResourceRepo{db: db}
}

func (r *platResourceRepo) Create(ctx context.Context, record *types.SysResource) (err error) {
	err = r.db.Create(ctx, record)
	return
}

func (r *platResourceRepo) Update(ctx context.Context, record *types.SysResource) (err error) {
	err = r.db.Update(ctx, record)
	return
}

func (r *platResourceRepo) Delete(ctx context.Context, id string) (err error) {
	err = r.db.Delete(ctx, id)
	return
}

func (r *platResourceRepo) FindById(ctx context.Context, id string) (*types.SysResource, error) {
	return r.db.FindByID(ctx, id)
}

func (r *platResourceRepo) List(ctx context.Context) ([]*types.SysResource, error) {
	return r.db.List(ctx)
}
func (r *platResourceRepo) ListByIds(ctx context.Context, ids []string) ([]*types.SysResource, error) {
	return r.db.ListByIds(ctx, ids)
}
