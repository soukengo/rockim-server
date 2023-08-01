package event

import (
	"context"
	"rockimserver/app/logic/group/biz"
	"rockimserver/pkg/component/event"
)

type GroupEventListener struct {
	repo biz.EventRepo
}

type GroupJoinedEvent struct {
}

func (g GroupJoinedEvent) Key() event.Key {
	return "GroupJoined"
}

func (lis *GroupEventListener) name() {
	lis.repo.Subscribe("GroupJoined", event.Listener{
		Mode: event.Async,
		Handler: event.Wrap(func(ctx context.Context, t *GroupJoinedEvent) {

		}),
	})
}
