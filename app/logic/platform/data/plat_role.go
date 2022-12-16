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

func (r *platRoleRepo) Create(ctx context.Context, record *types.PlatRole) error {
	return r.db.Create(ctx, record)
}
func (r *platRoleRepo) Update(ctx context.Context, record *types.PlatRole) error {
	return r.db.Update(ctx, record)
}
func (r *platRoleRepo) Delete(ctx context.Context, id string) error {
	return r.db.Delete(ctx, id)
}

func (r *platRoleRepo) FindById(ctx context.Context, id string) (*types.PlatRole, error) {
	return r.db.FindByID(ctx, id)
}

func (r *platRoleRepo) ListByIds(ctx context.Context, ids []string) ([]*types.PlatRole, error) {
	return r.db.ListByIds(ctx, ids)
}
func (r *platRoleRepo) Paging(ctx context.Context, req *biz.PlatRolePagingRequest) (res *biz.PlatRolePagingResponse, err error) {
	return r.db.Paging(ctx, req)
}

func (r *platRoleRepo) ListResourceId(ctx context.Context, ids []string) ([]string, error) {
	return r.db.ListResourceId(ctx, ids)
}
