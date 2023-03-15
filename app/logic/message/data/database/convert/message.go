package convert

import (
	"rockimserver/apis/rockim/service/message/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	biztypes "rockimserver/app/logic/message/biz/types"
	"rockimserver/app/logic/message/data/database/entity"
)

func MessageProto(source *entity.IMMessage) *types.IMMessage {
	if source == nil {
		return nil
	}
	_, conversationId := biztypes.DecodeConversationID(source.ConversationId)
	return &types.IMMessage{
		ProductId:      source.ProductId,
		MsgId:          source.MsgId,
		ConversationId: conversationId,
		Body:           MessageBodyProto(source.Body),
		Sequence:       source.Sequence,
		Status:         enums.Message_Status(source.Status),
		Version:        source.Version,
	}
}

func MessageBodyProto(source *entity.IMMessageBody) *types.IMMessageBody {
	if source == nil {
		return nil
	}
	return &types.IMMessageBody{
		Timestamp:   source.Timestamp,
		Sender:      MessageSenderProto(source.Sender),
		ClientMsgId: source.ClientMsgId,
		MessageType: enums.Message_MessageType(source.MessageType),
		Content:     source.Content,
		Payload:     source.Payload,
		NeedReceipt: source.NeedReceipt,
	}
}
func MessageSenderProto(source *entity.IMMessageSender) *types.IMMessageSender {
	if source == nil {
		return nil
	}
	return &types.IMMessageSender{
		Uid:       source.Uid,
		Account:   source.Account,
		Name:      source.Name,
		AvatarUrl: source.AvatarUrl,
	}
}

func MessageEntity(source *types.IMMessage) *entity.IMMessage {
	if source == nil {
		return nil
	}
	return &entity.IMMessage{
		Id:             source.MsgId,
		ProductId:      source.ProductId,
		MsgId:          source.MsgId,
		ConversationId: biztypes.EncodeConversationID(source.ProductId, source.ConversationId),
		Body:           MessageBodyEntity(source.Body),
		Sequence:       source.Sequence,
		Status:         int32(source.Status),
		Version:        source.Version,
	}
}

func MessageBodyEntity(source *types.IMMessageBody) *entity.IMMessageBody {
	if source == nil {
		return nil
	}
	return &entity.IMMessageBody{
		Timestamp:   source.Timestamp,
		Sender:      MessageSenderEntity(source.Sender),
		ClientMsgId: source.ClientMsgId,
		MessageType: int32(source.MessageType),
		Content:     source.Content,
		Payload:     source.Payload,
		NeedReceipt: source.NeedReceipt,
	}
}
func MessageSenderEntity(source *types.IMMessageSender) *entity.IMMessageSender {
	if source == nil {
		return nil
	}
	return &entity.IMMessageSender{
		Uid:       source.Uid,
		Account:   source.Account,
		Name:      source.Name,
		AvatarUrl: source.AvatarUrl,
	}
}
func MessageRelationEntity(source *types.IMMessageRelation) *entity.IMMessageRelation {
	if source == nil {
		return nil
	}
	return &entity.IMMessageRelation{
		Id:             source.Id,
		ProductId:      source.ProductId,
		ConversationId: biztypes.EncodeConversationID(source.ProductId, source.ConversationId),
		MsgId:          source.MsgId,
		Uid:            source.Uid,
		Deleted:        source.Deleted,
	}
}
func MessageRelationProto(source *entity.IMMessageRelation) *types.IMMessageRelation {
	if source == nil {
		return nil
	}
	_, conversationId := biztypes.DecodeConversationID(source.ConversationId)
	return &types.IMMessageRelation{
		Id:             source.Id,
		ProductId:      source.ProductId,
		ConversationId: conversationId,
		MsgId:          source.MsgId,
		Uid:            source.Uid,
		Deleted:        source.Deleted,
	}
}
