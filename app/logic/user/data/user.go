package data

import (
	"context"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/app/logic/user/biz"
	"rockimserver/app/logic/user/data/database"
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

func (r *userRepo) FindByID(ctx context.Context, productId string, id string) (*types.User, error) {
	return r.db.FindByID(ctx, productId, id)
}

func (r *userRepo) FindByAccount(ctx context.Context, productId string, account string) (uid string, err error) {
	return r.db.FindByAccount(ctx, productId, account)
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
