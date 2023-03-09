package grpc

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "rockimserver/apis/rockim/service/comet/v1"
	"rockimserver/app/access/comet/module/server/service"
)

type ServiceGroup struct {
	channelSrv *service.ChannelService
}

func NewServiceGroup(channelSrv *service.ChannelService) *ServiceGroup {
	return &ServiceGroup{channelSrv: channelSrv}
}

func (g *ServiceGroup) Register(srv *grpc.Server) {
	v1.RegisterChannelAPIServer(srv, g.channelSrv)
}
