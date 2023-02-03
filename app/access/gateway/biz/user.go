package biz

import (
	"context"
	"rockimserver/apis/rockim/service/user/v1/types"
)

type UserRepo interface {
	Register(ctx context.Context, user *types.User) (*types.User, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) RegisterUser(ctx context.Context, user *types.User) (*types.User, error) {
	return uc.repo.Register(ctx, user)
}
