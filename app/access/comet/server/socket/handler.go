package socket

import (
	"context"
	"github.com/antlabs/timer"
	"github.com/golang/protobuf/proto"
	v1 "rockimserver/apis/rockim/api/client/socket/v1"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/access/comet/module/client/service"
	"rockimserver/pkg/component/server/socket"
	"rockimserver/pkg/component/server/socket/packet"
	"rockimserver/pkg/component/trace"
	"rockimserver/pkg/errors"
	"rockimserver/pkg/log"
	"time"
)

const (
	handshakeTimeout = time.Second * 10
)

var (
	ErrApiNotFound    = errors.NotFound("", "")
	ErrInvalidRequest = errors.BadRequest("", "")
)

type HandleFunc func(context.Context, socket.Channel, []byte) (any, error)

type ClientHandler struct {
	srv       socket.Server
	timer     timer.Timer
	actions   map[v1.Operation]HandleFunc
	socketSrv *service.SocketService
}

func NewClientHandler() *ClientHandler {
	t := timer.NewTimer()
	t.Run()
	ins := &ClientHandler{timer: t}
	ins.register()
	return ins
}

func (h *ClientHandler) register() {
	h.actions[v1.Operation_AUTH] = wrapAction[v1.AuthRequest, v1.AuthResponse](h.socketSrv.Auth)
	h.actions[v1.Operation_HEARTBEAT] = wrapAction[v1.HeartBeatRequest, v1.HeartBeatResponse](h.socketSrv.HeartBeat)
}

func (h *ClientHandler) OnCreated(channel socket.Channel) {
	channelId := channel.Id()
	log.Debugf("channel created: %v", channelId)
	h.timer.AfterFunc(handshakeTimeout, func() {
		ch, ok := h.srv.Channel(channelId)
		if !ok {
			return
		}
		if !ch.Authenticated() {
			log.Warnf("channel handshake timed out: %v", channelId)
			_ = ch.Close()
		}
	})
}

func (h *ClientHandler) OnClosed(channel socket.Channel) {
	ctx := context.Background()
	h.socketSrv.DisConnect(ctx, channel.Id())
}

func (h *ClientHandler) OnReceived(channel socket.Channel, p packet.IPacket) {
	ctx := context.Background()
	pkt, ok := p.(*v1.Packet)
	if !ok {
		log.Errorf("OnReceived packet not invalid: %v", p)
		return
	}
	if int32(pkt.Typ) != int32(enums.Network_REQUEST) {
		log.Errorf("OnReceived packet not invalid: %v", p)
		return
	}
	var err error
	header := &v1.RequestPacketHeader{}
	err = decode(pkt.Header, header)
	if err != nil {
		log.Errorf("OnReceived packet header not invalid: %v", p)
		return
	}
	body := &v1.RequestPacketBody{}
	err = decode(pkt.Body, body)
	if err != nil {
		log.Errorf("OnReceived packet body not invalid: %v", p)
		return
	}
	h.handle(ctx, channel, header, body)
}

func (h *ClientHandler) handle(ctx context.Context, channel socket.Channel, header *v1.RequestPacketHeader, body *v1.RequestPacketBody) {
	var (
		data any
		err  error
	)
	operation := header.Operation
	handler, ok := h.actions[operation]
	respHeader := &v1.ResponsePacketHeader{
		RequestId: header.RequestId,
		TraceId:   trace.TraceID(ctx),
		Success:   true,
	}
	if !ok {
		err = ErrApiNotFound
		h.render(ctx, channel, respHeader, data, err)
		return
	}
	data, err = handler(ctx, channel, body.Body)
	h.render(ctx, channel, respHeader, data, err)
}

func (h *ClientHandler) render(ctx context.Context, ch socket.Channel, header *v1.ResponsePacketHeader, data any, err error) {
	if err != nil {
		data = err
		header.Success = false
	}
	headerBytes, err1 := encode(header)
	if err1 != nil {
		log.WithContext(ctx).Errorf("encode header error: %v", err1)
	}
	bodyBytes, err1 := encode(data)
	if err1 != nil {
		log.WithContext(ctx).Errorf("encode body error: %v", err1)
	}
	p := v1.NewPacket(uint8(enums.Network_RESPONSE), headerBytes, bodyBytes)
	err = ch.Send(p)
	if err != nil {
		log.WithContext(ctx).Errorf("channel.Send channelId: %v error: %v", ch.Id(), err)
	}
}

func wrapAction[Request any, Response any](fn func(context.Context, *Request) (*Response, error)) HandleFunc {
	return func(ctx context.Context, channel socket.Channel, in []byte) (out any, err error) {
		var req = new(Request)
		err = decode(in, req)
		if err != nil {
			return
		}
		return fn(ctx, req)
	}
}

func decode(data []byte, v any) error {
	in, ok := v.(proto.Message)
	if !ok {
		return ErrInvalidRequest
	}
	return proto.Unmarshal(data, in)
}

func encode(v any) ([]byte, error) {
	in, ok := v.(proto.Message)
	if !ok {
		return nil, ErrInvalidRequest
	}
	return proto.Marshal(in)
}
