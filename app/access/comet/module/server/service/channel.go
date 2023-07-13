package service

import (
	"context"
	"rockimserver/apis/rockim/service/comet/v1"
	"rockimserver/app/access/comet/module/server/biz"
	"rockimserver/app/access/comet/module/server/biz/options"
)

type ChannelService struct {
	v1.UnimplementedChannelAPIServer
	uc *biz.ChannelUseCase
}

func NewChannelService(uc *biz.ChannelUseCase) *ChannelService {
	return &ChannelService{uc: uc}
}

func (s *ChannelService) Dispatch(ctx context.Context, in *v1.DispatchRequest) (out *v1.DispatchResponse, err error) {
	err = s.uc.Push(ctx, &options.PushOptions{
		ChannelIds: in.ChannelIds,
		Operation:  in.Push.Operation,
		Body:       in.Push.Body,
	})
	if err != nil {
		return
	}
	out = &v1.DispatchResponse{}
	return
}

func (s *ChannelService) DispatchRoom(ctx context.Context, in *v1.DispatchRoomRequest) (out *v1.DispatchRoomResponse, err error) {
	err = s.uc.PushRoom(ctx, &options.PushRoomOptions{
		Room:      in.Room,
		Operation: in.Push.Operation,
		Body:      in.Push.Body,
	})
	if err != nil {
		return
	}
	out = &v1.DispatchRoomResponse{}
	return
}
