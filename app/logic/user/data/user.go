package data

import (
	"context"

	"rockim/app/logic/user/biz"
)

type userRepo struct {
	data *Data
}

// NewUserRepo .
func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (u userRepo) FindByID(ctx context.Context, appId string, id string) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) FindByAccount(ctx context.Context, appId string, account string) (uid string, err error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) GenID() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}
