package data

import (
	"context"
	"github.com/soukengo/gopkg/errors"
	"rockimserver/app/logic/user/biz"
	"rockimserver/app/logic/user/biz/types"
	"rockimserver/app/logic/user/data/cache"
)

type onlineRepo struct {
	cache *cache.OnlineData
}

func NewOnlineRepo(cache *cache.OnlineData) biz.OnlineRepo {
	return &onlineRepo{cache: cache}
}

func (r *onlineRepo) Add(ctx context.Context, ch *types.OnlineChannel) error {
	return r.cache.Add(ctx, ch)
}

func (r *onlineRepo) Exists(ctx context.Context, ch *types.OnlineChannel) (exists bool, err error) {
	exists, err = r.cache.Exists(ctx, ch)
	if err != nil {
		if errors.IsNotFound(err) {
			err = nil
		}
		return
	}
	return
}

func (r *onlineRepo) Delete(ctx context.Context, ch *types.OnlineChannel) error {
	return r.cache.Delete(ctx, ch)
}
