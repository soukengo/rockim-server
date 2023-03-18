package cache

import (
	"context"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/errors"
	"rockimserver/pkg/log"
)

type AuthCodeData struct {
	logger       *log.Helper
	cache        cache.ValueCache[string]
	reverseCache cache.ValueCache[string]
}
type AccessTokenData struct {
	logger       *log.Helper
	cache        cache.ValueCache[string]
	reverseCache cache.ValueCache[string]
}

func NewAuthCodeData(logger log.Logger, redisCli *redis.Client, cfg *cache.Config) *AuthCodeData {
	return &AuthCodeData{
		logger:       logger.Helper(),
		cache:        newValueCache[string](redisCli, cfg.Category(keyAuthCode)),
		reverseCache: newValueCache[string](redisCli, cfg.Category(keyAuthCodeReverse)),
	}
}
func NewAccessTokenData(logger log.Logger, redisCli *redis.Client, cfg *cache.Config) *AccessTokenData {
	return &AccessTokenData{
		logger:       logger.Helper(),
		cache:        newValueCache[string](redisCli, cfg.Category(keyAuthToken)),
		reverseCache: newValueCache[string](redisCli, cfg.Category(keyAuthTokenReverse)),
	}
}

func (d *AuthCodeData) SaveAuthCode(ctx context.Context, productId string, uid string, code string) (err error) {
	keySuffix := cache.Parts(productId, uid)
	oldCode, err := d.cache.Get(ctx, keySuffix)
	if err != nil && !errors.IsNotFound(err) {
		return
	}
	// 删除旧的code
	if oldCode != nil {
		err1 := d.reverseCache.Delete(ctx, cache.Parts(productId, *oldCode))
		if err1 != nil {
			d.logger.WithContext(ctx).Errorf("d.reverseCache.Delete productId: %s, oldCode: %s, err: %v", productId, *oldCode, err1)
		}
	}
	val := &code
	if len(code) == 0 {
		val = nil
	}
	err = d.cache.Set(ctx, keySuffix, val)
	if err != nil {
		return
	}
	return d.reverseCache.Set(ctx, cache.Parts(productId, code), &uid)
}

func (d *AuthCodeData) FindUidByAuthCode(ctx context.Context, productId string, code string) (id string, err error) {
	val, err := d.reverseCache.Get(ctx, cache.Parts(productId, code))
	if err != nil {
		return
	}
	id = *val
	return
}
func (d *AuthCodeData) DeleteAuthCode(ctx context.Context, productId string, code string) (err error) {
	return d.reverseCache.Delete(ctx, cache.Parts(productId, code))
}

func (d *AccessTokenData) SaveAccessToken(ctx context.Context, productId string, uid string, token string) (err error) {
	keySuffix := cache.Parts(productId, uid)
	oldToken, err := d.cache.Get(ctx, keySuffix)
	if err != nil && !errors.IsNotFound(err) {
		return
	}
	// 删除旧的token
	if oldToken != nil {
		err1 := d.reverseCache.Delete(ctx, cache.Parts(productId, *oldToken))
		if err1 != nil {
			d.logger.WithContext(ctx).Errorf("d.reverseCache.Delete productId: %s, oldToken: %s, err: %v", productId, *oldToken, err1)
		}
	}
	val := &token
	if len(token) == 0 {
		val = nil
	}
	err = d.cache.Set(ctx, keySuffix, val)
	if err != nil {
		return
	}
	return d.reverseCache.Set(ctx, cache.Parts(productId, token), &uid)
}

func (d *AccessTokenData) FindUidByAccessToken(ctx context.Context, productId string, token string) (id string, err error) {
	val, err := d.reverseCache.Get(ctx, cache.Parts(productId, token))
	if err != nil {
		return
	}
	id = *val
	return
}
