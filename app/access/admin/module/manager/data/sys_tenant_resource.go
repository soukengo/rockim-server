package data

import (
	"context"
	"rockim/app/access/admin/module/manager/biz"
	"rockim/app/access/admin/module/manager/biz/types"
	"rockim/app/access/admin/module/manager/data/database"
)

type tenantResourceRepo struct {
	db database.SysTenantResourceData
}

func NewSysTenantResourceRepo(db database.SysTenantResourceData) biz.SysTenantResourceRepo {
	return &tenantResourceRepo{db: db}
}

func (r *tenantResourceRepo) Create(ctx context.Context, record *types.SysTenantResource) (err error) {
	err = r.db.Create(ctx, record)
	return
}

func (r *tenantResourceRepo) Update(ctx context.Context, record *types.SysTenantResource) (err error) {
	err = r.db.Update(ctx, record)
	return
}

func (r *tenantResourceRepo) Delete(ctx context.Context, id string) (err error) {
	err = r.db.Delete(ctx, id)
	return
}
func (r *tenantResourceRepo) FindById(ctx context.Context, id string) (*types.SysTenantResource, error) {
	return r.db.FindByID(ctx, id)
}

func (r *tenantResourceRepo) List(ctx context.Context) ([]*types.SysTenantResource, error) {
	return r.db.List(ctx)
}
func (r *tenantResourceRepo) ListByIds(ctx context.Context, ids []string) ([]*types.SysTenantResource, error) {
	return r.db.ListByIds(ctx, ids)
}
