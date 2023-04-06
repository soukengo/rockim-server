package data

import (
	"context"
	"github.com/soukengo/gopkg/util/chain"
	"rockimserver/apis/rockim/service/message/v1/types"
	"rockimserver/app/logic/message/biz"
	"rockimserver/app/logic/message/data/cache"
	"rockimserver/app/logic/message/data/database"
)

type messageRepo struct {
	cache *cache.MessageData
	db    *database.MessageData
}

func NewMessageRepo(cache *cache.MessageData, db *database.MessageData) biz.MessageRepo {
	return &messageRepo{cache: cache, db: db}
}

func (r *messageRepo) GenSequence(ctx context.Context, productId string, conversation *types.ConversationID) (int64, error) {
	return r.db.GenSequence(ctx, productId, conversation)
}

func (r *messageRepo) Save(ctx context.Context, message *types.IMMessage) error {
	return chain.Call(func() (err error) {
		return r.db.Save(ctx, message)
	}, func() (err error) {
		return r.cache.Save(ctx, message)
	})
}

func (r *messageRepo) SaveRelations(ctx context.Context, productId string, relations []*types.IMMessageRelation) error {
	return r.db.SaveRelations(ctx, productId, relations)
}
