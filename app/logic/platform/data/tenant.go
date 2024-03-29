package data

import (
	"context"
	"rockimserver/apis/rockim/service/platform/v1/types"
	"rockimserver/apis/rockim/shared"
	"rockimserver/app/logic/platform/biz"
	"rockimserver/app/logic/platform/biz/options"
	"rockimserver/app/logic/platform/data/database"
)

type tenantRepo struct {
	db *database.TenantData
}

func NewTenantRepo(db *database.TenantData) biz.TenantRepo {
	return &tenantRepo{db: db}
}

func (r *tenantRepo) GenID(ctx context.Context) (string, error) {
	return r.db.GenID(ctx)
}

func (r *tenantRepo) Create(ctx context.Context, record *types.Tenant) error {
	return r.db.Create(ctx, record)
}
func (r *tenantRepo) Update(ctx context.Context, record *types.Tenant) error {
	return r.db.Update(ctx, record)
}
func (r *tenantRepo) Delete(ctx context.Context, id string) error {
	return r.db.Delete(ctx, id)
}

func (r *tenantRepo) FindById(ctx context.Context, id string) (*types.Tenant, error) {
	return r.db.FindByID(ctx, id)
}

func (r *tenantRepo) FindIdByEmail(ctx context.Context, account string) (id string, err error) {
	return r.db.FindIdByEmail(ctx, account)
}

func (r *tenantRepo) ListByIds(ctx context.Context, ids []string) ([]*types.Tenant, error) {
	return r.db.ListByIds(ctx, ids)
}

func (r *tenantRepo) Paging(ctx context.Context, opts *options.TenantPagingOptions) (ret []*types.Tenant, paginated *shared.Paginated, err error) {
	return r.db.Paging(ctx, opts)
}
