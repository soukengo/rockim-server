package cache

import (
	"context"
	"github.com/soukengo/gopkg/component/cache"
	biztypes "rockimserver/app/logic/session/biz/types"
)

type ChannelData struct {
	cache cache.HashCache[biztypes.Channel]
}

func NewChannelData(mgr *cache.Manager, cfg *cache.Config) *ChannelData {
	return &ChannelData{
		cache: cache.NewHashCache[biztypes.Channel](mgr, cfg.Category(keyChannel)),
	}
}

func (d *ChannelData) Add(ctx context.Context, ch *biztypes.Channel) error {
	// 暂时只让一个客户端在线
	_ = d.cache.Delete(ctx, cache.Parts(ch.ProductId, ch.Uid))
	return d.cache.Put(ctx, cache.Parts(ch.ProductId, ch.Uid), ch.ChannelId, ch)
}

func (d *ChannelData) Exists(ctx context.Context, ch *biztypes.Channel) (bool, error) {
	return d.cache.ExistsField(ctx, cache.Parts(ch.ProductId, ch.Uid), ch.ChannelId)
}

func (d *ChannelData) Delete(ctx context.Context, ch *biztypes.Channel) error {
	return d.cache.DeleteField(ctx, cache.Parts(ch.ProductId, ch.Uid), ch.ChannelId)
}
func (d *ChannelData) ExistsUid(ctx context.Context, productId string, uid string) (bool, error) {
	return d.cache.Exists(ctx, cache.Parts(productId, uid))
}
func (d *ChannelData) ListByUid(ctx context.Context, productId string, uid string) (list []*biztypes.Channel, err error) {
	hash, err := d.cache.GetAll(ctx, cache.Parts(productId, uid))
	if err != nil {
		return
	}
	for _, item := range hash {
		list = append(list, item)
	}
	return
}
