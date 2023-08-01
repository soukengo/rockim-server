package server

import (
	"github.com/soukengo/gopkg/component/transport/socket"
	"github.com/soukengo/gopkg/component/transport/socket/options"
	"rockimserver/app/access/comet/conf"
	"rockimserver/app/access/comet/module/protocol"
)

func NewSocketServer(c *conf.Server) socket.Server {
	srv := socket.NewServer(options.WithPacketParser(protocol.NewPacketParser()))
	srv.RegisterTCPServer(c.Socket.TCP)
	srv.RegisterWSServer(c.Socket.WS)
	return srv
}
