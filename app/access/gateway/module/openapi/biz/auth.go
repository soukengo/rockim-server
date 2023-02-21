package biz

import (
	"context"
	"rockimserver/app/access/gateway/module/openapi/biz/options"
	"rockimserver/app/access/gateway/module/openapi/biz/types"
)

type AuthRepo interface {
	CreateAuthCode(ctx context.Context, opts *options.AuthCodeCreateOptions) (*types.AuthCode, error)
}

type AuthUseCase struct {
	repo AuthRepo
}

func NewAuthUseCase(repo AuthRepo) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

func (uc *AuthUseCase) CreateAuthCode(ctx context.Context, opts *options.AuthCodeCreateOptions) (*types.AuthCode, error) {
	return uc.repo.CreateAuthCode(ctx, opts)
}
