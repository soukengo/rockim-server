package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "rockimserver/apis/rockim/service/comet/v1"
	"rockimserver/app/access/comet/module/server/service"
)

type GrpcRegistry struct {
	channelSrv *service.ChannelService
}

func NewServiceRegistry(channelSrv *service.ChannelService) *GrpcRegistry {
	return &GrpcRegistry{channelSrv: channelSrv}
}

func (g *GrpcRegistry) RegisterGrpc(srv *grpc.Server) {
	v1.RegisterChannelAPIServer(srv, g.channelSrv)
}
