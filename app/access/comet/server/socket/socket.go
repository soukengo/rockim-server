package socket

import (
	v1 "rockimserver/apis/rockim/api/client/socket/v1"
	"rockimserver/pkg/component/server"
	"rockimserver/pkg/component/server/socket"
	"rockimserver/pkg/component/server/socket/options"
)

func NewSocketServer(c *server.Config) socket.Server {
	handler := NewClientHandler()
	srv := socket.NewServer(handler, options.WithPacketFactory(v1.NewPacketFactory()))
	srv.CreateTCPServer(c.Socket.TCP)
	srv.CreateWSServer(c.Socket.WS)
	handler.srv = srv
	return srv
}
