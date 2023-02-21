package data

import (
	"context"
	v1 "rockimserver/apis/rockim/service/user/v1"
	"rockimserver/app/access/gateway/module/client/biz"
	"rockimserver/app/access/gateway/module/client/biz/options"
	"rockimserver/app/access/gateway/module/client/biz/types"
)

type authRepo struct {
	ac v1.AuthAPIClient
}

func NewAuthRepo(ac v1.AuthAPIClient) biz.AuthRepo {
	return &authRepo{ac: ac}
}

func (r *authRepo) CreateAuthCode(ctx context.Context, in *options.AuthCodeCreateOptions) (out *types.AuthCode, err error) {
	ret, err := r.ac.CreateAuthCode(ctx, &v1.AuthCodeRequest{ProductId: in.ProductId, Account: in.Account})
	if err != nil {
		return
	}
	return &types.AuthCode{Code: ret.AuthCode, ExpireTime: ret.ExpireTime}, nil
}

func (r *authRepo) Login(ctx context.Context, in *options.LoginOptions) (out *types.AccessToken, err error) {
	ret, err := r.ac.Login(ctx, &v1.LoginRequest{ProductId: in.ProductId, AuthCode: in.AuthCode})
	if err != nil {
		return
	}
	return &types.AccessToken{Token: ret.Token, ExpireTime: ret.ExpireTime}, nil
}
func (r *authRepo) CheckToken(ctx context.Context, in *options.TokenCheckOptions) (out string, err error) {
	ret, err := r.ac.CheckToken(ctx, &v1.TokenCheckRequest{ProductId: in.ProductId, Token: in.Token})
	if err != nil {
		return
	}
	return ret.Uid, nil
}
