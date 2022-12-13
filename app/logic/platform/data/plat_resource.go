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

func (r *platResourceRepo) ListByIds(ctx context.Context, ids []string) ([]*types.PlatResource, error) {
	return r.db.ListByIds(ctx, ids)
}
