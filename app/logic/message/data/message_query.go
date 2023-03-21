package data

import (
	"context"
	"rockimserver/apis/rockim/service/message/v1/types"
	"rockimserver/app/logic/message/biz"
	"rockimserver/app/logic/message/data/cache"
)

type messageQueryRepo struct {
	cache *cache.MessageData
}

func NewMessageQueryRepo(cache *cache.MessageData) biz.MessageQueryRepo {
	return &messageQueryRepo{cache: cache}
}

func (r *messageQueryRepo) List(ctx context.Context, productId string, msgIds []string) ([]*types.IMMessage, error) {
	return r.cache.List(ctx, productId, msgIds)
}

func (r *messageQueryRepo) ListFromCache(ctx context.Context, productId string, msgIds []string) ([]*types.IMMessage, error) {
	return r.cache.List(ctx, productId, msgIds)
}
