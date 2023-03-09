package socket

import (
	"context"
	log "github.com/golang/glog"
	"rockimserver/pkg/component/server/socket/network"
	"rockimserver/pkg/component/server/socket/packet"
)

func (m *server) OnConnect(conn network.Connection) {
	channelId := conn.Id()
	log.Infof("OnConnect: %v", channelId)
	ch := newChannel(conn, m.opt.SendQueueSize)
	m.Bucket(conn.Id()).PutChannel(ch)
	ctx := NewContext(context.Background(), ch)
	m.handler.OnCreated(ctx)
}

func (m *server) OnDisConnect(conn network.Connection) {
	channelId := conn.Id()
	log.Infof("OnDisConnect: %v", channelId)
	ch, ok := m.Channel(conn.Id())
	if !ok {
		return
	}
	m.Bucket(channelId).DelChannel(ch)
	ctx := NewContext(context.Background(), ch)
	m.handler.OnClosed(ctx)
}

func (m *server) OnReceived(conn network.Connection, p packet.IPacket) {
	log.Infof("OnReceived")
	ch, ok := m.Channel(conn.Id())
	if !ok {
		_ = conn.Close()
		return
	}
	ctx := NewContext(context.Background(), ch)
	m.handler.OnReceived(ctx, p)
}
