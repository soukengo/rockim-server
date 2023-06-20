package entity

import (
	"rockimserver/apis/rockim/shared/enums"
)

const (
	TableImMessage         = "im_message"
	TableImMessageSequence = "im_message_sequence"
	TableImMessageRelation = "im_message_relation"
)

type IMMessage struct {
	Id             string `bson:"_id,omitempty"`         // ID
	UpdateTime     int64  `bson:"update_time,omitempty"` // 更新时间
	ProductId      string `bson:"product_id"`
	MsgId          string `bson:"msg_id"`
	ConversationId string `bson:"conversation_id"`
	// 消息发送者
	Sender *IMMessageSender `bson:"sender"`
	// 来源目标ID
	From *MessageTargetID `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	// 接收者的目标ID
	To *MessageTargetID `protobuf:"bytes,6,opt,name=to,proto3" json:"to,omitempty"`
	// 消息主体
	Body     *IMMessageBody `bson:"body"`
	Sequence int64          `bson:"sequence"`
	Status   int32          `bson:"status"`
	Version  int64          `bson:"version"`
}

func (*IMMessage) TableName() string {
	return TableImMessage
}

type IMMessageBody struct {
	// 消息发送时间
	Timestamp int64 `bson:"timestamp"`
	// 客户端的消息ID
	ClientMsgId string `bson:"client_msg_id"`
	// 消息类型
	MessageType int32 `bson:"message_type"`
	// 消息内容
	Content []byte `bson:"content"`
	// 透传数据
	Payload map[string]string `bson:"payload"`
}

// IMMessageSender 消息发送者
type IMMessageSender struct {
	// 发送者uid
	Uid string `bson:"uid"`
	// 发送者账号
	Account string `bson:"account"`
	// 发送者名称
	Name string `bson:"name"`
	// 发送者头像
	AvatarUrl string `bson:"avatar_url"`
}

type MessageTargetID struct {
	// 目标类型
	Category enums.MessageTarget_Category `bson:"category"`
	// 值（群聊为bizId，私聊为用户account）
	Value string `bson:"value"`
}

type IMMessageSequence struct {
	Id  string `bson:"_id,omitempty"` // Id
	Seq int64  `bson:"seq"`
}

func (*IMMessageSequence) TableName() string {
	return TableImMessageSequence
}

// IMMessageRelation 消息关联
type IMMessageRelation struct {
	Id string `bson:"_id,omitempty"` // ID
	// 所属应用
	ProductId string `bson:"product_id"`
	// 会话ID
	ConversationId string `bson:"conversation_id"`
	// 消息ID
	MsgId string `bson:"msg_id"`
	// 用户ID
	Uid string `bson:"uid"`
	// 是否已删除
	Deleted bool `bson:"deleted"`
}

func (*IMMessageRelation) TableName() string {
	return TableImMessageRelation
}
