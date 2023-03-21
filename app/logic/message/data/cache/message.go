package cache

import (
	"context"
	"rockimserver/apis/rockim/service/message/v1/types"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
)

type MessageData struct {
	cache cache.ValueCache[types.IMMessage]
}

func NewMessageData(rds *redis.Client, cfg *cache.Config) *MessageData {
	return &MessageData{
		cache: newValueCache[types.IMMessage](rds, cfg.Category(keyMessage)),
	}
}

func (d *MessageData) Save(ctx context.Context, message *types.IMMessage) (err error) {
	return d.cache.Set(ctx, cache.Parts(message.ProductId, message.MsgId), message)
}
func (d *MessageData) List(ctx context.Context, productId string, msgIds []string) (list []*types.IMMessage, err error) {
	return d.cache.List(ctx, cache.Parts(productId), msgIds)
}
