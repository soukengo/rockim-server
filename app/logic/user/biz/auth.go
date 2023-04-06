package biz

import (
	"context"
	"github.com/google/uuid"
	"github.com/soukengo/gopkg/errors"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/logic/user/biz/options"
	"rockimserver/app/logic/user/biz/types"
	"strconv"
	"strings"
	"time"
)

var (
	ErrAuthCodeInvalid    = errors.BadRequest(reasons.User_AUTH_CODE_INVALID.String(), "授权码无效或已过期")
	ErrAccessTokenInvalid = errors.BadRequest(reasons.User_ACCESS_TOKEN_INVALID.String(), "访问令牌无效或已过期")
)

type AuthCodeRepo interface {
	SaveAuthCode(ctx context.Context, productId string, uid string, code string) (err error)
	FindByAuthCode(ctx context.Context, productId string, code string) (string, error)
	DeleteAuthCode(ctx context.Context, productId string, code string) error
}

type AccessTokenRepo interface {
	SaveAccessToken(ctx context.Context, productId string, uid string, token string) (err error)
	FindByAccessToken(ctx context.Context, productId string, token string) (string, error)
}

type AuthUseCase struct {
	repo      AuthCodeRepo
	tokenRepo AccessTokenRepo
	userRepo  UserRepo
}

const (
	authCodeExpires  = time.Minute * 10
	authTokenExpires = time.Hour * 24
	authSeparator    = "."
)

func NewAuthUseCase(repo AuthCodeRepo, tokenRepo AccessTokenRepo, userRepo UserRepo) *AuthUseCase {
	return &AuthUseCase{repo: repo, tokenRepo: tokenRepo, userRepo: userRepo}
}

func (uc *AuthUseCase) CreateAuthCode(ctx context.Context, in *options.AuthCodeCreateOptions) (out *types.AuthCode, err error) {
	uid, err := uc.userRepo.FindUidByAccount(ctx, in.ProductId, in.Account)
	if err != nil {
		return
	}
	expireTime := time.Now().Add(authCodeExpires).UnixMilli()
	authCode := uuid.New().String() + authSeparator + strconv.FormatInt(expireTime, 10)
	err = uc.repo.SaveAuthCode(ctx, in.ProductId, uid, authCode)
	if err != nil {
		return
	}
	out = &types.AuthCode{Code: authCode, ExpireTime: expireTime}
	return
}

func (uc *AuthUseCase) Login(ctx context.Context, in *options.LoginOptions) (out *types.AccessToken, err error) {
	pars := strings.Split(in.AuthCode, authSeparator)
	// eg: uuid.timestamp
	if len(pars) != 2 {
		err = ErrAuthCodeInvalid
		return
	}
	codeExpireTime, err := strconv.ParseInt(pars[1], 10, 64)
	nowts := time.Now().UnixMilli()
	if err != nil || codeExpireTime < nowts {
		err = ErrAuthCodeInvalid
		return
	}
	uid, err := uc.repo.FindByAuthCode(ctx, in.ProductId, in.AuthCode)
	if err != nil {
		return
	}
	// 生成Token
	expireTime := time.Now().Add(authTokenExpires).UnixMilli()
	token := uuid.New().String() + authSeparator + strconv.FormatInt(expireTime, 10)
	err = uc.tokenRepo.SaveAccessToken(ctx, in.ProductId, uid, token)
	if err != nil {
		return
	}
	// 清除code
	_ = uc.repo.DeleteAuthCode(ctx, in.ProductId, in.AuthCode)
	out = &types.AccessToken{Token: token, ExpireTime: expireTime}
	return
}

func (uc *AuthUseCase) CheckToken(ctx context.Context, in *options.TokenCheckOptions) (uid string, err error) {
	return uc.tokenRepo.FindByAccessToken(ctx, in.ProductId, in.Token)
}
