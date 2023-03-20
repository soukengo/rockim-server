package data

import (
	"context"
	"rockimserver/app/logic/user/biz"
	biztypes "rockimserver/app/logic/user/biz/types"
	"rockimserver/app/logic/user/data/cache"
)

type onlineQueryRepo struct {
	cache *cache.OnlineData
}

func NewOnlineQueryRepo(cache *cache.OnlineData) biz.OnlineQueryRepo {
	return &onlineQueryRepo{cache: cache}
}

func (r *onlineQueryRepo) Exists(ctx context.Context, productId string, uid string) (bool, error) {
	return r.cache.ExistsUid(ctx, productId, uid)
}

func (r *onlineQueryRepo) ListByUid(ctx context.Context, productId string, uid string) (list []*biztypes.OnlineChannel, err error) {
	return r.cache.ListByUid(ctx, productId, uid)
}
