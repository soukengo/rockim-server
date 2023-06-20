package types

import (
	grouptypes "rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/service/message/v1/types"
	usertypes "rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"sort"
	"strings"
)

const (
	conversationIdValueSeparator = "#" // 会话id分隔符
	conversationIdSeparator      = ":" // 会话id分隔符
)

var (
	conversationUnknown = &types.ConversationID{Category: enums.MessageTarget_UNKNOWN}
)

type MessageMeta struct {
	MsgId     string
	ProductId string
	Sender    *usertypes.User
	From      *types.TargetID
	To        *types.TargetID
	RecvUser  *usertypes.User
	Group     *grouptypes.Group
	Body      *types.IMMessageBody
	Receivers []string
}

func NewPersonMessageMeta(productId string, sender *usertypes.User, from *types.TargetID, to *types.TargetID, recvUser *usertypes.User, body *types.IMMessageBody) *MessageMeta {
	return &MessageMeta{ProductId: productId, Sender: sender, From: from, To: to, RecvUser: recvUser, Body: body}
}

func NewGroupMessageMeta(productId string, sender *usertypes.User, target *types.TargetID, group *grouptypes.Group, body *types.IMMessageBody) *MessageMeta {
	return &MessageMeta{ProductId: productId, Sender: sender, From: target, To: target, Group: group, Body: body}
}

func (m *MessageMeta) ConversationId() *types.ConversationID {
	switch m.To.Category {
	case enums.MessageTarget_PERSON:
		uids := []string{m.Sender.Id, m.RecvUser.Id}
		sort.Strings(uids)
		return &types.ConversationID{Category: m.To.Category, Value: strings.Join(uids, conversationIdValueSeparator)}
	case enums.MessageTarget_GROUP:
		return &types.ConversationID{Category: m.To.Category, Value: m.Group.Id}
	default:
		return &types.ConversationID{Category: enums.MessageTarget_UNKNOWN}
	}
}

// EncodeConversationID encodes the conversation id
// dist: productId:category:value
func EncodeConversationID(productId string, id *types.ConversationID) (dist string) {
	return strings.Join([]string{productId, id.Category.String(), id.Value}, conversationIdSeparator)
}

// DecodeConversationID decodes the conversation id from the string
// source: productId:category:value
func DecodeConversationID(source string) (productId string, id *types.ConversationID) {
	arr := strings.Split(source, conversationIdSeparator)
	if len(arr) != 3 {
		id = conversationUnknown
		return
	}
	productId = arr[0]
	category := arr[1]
	value := arr[2]
	categoryNum, ok := enums.MessageTarget_Category_value[category]
	if !ok {
		id = conversationUnknown
		return
	}
	id = &types.ConversationID{Category: enums.MessageTarget_Category(categoryNum), Value: value}
	return
}

func PersonUids(id *types.ConversationID) (uids []string, valid bool) {
	uids = strings.Split(id.Value, conversationIdValueSeparator)
	valid = len(uids) == 2
	return
}

// IMMessageLetter 消息信件（同步消息）
type IMMessageLetter struct {
	// 会话ID
	ConversationId *types.ConversationID
	// 消息ID
	MsgId string
	// 用户ID
	Uid string
	// 时间戳
	Timestamp int64
}
