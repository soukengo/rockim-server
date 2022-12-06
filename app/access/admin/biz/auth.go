package biz

import (
	"context"
	"rockim/api/rockim/service/user/v1/types"
)

type UserRepo interface {
	Register(ctx context.Context, user *types.User) (*types.User, error)
}

type AuthUseCase struct {
	repo UserRepo
}

func NewAuthUseCase(repo UserRepo) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

func (uc *AuthUseCase) RegisterUser(ctx context.Context, user *types.User) (*types.User, error) {
	return uc.repo.Register(ctx, user)
}
