package data

import (
	"context"
	"rockimserver/apis/rockim/service/message/v1/types"
	"rockimserver/app/logic/message/biz"
	"rockimserver/app/logic/message/data/database"
)

type messageRepo struct {
	db *database.MessageData
}

func NewMessageRepo(db *database.MessageData) biz.MessageRepo {
	return &messageRepo{db: db}
}

func (r *messageRepo) GenSequence(ctx context.Context, productId string, conversation *types.ConversationID) (int64, error) {
	return r.db.GenSequence(ctx, productId, conversation)
}

func (r *messageRepo) Save(ctx context.Context, message *types.IMMessage) error {
	return r.db.Save(ctx, message)
}

func (r *messageRepo) SaveRelations(ctx context.Context, relations []*types.IMMessageRelation) error {
	return r.db.SaveRelations(ctx, relations)
}
