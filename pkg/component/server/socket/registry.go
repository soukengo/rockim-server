package socket

import (
	"rockimserver/pkg/component/server/socket/network"
	"rockimserver/pkg/component/server/socket/network/tcp"
	"rockimserver/pkg/component/server/socket/network/tcp/nbio"
	"rockimserver/pkg/component/server/socket/network/ws"
)

func (m *server) RegisterTCPServer(cfg *tcp.Config) {
	s := nbio.NewServer(cfg, m.opt.Parser)
	//s := gnet.NewServer(cfg, m.opt.Parser)
	m.register(s)
	return
}
func (m *server) RegisterWSServer(cfg *ws.Config) {
	s := ws.NewServer(cfg, m.opt.Parser)
	m.register(s)
	return
}

func (m *server) register(srv network.Server) {
	srvId := srv.Id()
	m.servers[srvId] = srv
	srv.SetHandler(m)
}
