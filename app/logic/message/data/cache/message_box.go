package cache

import (
	"context"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/database/redis"
	"github.com/soukengo/gopkg/log"
	biztypes "rockimserver/app/logic/message/biz/types"
)

type MessageBoxData struct {
	cache cache.SortedSetCache[biztypes.IMMessageLetter]
}

func NewMessageBoxData(rds *redis.Client, cfg *cache.Config) *MessageBoxData {
	return &MessageBoxData{
		cache: newSortedSetCache[biztypes.IMMessageLetter](rds, cfg.Category(keyMessageBox)),
	}
}

func (d *MessageBoxData) SaveLetters(ctx context.Context, productId string, letters []*biztypes.IMMessageLetter) (err error) {
	for _, letter := range letters {
		err1 := d.cache.Add(ctx, cache.Parts(productId, letter.Uid), cache.Member(letter.Timestamp, letter))
		if err1 != nil {
			log.WithContext(ctx).Errorf("d.cache.Add error: %v", err1)
		}
	}
	return
}

func (d *MessageBoxData) Tail(ctx context.Context, productId string, uid string, size int64) (list []*biztypes.IMMessageLetter, err error) {
	return d.cache.Tail(ctx, cache.Parts(productId, uid), size)
}

func (d *MessageBoxData) Clear(ctx context.Context, productId string, uid string, timestamp int64) (err error) {
	return d.cache.ClearByScore(ctx, cache.Parts(productId, uid), timestamp, true)
}
