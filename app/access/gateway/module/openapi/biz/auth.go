package biz

import (
	"context"
	"github.com/soukengo/gopkg/errors"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/access/gateway/module/openapi/biz/options"
	"rockimserver/app/access/gateway/module/openapi/biz/types"
)

var (
	ErrSignInvalid = errors.BadRequest(reasons.OpenAPI_SIGN_INVALID.String(), "签名错误")
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
