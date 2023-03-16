package socket

import "rockimserver/pkg/component/server/socket"

type ChannelManager struct {
	server socket.Server
}

func NewChannelManager(server socket.Server) *ChannelManager {
	return &ChannelManager{server: server}
}
