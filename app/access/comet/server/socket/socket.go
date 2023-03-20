package socket

import (
	v1 "rockimserver/apis/rockim/api/client/v1/socket"
	"rockimserver/app/access/comet/conf"
	"rockimserver/app/access/comet/module/client/service"
	"rockimserver/pkg/component/server"
	"rockimserver/pkg/component/server/socket"
	"rockimserver/pkg/component/server/socket/options"
)

func NewSocketServer(c *server.Config, pc *conf.Protocol, channelSrv *service.ChannelService) socket.Server {
	handler := NewClientHandler(pc, channelSrv)
	srv := socket.NewServer(handler, options.WithPacketFactory(v1.NewPacketFactory()))
	srv.RegisterTCPServer(c.Socket.TCP)
	srv.RegisterWSServer(c.Socket.WS)
	handler.server = srv
	return srv
}
