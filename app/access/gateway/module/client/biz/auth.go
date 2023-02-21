package biz

import (
	"context"
	"rockimserver/app/access/gateway/module/client/biz/options"
	"rockimserver/app/access/gateway/module/client/biz/types"
)

const (
	ContextValueUID = "RockIM-Context-Uid"
)

type AuthRepo interface {
	Login(ctx context.Context, opts *options.LoginOptions) (*types.AccessToken, error)
	CheckToken(ctx context.Context, opts *options.TokenCheckOptions) (string, error)
}

type AuthUseCase struct {
	repo AuthRepo
}

func NewAuthUseCase(repo AuthRepo) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

func (uc *AuthUseCase) Login(ctx context.Context, opts *options.LoginOptions) (*types.AccessToken, error) {
	return uc.repo.Login(ctx, opts)
}

func (uc *AuthUseCase) CheckToken(ctx context.Context, opts *options.TokenCheckOptions) (string, error) {
	return uc.repo.CheckToken(ctx, opts)
}

func CurrentUidFromContext(ctx context.Context) (string, bool) {
	val := ctx.Value(ContextValueUID)
	if val == nil {
		return "", false
	}
	uid, ok := val.(string)
	if ok {
		ok = len(uid) > 0
	}
	return uid, ok
}
