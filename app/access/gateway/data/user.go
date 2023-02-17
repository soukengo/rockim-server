package data

import (
	"context"
	v1 "rockimserver/apis/rockim/service/user/v1"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/app/access/gateway/biz"
	"rockimserver/pkg/log"
)

type userRepo struct {
	uac v1.UserAPIClient
	aac v1.AuthAPIClient
}

func NewUserRepo(uac v1.UserAPIClient, aac v1.AuthAPIClient) biz.UserRepo {
	return &userRepo{uac: uac, aac: aac}
}

func (r *userRepo) Register(ctx context.Context, user *types.User) (*types.User, error) {
	ret, err := r.uac.Register(ctx, &v1.UserRegisterRequest{
		ProductId: user.ProductId,
		Bucket:    user.Bucket,
		Account:   user.Account,
		Name:      user.Name,
		Fields:    user.Fields,
	})
	if err != nil {
		return nil, err
	}
	return ret.User, nil
}

func (r *userRepo) Find(ctx context.Context, productId string, account string) (ret *types.User, err error) {
	lr, err := r.aac.Login(ctx, &v1.LoginRequest{})
	if err != nil {
		return
	}
	log.WithContext(ctx).Infof("lr: %v", lr)
	sret, err := r.uac.Find(ctx, &v1.UserFindRequest{ProductId: productId, Account: account})
	if err != nil {
		return
	}
	return sret.User, nil
}
