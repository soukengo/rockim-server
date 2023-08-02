package data

import (
	"context"
	"github.com/golang/protobuf/proto"
	comettypes "rockimserver/apis/rockim/service/comet/v1/types"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/group/biz"
	"rockimserver/app/logic/group/biz/options"
)

type cometRepo struct {
	ac v1.CometAPIClient
}

func NewCometRepo(ac v1.CometAPIClient) biz.CometRepo {
	return &cometRepo{ac: ac}
}

func (r *cometRepo) PushUser(ctx context.Context, opts *options.PushOptions) (err error) {
	body, err := proto.Marshal(opts.Data)
	if err != nil {
		return
	}
	message := &types.CometMessage{
		ProductId: opts.ProductId,
		Target: &types.CometMessage_Target{
			TargetType: types.CometMessage_Target_USER,
			Uids:       opts.Uids,
		},
		Message: &comettypes.Message{
			DataType: enums.Comet_PUSH,
			Push: &comettypes.PushMessage{
				Operation: opts.Operation,
				Body:      body,
			},
		},
	}
	_, err = r.ac.Dispatch(ctx, &v1.CometDispatchRequest{Message: message})
	return
}

func (r *cometRepo) PushRoom(ctx context.Context, opts *options.PushRoomOptions) (err error) {
	body, err := proto.Marshal(opts.Data)
	if err != nil {
		return
	}
	message := &types.CometMessage{
		ProductId: opts.ProductId,
		Target: &types.CometMessage_Target{
			TargetType: types.CometMessage_Target_ROOM,
			Room:       opts.Room,
		},
		Message: &comettypes.Message{
			DataType: enums.Comet_PUSH,
			Push: &comettypes.PushMessage{
				Operation: opts.Operation,
				Body:      body,
			},
		},
	}
	_, err = r.ac.Dispatch(ctx, &v1.CometDispatchRequest{Message: message})
	return
}

func (r *cometRepo) SendControl(ctx context.Context, opts *options.ControlOptions) (err error) {
	body, err := proto.Marshal(opts.Data)
	if err != nil {
		return
	}
	message := &types.CometMessage{
		ProductId: opts.ProductId,
		Target: &types.CometMessage_Target{
			TargetType: types.CometMessage_Target_USER,
			Uids:       opts.Uids,
		},
		Message: &comettypes.Message{
			DataType: enums.Comet_CONTROL,
			Control: &comettypes.ControlMessage{
				ControlType: opts.ControlType,
				Body:        body,
			},
		},
	}
	_, err = r.ac.Dispatch(ctx, &v1.CometDispatchRequest{Message: message})
	return
}
