package data

import (
	"context"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/logic/platform/biz"
	"rockim/app/logic/platform/data/database"
)

type platUserRepo struct {
	db *database.UserData
}

// NewUserRepo .
func NewUserRepo(db *database.UserData) biz.PlatUserRepo {
	return &platUserRepo{
		db: db,
	}
}

func (r *platUserRepo) FindById(ctx context.Context, id string) (*types.PlatUser, error) {
	return r.db.FindByID(ctx, id)
}

func (r *platUserRepo) FindIdByAccount(ctx context.Context, account string) (uid string, err error) {
	return r.db.FindByAccount(ctx, account)
}

func (r *platUserRepo) Create(ctx context.Context, user *types.PlatUser) error {
	return r.db.Create(ctx, user)
}
func (r *platUserRepo) Update(ctx context.Context, user *types.PlatUser) error {
	return r.db.Update(ctx, user)
}
