package data

import (
	"context"
	"rockimserver/apis/rockim/service/platform/v1/types"
	"rockimserver/app/logic/platform/biz"
	"rockimserver/app/logic/platform/data/database"
	"rockimserver/pkg/errors"
)

type productRepo struct {
	db *database.ProductData
}

func NewProductRepo(db *database.ProductData) biz.ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) GenID(ctx context.Context) (string, error) {
	return r.db.GenID(ctx)
}

func (r *productRepo) Create(ctx context.Context, record *types.Product) error {
	return r.db.Create(ctx, record)
}

func (r *productRepo) Update(ctx context.Context, record *types.Product) error {
	return r.db.Update(ctx, record)
}

func (r *productRepo) FindById(ctx context.Context, id string) (out *types.Product, err error) {
	out, err = r.db.FindByID(ctx, id)
	if err != nil {
		if errors.IsNotFound(err) {
			err = biz.ErrProductNotFound
		}
		return
	}
	return
}

func (r *productRepo) ListByIds(ctx context.Context, ids []string) ([]*types.Product, error) {
	return r.db.ListByIds(ctx, ids)
}

func (r *productRepo) ListByTenant(ctx context.Context, tenantId string) (res []*types.Product, err error) {
	return r.db.ListByTenant(ctx, tenantId)
}
