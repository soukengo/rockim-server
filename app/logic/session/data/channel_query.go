package data

import (
	"context"
	"rockimserver/app/logic/session/biz"
	biztypes "rockimserver/app/logic/session/biz/types"
	"rockimserver/app/logic/session/data/cache"
)

type channelQueryRepo struct {
	cache *cache.ChannelData
}

func NewChannelQueryRepo(cache *cache.ChannelData) biz.ChannelQueryRepo {
	return &channelQueryRepo{cache: cache}
}

func (r *channelQueryRepo) Exists(ctx context.Context, productId string, uid string) (bool, error) {
	return r.cache.ExistsUid(ctx, productId, uid)
}

func (r *channelQueryRepo) ListByUid(ctx context.Context, productId string, uid string) (list []*biztypes.Channel, err error) {
	return r.cache.ListByUid(ctx, productId, uid)
}
