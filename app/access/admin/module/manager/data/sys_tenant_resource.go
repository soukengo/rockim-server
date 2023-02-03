package data

import (
	"context"
	"rockimserver/app/access/admin/module/manager/biz"
	"rockimserver/app/access/admin/module/manager/biz/types"
	"rockimserver/app/access/admin/module/manager/data/database"
)

type sysTenantResourceRepo struct {
	db *database.SysTenantResourceData
}

func NewSysTenantResourceRepo(db *database.SysTenantResourceData) biz.SysTenantResourceRepo {
	return &sysTenantResourceRepo{db: db}
}

func (r *sysTenantResourceRepo) Create(ctx context.Context, record *types.SysTenantResource) (err error) {
	err = r.db.Create(ctx, record)
	return
}

func (r *sysTenantResourceRepo) Update(ctx context.Context, record *types.SysTenantResource) (err error) {
	err = r.db.Update(ctx, record)
	return
}

func (r *sysTenantResourceRepo) Delete(ctx context.Context, id string) (err error) {
	err = r.db.Delete(ctx, id)
	return
}
func (r *sysTenantResourceRepo) FindById(ctx context.Context, id string) (*types.SysTenantResource, error) {
	return r.db.FindByID(ctx, id)
}

func (r *sysTenantResourceRepo) List(ctx context.Context) ([]*types.SysTenantResource, error) {
	return r.db.List(ctx)
}
func (r *sysTenantResourceRepo) ListByIds(ctx context.Context, ids []string) ([]*types.SysTenantResource, error) {
	return r.db.ListByIds(ctx, ids)
}
