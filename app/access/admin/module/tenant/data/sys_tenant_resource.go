package data

import (
	"context"
	"rockim/app/access/admin/module/tenant/biz"
	"rockim/app/access/admin/module/tenant/biz/types"
	"rockim/app/access/admin/module/tenant/data/database"
)

type sysTenantResourceRepo struct {
	db *database.SysTenantResourceData
}

func NewSysTenantResourceRepo(db *database.SysTenantResourceData) biz.SysTenantResourceRepo {
	return &sysTenantResourceRepo{db: db}
}

func (r *sysTenantResourceRepo) List(ctx context.Context) ([]*types.SysTenantResource, error) {
	return r.db.List(ctx)
}
