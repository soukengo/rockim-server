package data

import (
	"github.com/google/wire"
	userV1 "rockim/api/logic/user/v1"
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

func (r *userRepo) Register() error {
	return nil
}
