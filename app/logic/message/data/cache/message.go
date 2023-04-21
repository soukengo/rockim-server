package cache

import (
	"context"
	"github.com/soukengo/gopkg/component/cache"
	"rockimserver/apis/rockim/service/message/v1/types"
)

type MessageData struct {
	cache cache.ValueCache[types.IMMessage]
}

func NewMessageData(rds *cache.Manager, cfg *cache.Config) *MessageData {
	return &MessageData{
		cache: cache.NewValueCache[types.IMMessage](rds, cfg.Category(keyMessage)),
	}
}

func (d *MessageData) Save(ctx context.Context, message *types.IMMessage) (err error) {
	return d.cache.Set(ctx, cache.Parts(message.ProductId, message.MsgId), message)
}
func (d *MessageData) List(ctx context.Context, productId string, msgIds []string) (list []*types.IMMessage, err error) {
	return d.cache.List(ctx, cache.Parts(productId), msgIds)
}
