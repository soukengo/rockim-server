package cache

import (
	"context"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/errors"
	"rockimserver/pkg/log"
)

type AuthCodeData struct {
	cache        cache.ValueCache[string]
	reverseCache cache.ValueCache[string]
}
type AccessTokenData struct {
	cache        cache.ValueCache[string]
	reverseCache cache.ValueCache[string]
}

func NewAuthCodeData(redisCli *redis.Client, cfg *cache.Config) *AuthCodeData {
	return &AuthCodeData{
		cache:        newValueCache[string](redisCli, cfg.Category(keyAuthCode)),
		reverseCache: newValueCache[string](redisCli, cfg.Category(keyAuthCodeReverse)),
	}
}
func NewAccessTokenData(redisCli *redis.Client, cfg *cache.Config) *AccessTokenData {
	return &AccessTokenData{
		cache:        newValueCache[string](redisCli, cfg.Category(keyAuthToken)),
		reverseCache: newValueCache[string](redisCli, cfg.Category(keyAuthTokenReverse)),
	}
}

func (d *AuthCodeData) SaveAuthCode(ctx context.Context, productId string, uid string, code string) (err error) {
	keySuffix := genKey(productId, uid)
	oldCode, err := d.cache.Get(ctx, keySuffix)
	if !errors.IsNotFound(err) {
		return
	}
	// 删除旧的code
	if oldCode != nil {
		err1 := d.reverseCache.Delete(ctx, genKey(productId, *oldCode))
		if err1 != nil {
			log.WithContext(ctx).Errorf("d.reverseCache.Delete productId: %s, oldCode: %s, err: %v", productId, *oldCode, err1)
		}
	}
	val := &code
	if len(code) == 0 {
		val = nil
	}
	return d.cache.Set(ctx, keySuffix, val)
}

func (d *AuthCodeData) FindUidByAuthCode(ctx context.Context, productId string, code string) (id string, err error) {
	val, err := d.reverseCache.Get(ctx, genKey(productId, code))
	if err != nil {
		return
	}
	id = *val
	return
}

func (d *AccessTokenData) SaveAccessToken(ctx context.Context, productId string, uid string, token string) (err error) {
	keySuffix := genKey(productId, uid)
	oldToken, err := d.cache.Get(ctx, keySuffix)
	if !errors.IsNotFound(err) {
		return
	}
	// 删除旧的token
	if oldToken != nil {
		err1 := d.reverseCache.Delete(ctx, genKey(productId, *oldToken))
		if err1 != nil {
			log.WithContext(ctx).Errorf("d.reverseCache.Delete productId: %s, oldToken: %s, err: %v", productId, *oldToken, err1)
		}
	}
	val := &token
	if len(token) == 0 {
		val = nil
	}
	return d.cache.Set(ctx, keySuffix, val)
}

func (d *AccessTokenData) FindUidByAccessToken(ctx context.Context, productId string, token string) (id string, err error) {
	val, err := d.reverseCache.Get(ctx, genKey(productId, token))
	if err != nil {
		return
	}
	id = *val
	return
}
