package data

import (
	"context"
	"github.com/soukengo/gopkg/errors"
	"rockimserver/app/logic/session/biz"
	"rockimserver/app/logic/session/biz/types"
	"rockimserver/app/logic/session/data/cache"
)

type channelRepo struct {
	cache *cache.ChannelData
}

func NewChannelRepo(cache *cache.ChannelData) biz.ChannelRepo {
	return &channelRepo{cache: cache}
}

func (r *channelRepo) Add(ctx context.Context, ch *types.Channel) error {
	return r.cache.Add(ctx, ch)
}

func (r *channelRepo) Exists(ctx context.Context, ch *types.Channel) (exists bool, err error) {
	exists, err = r.cache.Exists(ctx, ch)
	if err != nil {
		if errors.IsNotFound(err) {
			err = nil
		}
		return
	}
	return
}

func (r *channelRepo) Delete(ctx context.Context, ch *types.Channel) error {
	return r.cache.Delete(ctx, ch)
}
