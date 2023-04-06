package socket

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/soukengo/gopkg/component/server/socket"
	"github.com/soukengo/gopkg/component/server/socket/packet"
	"github.com/soukengo/gopkg/log"
	v1 "rockimserver/apis/rockim/api/client/v1/socket"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/access/comet/module/server/biz/options"
)

type ChannelManager struct {
	log    log.Logger
	server socket.Server
}

func NewChannelManager(log log.Logger, server socket.Server) *ChannelManager {
	return &ChannelManager{log: log, server: server}
}

func (m *ChannelManager) Push(ctx context.Context, opts *options.PushOptions) error {
	p := newPushPacket(opts.Operation, opts.Body)
	for _, id := range opts.ChannelIds {
		ch, ok := m.server.Channel(id)
		if !ok {
			continue
		}
		_ = ch.Send(p)
	}
	return nil
}

func (m *ChannelManager) PushGroup(ctx context.Context, opts *options.PushGroupOptions) (err error) {
	p := newPushPacket(opts.Operation, opts.Body)
	m.server.PushGroup(opts.GroupId, p)
	return
}

func newPushPacket(operation enums.Network_PushOperation, data []byte) packet.IPacket {
	header := &v1.PushPacketHeader{
		Operation: operation,
	}
	body := &v1.PushPacketBody{
		Body: data,
	}
	headerBytes, _ := encode(header)
	bodyBytes, _ := encode(body)
	return v1.NewPacket(uint8(enums.Network_PUSH), headerBytes, bodyBytes)
}

func encode(in proto.Message) ([]byte, error) {
	return proto.Marshal(in)
}
