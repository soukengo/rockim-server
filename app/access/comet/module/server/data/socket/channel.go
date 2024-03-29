package socket

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/soukengo/gopkg/component/transport/socket"
	"github.com/soukengo/gopkg/component/transport/socket/packet"
	"github.com/soukengo/gopkg/log"
	v1 "rockimserver/apis/rockim/api/client/v1/protocol/socket"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/access/comet/module/protocol"
	"rockimserver/app/access/comet/module/server/biz/options"
)

type ChannelManager struct {
	log    log.Logger
	server socket.Server

	controls map[enums.Comet_ControlType]ControlFunc
}

func NewChannelManager(log log.Logger, server socket.Server) *ChannelManager {
	ins := &ChannelManager{
		log:      log,
		server:   server,
		controls: map[enums.Comet_ControlType]ControlFunc{},
	}

	ins.controls[enums.Comet_ROOM_JOIN] = wrapControl(ins.RoomJoin)
	ins.controls[enums.Comet_ROOM_QUIT] = wrapControl(ins.RoomQuit)
	ins.controls[enums.Comet_KICK_OFF] = wrapControl(ins.KickOff)
	return ins
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

func (m *ChannelManager) PushRoom(ctx context.Context, opts *options.PushRoomOptions) (err error) {
	p := newPushPacket(opts.Operation, opts.Body)
	m.server.PushRoom(protocol.EncodeRoomId(opts.Room), p)
	return
}

func newPushPacket(operation enums.Comet_PushOperation, data []byte) packet.IPacket {
	header := &v1.PushPacketHeader{
		Operation: operation,
	}
	body := &v1.PushPacketBody{
		Body: data,
	}
	headerBytes, _ := encode(header)
	bodyBytes, _ := encode(body)
	return protocol.NewPacket(uint8(enums.Network_PUSH), headerBytes, bodyBytes)
}

func encode(in proto.Message) ([]byte, error) {
	return proto.Marshal(in)
}
