package task

import (
	"context"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/message/biz"
	biztypes "rockimserver/app/logic/message/biz/types"
)

type MessageTask struct {
	uc *biz.MessageDeliveryUseCase
}

func NewMessageTask(uc *biz.MessageDeliveryUseCase) *MessageTask {
	return &MessageTask{uc: uc}
}

func (t *MessageTask) Delivery(ctx context.Context, topic string, data []byte) (err error) {
	value := string(data)
	productId, conversationId := biztypes.DecodeConversationID(value)
	if conversationId.Category == enums.Conversation_UNKNOWN {
		return
	}
	return t.uc.Delivery(ctx, productId, conversationId)
}
