package listener

import (
	"context"
	"rockimserver/app/logic/group/biz"
	bizevent "rockimserver/app/logic/group/biz/event"
)

type GroupEventListener struct {
	biz *biz.GroupUseCase
}

func NewGroupEventListener(biz *biz.GroupUseCase) *GroupEventListener {
	return &GroupEventListener{biz: biz}
}

func (lis *GroupEventListener) OnGroupJoined(ctx context.Context, data *bizevent.GroupJoinedEvent) {

}
