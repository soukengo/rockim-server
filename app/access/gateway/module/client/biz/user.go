package biz

import (
	"context"
	"rockimserver/apis/rockim/service/user/v1/types"
)

type UserRepo interface {
	Find(ctx context.Context, productId string, account string) (*types.User, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Find(ctx context.Context, productId string, account string) (*types.User, error) {
	return uc.repo.Find(ctx, productId, account)
}
