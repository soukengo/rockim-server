package data

import (
	"context"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/user/v1"
	"rockimserver/app/access/comet/module/client/biz"
	"rockimserver/app/access/comet/module/client/biz/options"
)

type tokenRepo struct {
	ac v1.AuthAPIClient
}

func NewTokenRepo(ac v1.AuthAPIClient) biz.TokenRepo {
	return &tokenRepo{ac: ac}
}

func (r *tokenRepo) CheckAuth(ctx context.Context, opts *options.ChannelAuthOptions) (uid string, token error) {
	ret, err := r.ac.CheckToken(ctx, &v1.TokenCheckRequest{
		Base:  service.GenRequest(opts.ProductId),
		Token: opts.Token,
	})
	if err != nil {
		return
	}
	uid = ret.Uid
	return
}
