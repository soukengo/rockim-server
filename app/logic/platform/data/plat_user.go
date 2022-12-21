package data

import (
	"context"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/logic/platform/biz"
	"rockim/app/logic/platform/data/database"
)

type platUserRepo struct {
	db *database.PlatUserData
}

// NewPlatUserRepo .
func NewPlatUserRepo(db *database.PlatUserData) biz.PlatUserRepo {
	return &platUserRepo{
		db: db,
	}
}

func (r *platUserRepo) Create(ctx context.Context, user *types.PlatUser) error {
	return r.db.Create(ctx, user)
}
func (r *platUserRepo) Update(ctx context.Context, user *types.PlatUser) error {
	return r.db.Update(ctx, user)
}

func (r *platUserRepo) Delete(ctx context.Context, id string) error {
	return r.db.Delete(ctx, id)
}

func (r *platUserRepo) FindById(ctx context.Context, id string) (*types.PlatUser, error) {
	return r.db.FindByID(ctx, id)
}

func (r *platUserRepo) FindIdByAccount(ctx context.Context, account string) (uid string, err error) {
	return r.db.FindByAccount(ctx, account)
}

func (r *platUserRepo) Paging(ctx context.Context, req *biz.PlatUserPagingRequest) (res *biz.PlatUserPagingResponse, err error) {
	return r.db.Paging(ctx, req)
}

func (r *platUserRepo) ListRoleId(ctx context.Context, userId string) ([]string, error) {
	return r.db.ListRoleId(ctx, userId)
}
func (r *platUserRepo) SaveRoleId(ctx context.Context, userId string, roleIds []string) (err error) {
	return r.db.SaveRoleId(ctx, userId, roleIds)
}
