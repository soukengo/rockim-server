package task

import (
	"context"
	"rockimserver/apis/rockim/service/message/v1/types"
	"rockimserver/app/logic/message/biz"
)

type MessageTask struct {
	uc *biz.MessageDeliveryUseCase
}

func NewMessageTask(uc *biz.MessageDeliveryUseCase) *MessageTask {
	return &MessageTask{uc: uc}
}

func (t *MessageTask) Delivery(ctx context.Context, task *types.DeliveryTask) (err error) {
	return t.uc.Delivery(ctx, task)
}
