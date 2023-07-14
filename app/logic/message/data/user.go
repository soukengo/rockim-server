package data

import (
	"context"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/user/v1"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/app/logic/message/biz"
)

type userRepo struct {
	ac v1.UserAPIClient
}

func NewUserRepo(ac v1.UserAPIClient) biz.UserRepo {
	return &userRepo{ac: ac}
}

func (r *userRepo) FindById(ctx context.Context, productId string, uid string) (out *types.User, err error) {
	ret, err := r.ac.Find(ctx, &v1.UserFindRequest{
		Base: service.GenRequest(productId),
		Uid:  uid,
	})
	if err != nil {
		return
	}
	out = ret.User
	return
}
func (r *userRepo) FindByAccount(ctx context.Context, productId string, account string) (out *types.User, err error) {
	ret, err := r.ac.FindByAccount(ctx, &v1.UserFindByAccountRequest{
		Base:    service.GenRequest(productId),
		Account: account,
	})
	if err != nil {
		return
	}
	out = ret.User
	return
}
