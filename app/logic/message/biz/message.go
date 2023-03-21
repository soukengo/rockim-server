package biz

import (
	"context"
	"rockimserver/apis/rockim/service/message/v1/types"
	biztypes "rockimserver/app/logic/message/biz/types"
)

type MessageRepo interface {
	GenSequence(ctx context.Context, productId string, conversation *types.ConversationID) (int64, error)
	Save(ctx context.Context, message *types.IMMessage) error
	SaveRelations(ctx context.Context, productId string, relations []*types.IMMessageRelation) error
}

type MessageQueryRepo interface {
	List(ctx context.Context, productId string, msgIds []string) ([]*types.IMMessage, error)
	ListFromCache(ctx context.Context, productId string, msgIds []string) ([]*types.IMMessage, error)
}

type MessageDeliveryRepo interface {
	Save(ctx context.Context, productId string, record *biztypes.IMMessageDelivery) error
	Pop(ctx context.Context, productId string, conversationId *types.ConversationID, size int64) (list []*biztypes.IMMessageDelivery, err error)
}

// MessageBoxRepo 信箱
type MessageBoxRepo interface {
	SaveLetters(ctx context.Context, productId string, letters []*biztypes.IMMessageLetter) error
	Tail(ctx context.Context, productId string, uid string, size int64) (list []*biztypes.IMMessageLetter, err error)
	Clear(ctx context.Context, productId string, uid string, timestamp int64) (err error)
}
