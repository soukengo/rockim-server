package biz

import (
	"context"
	"github.com/samber/lo"
	"github.com/soukengo/gopkg/log"
	clienttypes "rockimserver/apis/rockim/api/client/v1/types"
	"rockimserver/apis/rockim/service/message/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	biztypes "rockimserver/app/logic/message/biz/types"
)

type MessageDeliveryUseCase struct {
	repo      MessageDeliveryRepo
	queryRepo MessageQueryRepo
	pushRepo  PushRepo
}

func NewMessageDeliveryUseCase(repo MessageDeliveryRepo, queryRepo MessageQueryRepo, pushRepo PushRepo) *MessageDeliveryUseCase {
	return &MessageDeliveryUseCase{repo: repo, queryRepo: queryRepo, pushRepo: pushRepo}
}

func (uc *MessageDeliveryUseCase) Delivery(ctx context.Context, productId string, conversationId *types.ConversationID) (err error) {
	list, err := uc.repo.Pop(ctx, productId, conversationId, 10)
	if err != nil {
		return
	}
	if len(list) == 0 {
		return
	}
	msgIds := lo.Map(list, func(item *biztypes.IMMessageDelivery, index int) string {
		return item.MsgId
	})
	messages, err := uc.queryRepo.List(ctx, productId, msgIds)
	if err != nil {
		return
	}
	messageMap := make(map[string]*types.IMMessage)
	for _, message := range messages {
		messageMap[message.MsgId] = message
	}
	packet := &clienttypes.IMMessagePacket{List: make([]*clienttypes.IMMessage, 0)}
	for _, record := range list {
		message, ok := messageMap[record.MsgId]
		if !ok {
			continue
		}
		packet.List = append(packet.List, uc.toClientMessage(record, message))
	}

	switch conversationId.Category {
	case enums.MessageTarget_GROUP:
		err = uc.pushRepo.PushGroup(ctx, productId, conversationId.Value, packet)
	case enums.MessageTarget_PERSON:
		uids, valid := biztypes.PersonUids(conversationId)
		if !valid {
			log.WithContext(ctx).Errorf("conversationId is not valid: %v", conversationId)
			return
		}
		err = uc.pushRepo.PushUsers(ctx, productId, uids, packet)
	}
	return
}

func (uc *MessageDeliveryUseCase) toClientMessage(delivery *biztypes.IMMessageDelivery, message *types.IMMessage) *clienttypes.IMMessage {
	return &clienttypes.IMMessage{
		ProductId: message.ProductId,
		MsgId:     message.MsgId,
		Sender: &clienttypes.IMMessageSender{
			Account:   message.Sender.Account,
			Name:      message.Sender.Name,
			AvatarUrl: message.Sender.AvatarUrl,
		},
		Target: &clienttypes.TargetID{Category: delivery.From.Category, Value: delivery.From.Value},
		Body: &clienttypes.IMMessageBody{
			Timestamp:   message.Body.Timestamp,
			ClientMsgId: message.Body.ClientMsgId,
			MessageType: message.Body.MessageType,
			Content:     message.Body.Content,
			Payload:     message.Body.Payload,
		},
		Sequence: message.Sequence,
		Status:   message.Status,
		Version:  message.Version,
	}
}
