package types

import "rockimserver/apis/rockim/service/message/v1/types"

type IMMessageDelivery struct {
	ConversationId *types.ConversationID
	MsgId          string
	From           *types.TargetID // 来自哪个目标
	To             *types.TargetID // 去向哪个目标
	Timestamp      int64
}

func NewIMMessageDelivery(conversationId *types.ConversationID, msgId string, from *types.TargetID, to *types.TargetID) *IMMessageDelivery {
	return &IMMessageDelivery{ConversationId: conversationId, MsgId: msgId, From: from, To: to}
}
