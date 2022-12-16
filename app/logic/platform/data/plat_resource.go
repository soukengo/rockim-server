package data

import (
	"context"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/logic/platform/biz"
	"rockim/app/logic/platform/data/database"
)

type platResourceRepo struct {
	db *database.PlatResourceData
}

func NewPlatResourceRepo(db *database.PlatResourceData) biz.PlatResourceRepo {
	return &platResourceRepo{db: db}
}

func (r *platResourceRepo) Create(ctx context.Context, resource *types.PlatResource) error {
	return r.db.Create(ctx, resource)
}

func (r *platResourceRepo) Update(ctx context.Context, resource *types.PlatResource) error {
	return r.db.Update(ctx, resource)
}

func (r *platResourceRepo) Delete(ctx context.Context, id string) error {
	return r.db.Delete(ctx, id)
}

func (r *platResourceRepo) FindById(ctx context.Context, id string) (*types.PlatResource, error) {
	return r.db.FindByID(ctx, id)
}

func (r *platResourceRepo) List(ctx context.Context) ([]*types.PlatResource, error) {
	return r.db.List(ctx)
}
func (r *platResourceRepo) ListByIds(ctx context.Context, ids []string) ([]*types.PlatResource, error) {
	return r.db.ListByIds(ctx, ids)
}
