package cache

import (
	"context"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/database/redis"
	"rockimserver/apis/rockim/service/group/v1/types"
)

type GroupData struct {
	cache        cache.ValueCache[types.Group]
	accountCache cache.ValueCache[string]
}

func NewGroupData(rds *redis.Client, cfg *cache.Config) *GroupData {
	return &GroupData{
		cache:        newValueCache[types.Group](rds, cfg.Category(keyGroup)),
		accountCache: newValueCache[string](rds, cfg.Category(keyGroupCustomGroupId)),
	}
}

func (d *GroupData) FindByID(ctx context.Context, productId string, uid string) (*types.Group, error) {
	val, err := d.cache.Get(ctx, cache.Parts(productId, uid))
	return val, err
}
func (d *GroupData) SaveGroup(ctx context.Context, productId string, uid string, record *types.Group) error {
	return d.cache.Set(ctx, cache.Parts(productId, uid), record)
}

func (d *GroupData) DeleteGroup(ctx context.Context, productId string, uid string) error {
	return d.cache.Delete(ctx, cache.Parts(productId, uid))
}

func (d *GroupData) FindGroupId(ctx context.Context, productId string, customGroupId string) (id string, err error) {
	val, err := d.accountCache.Get(ctx, cache.Parts(productId, customGroupId))
	if err != nil {
		return
	}
	id = *val
	return
}

func (d *GroupData) SaveCustomGroupId(ctx context.Context, productId string, customGroupId string, uid string) error {
	val := &uid
	if len(uid) == 0 {
		val = nil
	}
	return d.accountCache.Set(ctx, cache.Parts(productId, customGroupId), val)
}
func (d *GroupData) DeleteCustomGroupId(ctx context.Context, productId string, customGroupId string) error {
	return d.accountCache.Delete(ctx, cache.Parts(productId, customGroupId))
}
