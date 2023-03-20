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

func (s *ChannelService) Push(ctx context.Context, in *v1.PushRequest) (out *v1.PushResponse, err error) {
	err = s.uc.Push(ctx, &options.PushOptions{
		ChannelIds: in.ChannelIds,
		Operation:  in.Operation,
		Body:       in.Body,
	})
	if err != nil {
		return
	}
	out = &v1.PushResponse{}
	return
}

func (s *ChannelService) PushGroup(ctx context.Context, in *v1.PushGroupRequest) (out *v1.PushGroupResponse, err error) {
	err = s.uc.PushGroup(ctx, &options.PushGroupOptions{
		GroupId:   in.GroupId,
		Operation: in.Operation,
		Body:      in.Body,
	})
	if err != nil {
		return
	}
	out = &v1.PushGroupResponse{}
	return
}
