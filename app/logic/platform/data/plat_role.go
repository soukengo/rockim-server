package data

import (
	"context"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/logic/platform/biz"
	"rockim/app/logic/platform/data/database"
)

type platRoleRepo struct {
	db *database.PlatRoleData
}

func NewPlatRoleRepo(db *database.PlatRoleData) biz.PlatRoleRepo {
	return &platRoleRepo{db: db}
}

func (r *platRoleRepo) ListByIds(ctx context.Context, ids []string) ([]*types.PlatRole, error) {
	return r.db.ListByIds(ctx, ids)
}
func (r *platRoleRepo) ListResourceId(ctx context.Context, ids []string) ([]string, error) {
	return r.db.ListResourceId(ctx, ids)
}
