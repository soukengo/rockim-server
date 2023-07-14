package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "rockimserver/apis/rockim/service/session/v1"
	"rockimserver/app/logic/session/service"
)

type ServiceGroup struct {
	channelSrv      *service.ChannelService
	channelQuerySrv *service.ChannelQueryService
}

func NewServiceGroup(channelSrv *service.ChannelService, channelQuerySrv *service.ChannelQueryService) *ServiceGroup {
	return &ServiceGroup{channelSrv: channelSrv, channelQuerySrv: channelQuerySrv}
}

func (g *ServiceGroup) RegisterGrpc(srv *grpc.Server) {
	v1.RegisterChannelAPIServer(srv, g.channelSrv)
	v1.RegisterChannelQueryAPIServer(srv, g.channelQuerySrv)
}
