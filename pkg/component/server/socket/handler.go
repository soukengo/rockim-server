package socket

import (
	"context"
	"rockimserver/pkg/component/server/socket/network"
	"rockimserver/pkg/component/server/socket/packet"
)

func (m *server) OnConnect(conn network.Connection) {
	ch := newChannel(conn, m.opt.RecvQueueSize, m.opt.SendQueueSize)
	m.Bucket(conn.Id()).PutChannel(ch)
	ctx := NewContext(context.Background(), ch)
	m.handler.OnCreated(ctx)
}

func (m *server) OnDisConnect(conn network.Connection) {
	channelId := conn.Id()
	ch, ok := m.Channel(conn.Id())
	if !ok {
		return
	}
	m.Bucket(channelId).DelChannel(ch)
	ctx := NewContext(context.Background(), ch)
	m.handler.OnClosed(ctx)
}

func (m *server) OnReceived(conn network.Connection, p packet.IPacket) {
	ch, ok := m.Channel(conn.Id())
	if !ok {
		_ = conn.Close()
		return
	}
	ctx := NewContext(context.Background(), ch)
	m.handler.OnReceived(ctx, p)
}
