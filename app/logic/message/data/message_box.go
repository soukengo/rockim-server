package data

import (
	"context"
	"rockimserver/app/logic/message/biz"
	biztypes "rockimserver/app/logic/message/biz/types"
	"rockimserver/app/logic/message/data/cache"
)

type messageBoxRepo struct {
	cache *cache.MessageBoxData
}

func NewMessageBoxRepo(cache *cache.MessageBoxData) biz.MessageBoxRepo {
	return &messageBoxRepo{cache: cache}
}

func (r *messageBoxRepo) SaveLetters(ctx context.Context, productId string, letters []*biztypes.IMMessageLetter) error {
	return r.cache.SaveLetters(ctx, productId, letters)
}

func (r *messageBoxRepo) Tail(ctx context.Context, productId string, uid string, size int64) (list []*biztypes.IMMessageLetter, err error) {
	return r.cache.Tail(ctx, productId, uid, size)
}

func (r *messageBoxRepo) Clear(ctx context.Context, productId string, uid string, timestamp int64) (err error) {
	return r.cache.Clear(ctx, productId, uid, timestamp)
}
