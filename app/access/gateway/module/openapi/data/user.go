package data

import (
	"context"
	v1 "rockimserver/apis/rockim/service/user/v1"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/app/access/gateway/module/openapi/biz"
	"rockimserver/app/access/gateway/module/openapi/biz/options"
)

type userRepo struct {
	uac v1.UserAPIClient
	aac v1.AuthAPIClient
}

func NewUserRepo(uac v1.UserAPIClient, aac v1.AuthAPIClient) biz.UserRepo {
	return &userRepo{uac: uac, aac: aac}
}

func (r *userRepo) Register(ctx context.Context, opts *options.UserRegisterOptions) (*types.User, error) {
	ret, err := r.uac.Register(ctx, &v1.UserRegisterRequest{
		ProductId: opts.ProductId,
		Account:   opts.Account,
		Name:      opts.Name,
		AvatarUrl: opts.AvatarUrl,
		Fields:    opts.Fields,
	})
	if err != nil {
		return nil, err
	}
	return ret.User, nil
}
