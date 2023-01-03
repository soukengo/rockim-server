package data

import (
	"context"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/logic/platform/biz"
	"rockim/app/logic/platform/data/database"
)

type tenantResourceRepo struct {
	db *database.TenantResourceData
}

func NewTenantResourceRepo(db *database.TenantResourceData) biz.TenantResourceRepo {
	return &tenantResourceRepo{db: db}
}

func (r *tenantResourceRepo) Create(ctx context.Context, resource *types.TenantResource) error {
	return r.db.Create(ctx, resource)
}

func (r *tenantResourceRepo) Update(ctx context.Context, resource *types.TenantResource) error {
	return r.db.Update(ctx, resource)
}

func (r *tenantResourceRepo) Delete(ctx context.Context, id string) error {
	return r.db.Delete(ctx, id)
}

func (r *tenantResourceRepo) FindById(ctx context.Context, id string) (*types.TenantResource, error) {
	return r.db.FindByID(ctx, id)
}

func (r *tenantResourceRepo) List(ctx context.Context) ([]*types.TenantResource, error) {
	return r.db.List(ctx)
}
func (r *tenantResourceRepo) ListByIds(ctx context.Context, ids []string) ([]*types.TenantResource, error) {
	return r.db.ListByIds(ctx, ids)
}
