package biz

import (
	"context"
	"github.com/soukengo/gopkg/errors"
	usertypes "rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/access/gateway/module/client/biz/options"
	"rockimserver/app/access/gateway/module/client/biz/types"
)

const (
	ContextValueUID = "RockIM-Context-Uid"
)

var (
	ErrSignInvalid        = errors.BadRequest(reasons.Client_SIGN_INVALID.String(), "签名错误")
	ErrAccessTokenInvalid = errors.Unauthorized(reasons.User_ACCESS_TOKEN_INVALID.String(), "访问令牌无效或已过期")
)

type AuthRepo interface {
	Login(ctx context.Context, opts *options.LoginOptions) (*types.AccessToken, *usertypes.User, error)
	CheckToken(ctx context.Context, opts *options.TokenCheckOptions) (string, error)
}

type AuthUseCase struct {
	repo AuthRepo
}

func NewAuthUseCase(repo AuthRepo) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

func (uc *AuthUseCase) Login(ctx context.Context, opts *options.LoginOptions) (*types.AccessToken, *usertypes.User, error) {
	return uc.repo.Login(ctx, opts)
}

func (uc *AuthUseCase) CheckToken(ctx context.Context, opts *options.TokenCheckOptions) (string, error) {
	return uc.repo.CheckToken(ctx, opts)
}

func CurrentUidFromContext(ctx context.Context) (string, error) {
	val := ctx.Value(ContextValueUID)
	if val == nil {
		return "", ErrAccessTokenInvalid
	}
	uid, ok := val.(string)
	if ok {
		ok = len(uid) > 0
	}
	return uid, nil
}
