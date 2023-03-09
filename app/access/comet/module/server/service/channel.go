package service

import (
	"context"
	"rockimserver/apis/rockim/service/comet/v1"
	"rockimserver/app/access/comet/module/server/biz"
)

type ChannelService struct {
	v1.UnimplementedChannelAPIServer
	uc *biz.ChannelUseCase
}

func NewChannelService(uc *biz.ChannelUseCase) *ChannelService {
	return &ChannelService{uc: uc}
}

func (s *ChannelService) Push(ctx context.Context, request *v1.PushRequest) (*v1.PushResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChannelService) PushGroup(ctx context.Context, request *v1.PushGroupRequest) (*v1.PushGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}
