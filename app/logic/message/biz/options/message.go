package options

import (
	"rockimserver/apis/rockim/service/message/v1/types"
	"rockimserver/apis/rockim/shared/enums"
)

type MessageSendOptions struct {
	ProductId      string                    // 应用ID
	Uid            string                    // 发送者用户ID
	ConversationId *types.ConversationID     // 会话ID
	ClientMsgId    string                    // 客户端的消息ID
	MessageType    enums.Message_MessageType // 消息类型
	Content        string                    // 消息内容
	Payload        map[string]string         // 透传数据
}
