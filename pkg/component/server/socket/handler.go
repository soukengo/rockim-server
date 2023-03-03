package socket

import (
	"rockimserver/pkg/component/server/socket/network"
	"rockimserver/pkg/component/server/socket/packet"
)

func (m *server) OnConnect(conn network.Connection) {
	var ch Channel
	ch = newChannel(conn, m.opt.RecvQueueSize, m.opt.SendQueueSize, func(p packet.IPacket) {
		m.handler.OnReceived(ch, p)
	})
	m.Bucket(conn.Id()).PutChannel(ch)
	m.handler.OnCreated(ch)
}

func (m *server) OnDisConnect(conn network.Connection) {
	channelId := conn.Id()
	ch, ok := m.Channel(conn.Id())
	if !ok {
		return
	}
	m.Bucket(channelId).DelChannel(ch)
	m.handler.OnClosed(ch)
}

func (m *server) OnReceived(conn network.Connection, p packet.IPacket) {
	ch, ok := m.Channel(conn.Id())
	if !ok {
		_ = conn.Close()
		return
	}
	if r, ok := ch.(receiver); ok {
		r.Receive(p)
	} else {
		m.handler.OnReceived(ch, p)
	}
}
