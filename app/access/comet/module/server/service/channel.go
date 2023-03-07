package service

import (
	"context"
	"rockimserver/apis/rockim/service/comet/v1"
)

type ChannelService struct {
	v1.UnimplementedChannelAPIServer
}

func (s *ChannelService) Push(ctx context.Context, request *v1.PushRequest) (*v1.PushResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChannelService) PushGroup(ctx context.Context, request *v1.PushGroupRequest) (*v1.PushGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}
