package event

import (
	"github.com/soukengo/gopkg/component/transport/event"
)

const (
	KeyGroupJoined = event.Key("GroupJoined")
)

type GroupJoinedEvent struct {
	ProductId string
	GroupId   string
	Uid       string
}

func NewGroupJoinedEvent(productId string, groupId string, uid string) *GroupJoinedEvent {
	return &GroupJoinedEvent{ProductId: productId, GroupId: groupId, Uid: uid}
}

func (g GroupJoinedEvent) Key() event.Key {
	return KeyGroupJoined
}
