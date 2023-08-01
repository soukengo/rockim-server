package event

import (
	"github.com/soukengo/gopkg/component/transport/event"
)

const (
	KeyGroupJoined = event.Key("GroupJoined")
)

type GroupJoinedEvent struct {
	GroupId string
}

func NewGroupJoinedEvent(groupId string) *GroupJoinedEvent {
	return &GroupJoinedEvent{GroupId: groupId}
}

func (g GroupJoinedEvent) Key() event.Key {
	return KeyGroupJoined
}
