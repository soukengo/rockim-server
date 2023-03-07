package data

import (
	"context"
	"rockimserver/apis/rockim/service"
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
		Base:      service.GenRequest(opts.ProductId),
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

func (r *userRepo) FindUid(ctx context.Context, productId string, account string) (out string, err error) {
	ret, err := r.uac.FindUid(ctx, &v1.UserIdFindRequest{Base: service.GenRequest(productId), Account: account})
	if err != nil {
		return
	}
	return ret.Uid, nil
}
