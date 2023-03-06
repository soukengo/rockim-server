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
	srv.RegisterTCPServer(c.Socket.TCP)
	srv.RegisterWSServer(c.Socket.WS)
	handler.server = srv
	return srv
}
