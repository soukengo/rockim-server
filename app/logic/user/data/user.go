package data

import (
	"context"
	"rockim/api/logic/user/v1/types"
	"rockim/app/logic/user/data/database"

	"rockim/app/logic/user/biz"
)

type userRepo struct {
	db *database.UserData
}

// NewUserRepo .
func NewUserRepo(db *database.UserData) biz.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) FindByID(ctx context.Context, appId string, id string) (*types.User, error) {
	return r.db.FindByID(ctx, appId, id)
}

func (r *userRepo) FindByAccount(ctx context.Context, appId string, account string) (uid string, err error) {
	return r.db.FindByAccount(ctx, appId, account)
}

func (r *userRepo) GenID(ctx context.Context) (string, error) {
	return r.db.GenID(ctx)
}

func (r *userRepo) Create(ctx context.Context, user *types.User) error {
	return r.db.Create(ctx, user)
}
func (r *userRepo) Update(ctx context.Context, user *types.User) error {
	return r.db.Update(ctx, user)
}
