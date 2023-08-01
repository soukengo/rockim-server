package data

import (
	"context"
	"google.golang.org/protobuf/proto"
	clienttypes "rockimserver/apis/rockim/api/client/v1/types"
	comettypes "rockimserver/apis/rockim/service/comet/v1/types"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/message/biz"
)

type pushMessageRepo struct {
	ac v1.CometAPIClient
}

func NewPushMessageRepo(ac v1.CometAPIClient) biz.PushRepo {
	return &pushMessageRepo{ac: ac}
}

func (r *pushMessageRepo) PushUsers(ctx context.Context, productId string, uids []string, packet *clienttypes.IMMessagePacket) (err error) {
	body, err := proto.Marshal(packet)
	if err != nil {
		return
	}
	operation := enums.Comet_MESSAGES
	message := &types.CometMessage{
		ProductId: productId,
		Target: &types.CometMessage_Target{
			TargetType: types.CometMessage_Target_USER,
			Uids:       uids,
		},
		Message: &comettypes.Message{
			Push: &comettypes.PushMessage{
				Operation: operation,
				Body:      body,
			},
		},
	}
	_, err = r.ac.DispatchAsync(ctx, &v1.CometDispatchRequest{Message: message})
	return
}

func (r *pushMessageRepo) PushGroup(ctx context.Context, productId string, groupId string, packet *clienttypes.IMMessagePacket) (err error) {
	body, err := proto.Marshal(packet)
	if err != nil {
		return
	}
	operation := enums.Comet_MESSAGES
	message := &types.CometMessage{
		ProductId: productId,
		Target: &types.CometMessage_Target{
			TargetType: types.CometMessage_Target_ROOM,
			Room:       &comettypes.Room{RoomType: enums.Comet_GROUP, BizId: groupId},
		},
		Message: &comettypes.Message{
			Push: &comettypes.PushMessage{
				Operation: operation,
				Body:      body,
			},
		},
	}
	_, err = r.ac.DispatchAsync(ctx, &v1.CometDispatchRequest{Message: message})
	return
}
