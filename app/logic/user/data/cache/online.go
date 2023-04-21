package cache

import (
	"context"
	"github.com/soukengo/gopkg/component/cache"
	biztypes "rockimserver/app/logic/user/biz/types"
)

type OnlineData struct {
	cache cache.HashCache[biztypes.OnlineChannel]
}

func NewOnlineData(mgr *cache.Manager, cfg *cache.Config) *OnlineData {
	return &OnlineData{
		cache: cache.NewHashCache[biztypes.OnlineChannel](mgr, cfg.Category(keyOnline)),
	}
}

func (d *OnlineData) Add(ctx context.Context, ch *biztypes.OnlineChannel) error {
	// 暂时只让一个客户端在线
	_ = d.cache.Delete(ctx, cache.Parts(ch.ProductId, ch.Uid))
	return d.cache.Put(ctx, cache.Parts(ch.ProductId, ch.Uid), ch.ChannelId, ch)
}

func (d *OnlineData) Exists(ctx context.Context, ch *biztypes.OnlineChannel) (bool, error) {
	return d.cache.ExistsField(ctx, cache.Parts(ch.ProductId, ch.Uid), ch.ChannelId)
}

func (d *OnlineData) Delete(ctx context.Context, ch *biztypes.OnlineChannel) error {
	return d.cache.DeleteField(ctx, cache.Parts(ch.ProductId, ch.Uid), ch.ChannelId)
}
func (d *OnlineData) ExistsUid(ctx context.Context, productId string, uid string) (bool, error) {
	return d.cache.Exists(ctx, cache.Parts(productId, uid))
}
func (d *OnlineData) ListByUid(ctx context.Context, productId string, uid string) (list []*biztypes.OnlineChannel, err error) {
	hash, err := d.cache.GetAll(ctx, cache.Parts(productId, uid))
	if err != nil {
		return
	}
	for _, item := range hash {
		list = append(list, item)
	}
	return
}
