package data

import (
	"context"
	v1 "rockimserver/apis/rockim/service/user/v1"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/app/access/gateway/module/client/biz"
)

type userRepo struct {
	uac v1.UserAPIClient
	aac v1.AuthAPIClient
}

func NewUserRepo(uac v1.UserAPIClient, aac v1.AuthAPIClient) biz.UserRepo {
	return &userRepo{uac: uac, aac: aac}
}

func (r *userRepo) Find(ctx context.Context, productId string, account string) (ret *types.User, err error) {
	sret, err := r.uac.Find(ctx, &v1.UserFindRequest{ProductId: productId, Account: account})
	if err != nil {
		return
	}
	return sret.User, nil
}
