package cache

import (
	"context"
	"github.com/soukengo/gopkg/component/cache"
	"rockimserver/apis/rockim/service/group/v1/types"
)

type GroupData struct {
	cache        cache.ValueCache[types.Group]
	accountCache cache.ValueCache[string]
}

func NewGroupData(mgr *cache.Manager, cfg *cache.Config) *GroupData {
	return &GroupData{
		cache:        cache.NewValueCache[types.Group](mgr, cfg.Category(keyGroup)),
		accountCache: cache.NewValueCache[string](mgr, cfg.Category(keyGroupBizId)),
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

func (d *GroupData) FindGroupId(ctx context.Context, productId string, bizId string) (id string, err error) {
	val, err := d.accountCache.Get(ctx, cache.Parts(productId, bizId))
	if err != nil {
		return
	}
	id = *val
	return
}

func (d *GroupData) SaveBizId(ctx context.Context, productId string, bizId string, uid string) error {
	val := &uid
	if len(uid) == 0 {
		val = nil
	}
	return d.accountCache.Set(ctx, cache.Parts(productId, bizId), val)
}
func (d *GroupData) DeleteBizId(ctx context.Context, productId string, bizId string) error {
	return d.accountCache.Delete(ctx, cache.Parts(productId, bizId))
}
