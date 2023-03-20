package socket

import (
	"rockimserver/pkg/component/server/socket/network"
	"rockimserver/pkg/component/server/socket/network/tcp"
	"rockimserver/pkg/component/server/socket/network/tcp/nbio"
	"rockimserver/pkg/component/server/socket/network/ws"
)

func (s *server) RegisterTCPServer(cfg *tcp.Config) {
	ins := nbio.NewServer(cfg, s.opt.Parser)
	//s := gnet.NewServer(cfg, s.opt.Parser)
	s.register(ins)
	return
}
func (s *server) RegisterWSServer(cfg *ws.Config) {
	ins := ws.NewServer(cfg, s.opt.Parser)
	s.register(ins)
	return
}

func (s *server) register(srv network.Server) {
	srvId := srv.Id()
	s.servers[srvId] = srv
	srv.SetHandler(s)
}
