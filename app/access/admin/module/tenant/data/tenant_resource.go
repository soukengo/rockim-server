package data

import (
	"context"
	"rockim/app/access/admin/module/tenant/biz"
	"rockim/app/access/admin/module/tenant/biz/types"
	"rockim/app/access/admin/module/tenant/data/database"
)

type tenantResourceRepo struct {
	db database.SysTenantResourceData
}

func NewTenantResourceRepo(db database.SysTenantResourceData) biz.SysTenantResourceRepo {
	return &tenantResourceRepo{db: db}
}

func (r *tenantResourceRepo) List(ctx context.Context) ([]*types.SysTenantResource, error) {
	return r.db.List(ctx)
}
