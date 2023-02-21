package cache

import (
	"context"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
)

type UserData struct {
	cache        cache.ValueCache[types.User]
	accountCache cache.ValueCache[string]
}

func NewUserData(redisCli *redis.Client, cfg *cache.Config) *UserData {
	return &UserData{
		cache:        newValueCache[types.User](redisCli, cfg.Category(keyUser)),
		accountCache: newValueCache[string](redisCli, cfg.Category(keyUserAccount)),
	}
}

func (d *UserData) FindUserByID(ctx context.Context, productId string, uid string) (*types.User, error) {
	val, err := d.cache.Get(ctx, genKey(productId, uid))
	return val, err
}
func (d *UserData) SaveUser(ctx context.Context, productId string, uid string, record *types.User) error {
	return d.cache.Set(ctx, genKey(productId, uid), record)
}

func (d *UserData) FindUidByAccount(ctx context.Context, productId string, account string) (id string, err error) {
	val, err := d.accountCache.Get(ctx, genKey(productId, account))
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
	return d.accountCache.Set(ctx, genKey(productId, account), val)
}
func (d *UserData) DeleteAccountUid(ctx context.Context, productId string, account string) error {
	return d.accountCache.Delete(ctx, genKey(productId, account))
}
