package socket

import (
	"context"
	"rockimserver/pkg/component/server/socket/network/tcp"
	"rockimserver/pkg/component/server/socket/network/ws"
	"rockimserver/pkg/component/server/socket/packet"
)

type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
	Manager
	Register
}

type Manager interface {
	Channel(channelId string) (Channel, bool)
	Group(groupId string) (Group, bool)
	JoinGroup(groupId string, channel Channel) error
	QuitGroup(groupId string, channel Channel) error
}

type Register interface {
	CreateTCPServer(cfg *tcp.Config)
	CreateWSServer(cfg *ws.Config)
}

type Handler interface {
	OnCreated(Channel)
	OnClosed(Channel)
	OnReceived(Channel, packet.IPacket)
}

type Channel interface {
	Id() string
	ClientIP() string
	Send(packet packet.IPacket) error
	Close() error
	AddGroup(groupId string)
	DelGroup(groupId string)
	Groups() []string

	MarkAuthenticated()
	Authenticated() bool
	Attrs() map[string]any
	SetAttr(key string, value any)
}

type Group interface {
	ID() string
	Put(ch Channel) (err error)
	Del(ch Channel) bool
	Push(p packet.IPacket)
	Close()
}

type receiver interface {
	Receive(p packet.IPacket)
}
