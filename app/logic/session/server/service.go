package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "rockimserver/apis/rockim/service/session/v1"
	"rockimserver/app/logic/session/service"
)

type ServiceRegistry struct {
	channelSrv      *service.ChannelService
	channelQuerySrv *service.ChannelQueryService
}

func NewServiceRegistry(channelSrv *service.ChannelService, channelQuerySrv *service.ChannelQueryService) *ServiceRegistry {
	return &ServiceRegistry{channelSrv: channelSrv, channelQuerySrv: channelQuerySrv}
}

func (g *ServiceRegistry) RegisterGrpc(srv *grpc.Server) {
	v1.RegisterChannelAPIServer(srv, g.channelSrv)
	v1.RegisterChannelQueryAPIServer(srv, g.channelQuerySrv)
}
