package options

import (
	"rockimserver/apis/rockim/service/message/v1/types"
	"rockimserver/apis/rockim/shared/enums"
)

type MessageSendOptions struct {
	ProductId   string                    // 应用ID
	From        string                    // 来源ID
	Target      *types.TargetID           // 目标ID
	IsSystem    bool                      // 是否是系统消息
	ClientMsgId string                    // 客户端的消息ID
	MessageType enums.Message_MessageType // 消息类型
	Content     []byte                    // 消息内容
	Payload     map[string]string         // 透传数据
}
