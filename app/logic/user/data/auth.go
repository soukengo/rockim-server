package data

import (
	"context"
	"github.com/soukengo/gopkg/errors"
	"rockimserver/app/logic/user/biz"
	"rockimserver/app/logic/user/data/cache"
)

type authCodeRepo struct {
	cache *cache.AuthCodeData
}

func NewAuthCodeRepo(cache *cache.AuthCodeData) biz.AuthCodeRepo {
	return &authCodeRepo{cache: cache}
}

func (r *authCodeRepo) SaveAuthCode(ctx context.Context, productId string, uid string, code string) (err error) {
	return r.cache.SaveAuthCode(ctx, productId, uid, code)
}

func (r *authCodeRepo) FindByAuthCode(ctx context.Context, productId string, code string) (uid string, err error) {
	uid, err = r.cache.FindUidByAuthCode(ctx, productId, code)
	if err != nil {
		if errors.IsNotFound(err) {
			err = biz.ErrAuthCodeInvalid
		}
		return
	}
	return
}
func (r *authCodeRepo) DeleteAuthCode(ctx context.Context, productId string, code string) (err error) {
	return r.cache.DeleteAuthCode(ctx, productId, code)
}

type accessTokenRepo struct {
	cache *cache.AccessTokenData
}

func NewAccessTokenRepo(cache *cache.AccessTokenData) biz.AccessTokenRepo {
	return &accessTokenRepo{cache: cache}
}

func (r *accessTokenRepo) SaveAccessToken(ctx context.Context, productId string, uid string, token string) (err error) {
	return r.cache.SaveAccessToken(ctx, productId, uid, token)
}

func (r *accessTokenRepo) FindByAccessToken(ctx context.Context, productId string, token string) (uid string, err error) {
	uid, err = r.cache.FindUidByAccessToken(ctx, productId, token)
	if err != nil {
		if errors.IsNotFound(err) {
			err = biz.ErrAuthCodeInvalid
		}
		return
	}
	return
}
