package biz

import (
	"context"
	"github.com/golang/protobuf/proto"
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
	pushMgr   PushManager
}

func NewMessageDeliveryUseCase(repo MessageDeliveryRepo, queryRepo MessageQueryRepo, pushMgr PushManager) *MessageDeliveryUseCase {
	return &MessageDeliveryUseCase{repo: repo, queryRepo: queryRepo, pushMgr: pushMgr}
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
	body, err := proto.Marshal(packet)
	if err != nil {
		return
	}
	operation := enums.Network_MESSAGES
	switch conversationId.Category {
	case enums.Conversation_GROUP:
		err = uc.pushMgr.PushGroup(ctx, operation, productId, conversationId.Value, body)
	case enums.Conversation_PERSON:
		uids, valid := biztypes.PersonUids(conversationId)
		if !valid {
			log.WithContext(ctx).Errorf("conversationId is not valid: %v", conversationId)
			return
		}
		err = uc.pushMgr.PushUsers(ctx, operation, productId, uids, body)
	}
	return
}

func (uc *MessageDeliveryUseCase) toClientMessage(delivery *biztypes.IMMessageDelivery, message *types.IMMessage) *clienttypes.IMMessage {
	return &clienttypes.IMMessage{
		ProductId: message.ProductId,
		MsgId:     message.MsgId,
		From:      &clienttypes.TargetID{Category: delivery.From.Category, Value: delivery.From.Value},
		To:        &clienttypes.TargetID{Category: delivery.To.Category, Value: delivery.To.Value},
		Body: &clienttypes.IMMessageBody{
			Timestamp: message.Body.Timestamp,
			Sender: &clienttypes.IMMessageSender{
				Uid:       message.Body.Sender.Uid,
				Account:   message.Body.Sender.Account,
				Name:      message.Body.Sender.Name,
				AvatarUrl: message.Body.Sender.AvatarUrl,
			},
			ClientMsgId: message.Body.ClientMsgId,
			MessageType: message.Body.MessageType,
			Content:     message.Body.Content,
			Payload:     message.Body.Payload,
			NeedReceipt: message.Body.NeedReceipt,
		},
		Sequence: message.Sequence,
		Status:   message.Status,
		Version:  message.Version,
	}
}
