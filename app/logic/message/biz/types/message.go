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
	conversationUnknown = &types.ConversationID{Category: enums.Conversation_UNKNOWN}
)

type MessageMeta struct {
	MsgId     string
	ProductId string
	Sender    *usertypes.User
	Target    *types.TargetID
	RecvUser  *usertypes.User
	Group     *grouptypes.Group
	Body      *types.IMMessageBody
	Receivers []string
}

func NewPersonMessageMeta(productId string, sender *usertypes.User, target *types.TargetID, recvUser *usertypes.User, body *types.IMMessageBody) *MessageMeta {
	return &MessageMeta{ProductId: productId, Sender: sender, Target: target, RecvUser: recvUser, Body: body}
}

func NewGroupMessageMeta(productId string, sender *usertypes.User, target *types.TargetID, group *grouptypes.Group, body *types.IMMessageBody) *MessageMeta {
	return &MessageMeta{ProductId: productId, Sender: sender, Target: target, Group: group, Body: body}
}

func (m *MessageMeta) ConversationId() *types.ConversationID {
	switch m.Target.Category {
	case enums.Conversation_PERSON:
		uids := []string{m.Sender.Id, m.RecvUser.Id}
		sort.Strings(uids)
		return &types.ConversationID{Category: m.Target.Category, Value: strings.Join(uids, conversationIdValueSeparator)}
	case enums.Conversation_GROUP:
		return &types.ConversationID{Category: m.Target.Category, Value: m.Group.Id}
	default:
		return &types.ConversationID{Category: enums.Conversation_UNKNOWN}
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
	categoryNum, ok := enums.Conversation_Category_value[category]
	if !ok {
		id = conversationUnknown
		return
	}
	id = &types.ConversationID{Category: enums.Conversation_Category(categoryNum), Value: value}
	return
}
