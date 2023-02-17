package cache

import (
	"context"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
)

type UserData struct {
	redisCli *redis.Client
}

func NewUserData(redisCli *redis.Client) *UserData {
	return &UserData{redisCli: redisCli}
}

func (d *UserData) FindUserByID(ctx context.Context, productId string, uid string) (*types.User, error) {
	val, err := d.genCache(productId, uid).Get(ctx)
	return val, err
}
func (d *UserData) SaveUser(ctx context.Context, productId string, uid string, record *types.User) error {
	return d.genCache(productId, uid).Set(ctx, record)
}

func (d *UserData) FindUidByAccount(ctx context.Context, productId string, account string) (id string, err error) {
	val, err := d.genAccountCache(productId, account).Get(ctx)
	if err != nil {
		return
	}
	id = *val
	return
}

func (d *UserData) SaveAccountUid(ctx context.Context, productId string, account string, uid string) error {
	val := &uid
	if len(uid) == 0 {
		val = nil
	}
	return d.genAccountCache(productId, account).Set(ctx, val)
}
func (d *UserData) DeleteAccountUid(ctx context.Context, productId string, account string) error {
	return d.genAccountCache(productId, account).Delete(ctx)
}

// genCache generates a cache for user record
func (d *UserData) genCache(productId string, uid string) cache.ValueCache[types.User] {
	return newValueCache[types.User](d.redisCli, GenUserCacheKey(productId, uid))
}

// genAccountCache generates a cache for user account to uid
func (d *UserData) genAccountCache(productId string, account string) cache.ValueCache[string] {
	return newValueCache[string](d.redisCli, GenUserAccountCacheKey(productId, account))
}
