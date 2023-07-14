package socket

import (
	"context"
	"github.com/soukengo/gopkg/component/server/socket"
	"github.com/soukengo/gopkg/errors"
	"github.com/soukengo/gopkg/log"
	"google.golang.org/protobuf/proto"
	"rockimserver/apis/rockim/service/comet/v1/types"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/access/comet/module/protocol"
	"rockimserver/app/access/comet/module/server/biz/options"
)

var (
	ErrControlNotSupported = errors.BadRequest(reasons.Comet_CONTROL_NOT_SUPPORTED.String(), "无效的数据")
	ErrControlDataInvalid  = errors.BadRequest(reasons.Comet_CONTROL_DATA_INVALID.String(), "无效的数据")
)

type ControlFunc func(context.Context, socket.Channel, []byte) error

func (m *ChannelManager) Control(ctx context.Context, opts *options.ControlOptions) error {
	fn, ok := m.controls[opts.ControlType]
	if !ok {
		return ErrControlNotSupported
	}
	for _, cid := range opts.ChannelIds {
		ch, exists := m.server.Channel(cid)
		if !exists {
			log.WithContext(ctx).Warn("Control Channel not found", log.Pairs{"channelId": cid})
			continue
		}
		err1 := fn(ctx, ch, opts.Body)
		if err1 != nil {
			log.WithContext(ctx).Errorf("Control error", log.Pairs{"channelId": cid, "controlType": opts.ControlType})
			continue
		}
	}
	return nil
}

func (m *ChannelManager) RoomJoin(ctx context.Context, ch socket.Channel, in *types.RoomJoinControl) (err error) {
	for _, room := range in.Rooms {
		roomId := protocol.EncodeRoomId(room)
		err1 := m.server.JoinRoom(roomId, ch)
		if err1 != nil {
			log.WithContext(ctx).Errorf("JoinRoom error", log.Pairs{"channelId": ch.Id(), "roomId": roomId})
			continue
		}
	}
	return
}

func (m *ChannelManager) RoomQuit(ctx context.Context, ch socket.Channel, in *types.RoomQuitControl) (err error) {
	for _, room := range in.Rooms {
		roomId := protocol.EncodeRoomId(room)
		err1 := m.server.QuitRoom(roomId, ch)
		if err1 != nil {
			log.WithContext(ctx).Errorf("QuitRoom error", log.Pairs{"channelId": ch.Id(), "roomId": roomId})
			continue
		}
	}
	return
}

func (m *ChannelManager) KickOff(ctx context.Context, ch socket.Channel, in *types.KickOffControl) (err error) {

	return
}

func wrapControl[T any](fn func(context.Context, socket.Channel, *T) error) func(ctx context.Context, ch socket.Channel, data []byte) (err error) {
	return func(ctx context.Context, ch socket.Channel, data []byte) (err error) {
		var req = new(T)
		err = decode(data, req)
		if err != nil {
			return
		}
		err = fn(ctx, ch, req)
		return
	}
}

func decode(data []byte, v any) error {
	in, ok := v.(proto.Message)
	if !ok {
		return ErrControlDataInvalid
	}
	return proto.Unmarshal(data, in)
}
