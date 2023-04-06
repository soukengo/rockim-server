package data

import (
	"context"
	"github.com/soukengo/gopkg/errors"
	"rockimserver/apis/rockim/service/message/v1/types"
	"rockimserver/app/logic/message/biz"
	biztypes "rockimserver/app/logic/message/biz/types"
	"rockimserver/app/logic/message/data/cache"
)

type messageDeliveryRepo struct {
	cache *cache.MessageDeliveryData
}

func NewMessageDeliveryRepo(cache *cache.MessageDeliveryData) biz.MessageDeliveryRepo {
	return &messageDeliveryRepo{cache: cache}
}

func (r *messageDeliveryRepo) Save(ctx context.Context, productId string, record *biztypes.IMMessageDelivery) error {
	return r.cache.Save(ctx, productId, record)
}

func (r *messageDeliveryRepo) Pop(ctx context.Context, productId string, conversationId *types.ConversationID, size int64) (list []*biztypes.IMMessageDelivery, err error) {
	list, err = r.cache.Pop(ctx, productId, conversationId, size)
	if err != nil {
		if errors.IsNotFound(err) {
			err = nil
		}
		return
	}
	return
}
