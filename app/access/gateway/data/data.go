package data

import (
	"context"
	"github.com/google/wire"
	userV1 "rockim/api/logic/user/v1"
	"rockim/api/logic/user/v1/types"
	"rockim/app/access/gateway/biz"
	"rockim/app/access/gateway/data/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(grpc.ProviderSet, NewUserRepo)

type userRepo struct {
	uac userV1.UserAPIClient
}

func NewUserRepo(uac userV1.UserAPIClient) biz.UserRepo {
	return &userRepo{uac: uac}
}

func (r *userRepo) Register(ctx context.Context, user *types.User) (*types.User, error) {
	ret, err := r.uac.Register(ctx, &userV1.UserRegisterRequest{
		AppId:   user.AppId,
		Bucket:  user.Bucket,
		Account: user.Account,
		Name:    user.Name,
		Fields:  user.Fields,
	})
	if err != nil {
		return nil, err
	}
	return ret.User, nil
}
