package listener

import (
	"context"
	comettypes "rockimserver/apis/rockim/service/comet/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/group/biz"
	bizevent "rockimserver/app/logic/group/biz/event"
	"rockimserver/app/logic/group/biz/options"
	"time"
)

type GroupEventListener struct {
	uc *biz.CometUseCase
}

func NewGroupEventListener(uc *biz.CometUseCase) *GroupEventListener {
	return &GroupEventListener{uc: uc}
}

func (lis *GroupEventListener) OnGroupJoined(ctx context.Context, data *bizevent.GroupJoinedEvent) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	lis.uc.SendControl(ctx, &options.ControlOptions{
		ProductId:   data.ProductId,
		ControlType: enums.Comet_ROOM_JOIN,
		Uids:        []string{data.Uid},
		Data: &comettypes.RoomJoinControl{
			Rooms: []*comettypes.Room{{RoomType: enums.Comet_GROUP, BizId: data.GroupId}},
		},
	})
}
