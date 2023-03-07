package service

import (
	"context"
	"rockimserver/apis/rockim/api/client/socket/v1"
	"rockimserver/app/access/comet/module/client/biz"
	"rockimserver/app/access/comet/module/client/biz/options"
	"rockimserver/pkg/component/server/socket"
)

type ChannelService struct {
	v1.UnimplementedChannelAPIServer
	uc *biz.ChannelUseCase
}

func NewClientService(uc *biz.ChannelUseCase) *ChannelService {
	return &ChannelService{uc: uc}
}

func (s *ChannelService) Connect(ctx *socket.Context) (err error) {
	return s.uc.Connect(ctx)
}

func (s *ChannelService) DisConnect(ctx *socket.Context) (err error) {
	return s.uc.DisConnect(ctx)
}

func (s *ChannelService) Auth(ctx context.Context, in *v1.AuthRequest) (out *v1.AuthResponse, err error) {
	err = s.uc.Auth(ctx, &options.ChannelAuthOptions{
		ProductId: in.ProductId,
		Token:     in.Token,
	})
	if err != nil {
		return
	}
	out = &v1.AuthResponse{}
	return
}

func (s *ChannelService) HeartBeat(ctx context.Context, in *v1.HeartBeatRequest) (out *v1.HeartBeatResponse, err error) {
	err = s.uc.HeartBeat(ctx)
	if err != nil {
		return
	}
	out = &v1.HeartBeatResponse{}
	return
}
