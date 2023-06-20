package socket

import (
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/component/server/socket"
	"github.com/soukengo/gopkg/component/server/socket/options"
	"rockimserver/app/access/comet/conf"
	"rockimserver/app/access/comet/module/client/service"
	"rockimserver/app/access/comet/protocol"
)

func NewSocketServer(c *server.Config, pc *conf.Protocol, channelSrv *service.ChannelService) socket.Server {
	handler := NewClientHandler(pc, channelSrv)
	srv := socket.NewServer(handler, options.WithPacketFactory(protocol.NewPacketFactory()))
	srv.RegisterTCPServer(c.Socket.TCP)
	srv.RegisterWSServer(c.Socket.WS)
	handler.server = srv
	return srv
}
