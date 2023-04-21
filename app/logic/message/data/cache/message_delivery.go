package cache

import (
	"context"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service/message/v1/types"
	biztypes "rockimserver/app/logic/message/biz/types"
)

type MessageDeliveryData struct {
	maxSize uint64
	cache   cache.SortedSetCache[biztypes.IMMessageDelivery]
}

func NewMessageDeliveryData(mgr *cache.Manager, cfg *cache.Config) *MessageDeliveryData {
	return &MessageDeliveryData{
		maxSize: 100,
		cache:   cache.NewSortedSetCache[biztypes.IMMessageDelivery](mgr, cfg.Category(keyMessageDelivery)),
	}
}

func (d *MessageDeliveryData) Save(ctx context.Context, productId string, record *biztypes.IMMessageDelivery) (err error) {
	keyParts := cache.Parts(biztypes.EncodeConversationID(productId, record.ConversationId))
	err = d.cache.Add(ctx, keyParts, cache.Member[biztypes.IMMessageDelivery](record.Timestamp, record))
	if err != nil {
		return
	}
	err1 := d.cache.Clear(ctx, keyParts, d.maxSize, false)
	if err1 != nil {
		log.WithContext(ctx).Errorf("cache.Delete error: %v", err1)
	}
	return
}

func (d *MessageDeliveryData) Pop(ctx context.Context, productId string, conversationId *types.ConversationID, size int64) (list []*biztypes.IMMessageDelivery, err error) {
	keyParts := cache.Parts(biztypes.EncodeConversationID(productId, conversationId))
	list, err = d.cache.Tail(ctx, keyParts, size)
	if err != nil {
		return
	}
	if len(list) > 0 {
		err1 := d.cache.Delete(ctx, keyParts)
		if err1 != nil {
			log.WithContext(ctx).Errorf("cache.Delete error: %v", err1)
		}
	}
	return
}
