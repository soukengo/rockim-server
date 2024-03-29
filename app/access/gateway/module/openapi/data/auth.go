package data

import (
	"context"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/user/v1"
	"rockimserver/app/access/gateway/module/openapi/biz"
	"rockimserver/app/access/gateway/module/openapi/biz/options"
	"rockimserver/app/access/gateway/module/openapi/biz/types"
)

type authRepo struct {
	ac v1.AuthAPIClient
}

func NewAuthRepo(ac v1.AuthAPIClient) biz.AuthRepo {
	return &authRepo{ac: ac}
}

func (r *authRepo) CreateAuthCode(ctx context.Context, in *options.AuthCodeCreateOptions) (out *types.AuthCode, err error) {
	ret, err := r.ac.CreateAuthCode(ctx, &v1.AuthCodeRequest{Base: service.GenRequest(in.ProductId), Account: in.Account})
	if err != nil {
		return
	}
	return &types.AuthCode{Code: ret.AuthCode, ExpireTime: ret.ExpireTime}, nil
}
