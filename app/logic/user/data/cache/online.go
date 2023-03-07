package cache

import (
	"context"
	"rockimserver/app/logic/user/biz/types"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
)

type OnlineData struct {
	cache cache.HashCache[types.OnlineChannel]
}

func NewOnlineData(rds *redis.Client, cfg *cache.Config) *OnlineData {
	return &OnlineData{
		cache: newHashCache[types.OnlineChannel](rds, cfg.Category(keyOnline)),
	}
}

func (d *OnlineData) Add(ctx context.Context, ch *types.OnlineChannel) error {
	return d.cache.Put(ctx, cache.Parts(ch.ProductId, ch.Uid), ch.ChannelId, ch)
}

func (d *OnlineData) Exists(ctx context.Context, ch *types.OnlineChannel) (bool, error) {
	return d.cache.ExistsField(ctx, cache.Parts(ch.ProductId, ch.Uid), ch.ChannelId)
}

func (d *OnlineData) Delete(ctx context.Context, ch *types.OnlineChannel) error {
	return d.cache.DeleteField(ctx, cache.Parts(ch.ProductId, ch.Uid), ch.ChannelId)
}
