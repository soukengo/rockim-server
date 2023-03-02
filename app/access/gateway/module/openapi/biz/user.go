package biz

import (
	"context"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/app/access/gateway/module/openapi/biz/options"
)

type UserRepo interface {
	Register(ctx context.Context, opts *options.UserRegisterOptions) (*types.User, error)
	FindUid(ctx context.Context, productId string, account string) (string, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Register(ctx context.Context, opts *options.UserRegisterOptions) (*types.User, error) {
	return uc.repo.Register(ctx, opts)
}
func (uc *UserUseCase) FindUid(ctx context.Context, productId string, account string) (string, error) {
	return uc.repo.FindUid(ctx, productId, account)
}
