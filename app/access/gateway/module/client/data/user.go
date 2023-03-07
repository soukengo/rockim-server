package data

import (
	"context"
	"rockimserver/apis/rockim/service"
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

func (r *userRepo) FindByAccount(ctx context.Context, productId string, account string) (ret *types.User, err error) {
	sret, err := r.uac.FindByAccount(ctx, &v1.UserFindByAccountRequest{Base: service.GenRequest(productId), Account: account})
	if err != nil {
		return
	}
	return sret.User, nil
}
