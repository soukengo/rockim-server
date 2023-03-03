package network

import (
	"net"
	"rockimserver/pkg/component/server/socket/packet"
)

type Server interface {
	Id() string
	Start() error
	Close() error
	SetHandler(listener Handler)
}

type Connection interface {
	Id() string
	Send(packet.IPacket) error
	RemoteAddr() net.Addr
	Close() error
}

type Handler interface {
	OnConnect(Connection)
	OnDisConnect(Connection)
	OnReceived(Connection, packet.IPacket)
}

var (
	DefaultHandler = &handler{}
)

// handler default server handler
type handler struct {
}

func (h handler) OnConnect(conn Connection) {
}

func (h handler) OnDisConnect(conn Connection) {
}

func (h handler) OnReceived(conn Connection, packet packet.IPacket) {
}
