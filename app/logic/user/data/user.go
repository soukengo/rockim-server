package data

import (
	"context"
	"rockim/app/logic/user/data/database"

	"rockim/app/logic/user/biz"
)

type userRepo struct {
	data *database.UserData
}

// NewUserRepo .
func NewUserRepo(data *database.UserData) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (u *userRepo) FindByID(ctx context.Context, appId string, id string) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) FindByAccount(ctx context.Context, appId string, account string) (uid string, err error) {
	return u.data.FindByAccount(ctx, appId, account)
}

func (u *userRepo) GenID(ctx context.Context) (string, error) {
	return u.data.GenID(ctx)
}

func (u *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}
