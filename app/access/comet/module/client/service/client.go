package service

import (
	"context"
	"rockimserver/apis/rockim/api/client/socket/v1"
	"rockimserver/app/access/comet/module/client/biz"
	"rockimserver/app/access/comet/module/client/biz/options"
	"rockimserver/pkg/component/server/socket"
)

type ClientService struct {
	v1.UnimplementedClientAPIServer
	uc *biz.ClientUseCase
}

func NewClientService(uc *biz.ClientUseCase) *ClientService {
	return &ClientService{uc: uc}
}

func (s *ClientService) Connect(ctx *socket.Context) (err error) {
	return s.uc.Connect(ctx)
}

func (s *ClientService) DisConnect(ctx *socket.Context) (err error) {
	return s.uc.DisConnect(ctx)
}

func (s *ClientService) Auth(ctx context.Context, in *v1.AuthRequest) (out *v1.AuthResponse, err error) {
	err = s.uc.Auth(ctx, &options.ClientAuthOptions{
		ProductId: in.ProductId,
		Token:     in.Token,
	})
	if err != nil {
		return
	}
	out = &v1.AuthResponse{}
	return
}

func (s *ClientService) HeartBeat(ctx context.Context, in *v1.HeartBeatRequest) (out *v1.HeartBeatResponse, err error) {
	err = s.uc.HeartBeat(ctx)
	if err != nil {
		return
	}
	out = &v1.HeartBeatResponse{}
	return
}
