package biz

import (
	"context"
	"rockimserver/apis/rockim/service/user/v1/types"
)

type UserRepo interface {
	FindById(ctx context.Context, productId string, uid string) (*types.User, error)
	FindByAccount(ctx context.Context, productId string, account string) (*types.User, error)
}
