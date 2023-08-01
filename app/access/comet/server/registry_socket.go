package server

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/soukengo/gopkg/component/trace"
	"github.com/soukengo/gopkg/component/transport/socket"
	"github.com/soukengo/gopkg/component/transport/socket/packet"
	"github.com/soukengo/gopkg/errors"
	"github.com/soukengo/gopkg/log"
	v1 "rockimserver/apis/rockim/api/client/v1/protocol/socket"
	"rockimserver/apis/rockim/shared"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/access/comet/conf"
	"rockimserver/app/access/comet/module/client/service"
	"rockimserver/app/access/comet/module/protocol"
)

var (
	ErrOperationNotSupported = errors.BadRequest(reasons.Socket_OPERATION_NOT_SUPPORTED.String(), "不支持的操作类型")
	ErrPacketInvalid         = errors.BadRequest(reasons.Socket_PACKET_INVALID.String(), "无效的数据包")
)

type HandleFunc func(context.Context, socket.Channel, []byte) (any, error)

type SocketRegistry struct {
	actions map[v1.Operation]HandleFunc

	cfg        *conf.Protocol
	channelSrv *service.ChannelService
}

func NewSocketRegistry(cfg *conf.Protocol, channelSrv *service.ChannelService) *SocketRegistry {
	ins := &SocketRegistry{
		actions:    map[v1.Operation]HandleFunc{},
		cfg:        cfg,
		channelSrv: channelSrv,
	}
	return ins
}

func (h *SocketRegistry) RegisterSocket(srv socket.Server) {
	srv.SetHandler(h)
	h.actions[v1.Operation_AUTH] = wrapAction[v1.AuthRequest, v1.AuthResponse](h.channelSrv.Auth)
	h.actions[v1.Operation_HEARTBEAT] = wrapAction[v1.HeartBeatRequest, v1.HeartBeatResponse](h.channelSrv.HeartBeat)
}

func (h *SocketRegistry) OnCreated(ctx *socket.Context) {
	ch := ctx.Channel()
	channelId := ch.Id()
	log.WithContext(ctx).Debugf("channel created: %v", channelId)
	ch.SetSession(&protocol.Session{})
	_ = h.channelSrv.Connect(ctx)
}

func (h *SocketRegistry) OnClosed(ctx *socket.Context) {
	_ = h.channelSrv.DisConnect(ctx)
}

func (h *SocketRegistry) OnReceived(ctx *socket.Context, p packet.IPacket) {
	ch := ctx.Channel()
	pkt, ok := p.(*protocol.Packet)
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
	h.handle(ctx, ch, header, body)
}

func (h *SocketRegistry) handle(ctx context.Context, channel socket.Channel, header *v1.RequestPacketHeader, body *v1.RequestPacketBody) {
	var (
		data any
		err  error
	)
	operation := header.Operation
	handler, ok := h.actions[operation]
	respHeader := &v1.ResponsePacketHeader{
		Operation: operation,
		RequestId: header.RequestId,
		TraceId:   trace.TraceID(ctx),
		Success:   true,
	}
	if !ok {
		err = ErrOperationNotSupported
		h.render(ctx, channel, respHeader, data, err)
		return
	}
	data, err = handler(ctx, channel, body.Body)
	h.render(ctx, channel, respHeader, data, err)
}

func (h *SocketRegistry) render(ctx context.Context, ch socket.Channel, header *v1.ResponsePacketHeader, data any, err error) {
	if err != nil {
		var e = errors.FromError(err)
		data = &shared.Error{
			Code:     e.Code(),
			Reason:   e.Reason(),
			Message:  e.Message(),
			Metadata: e.Metadata(),
		}
		header.Success = false
	}
	headerBytes, err1 := encode(header)
	if err1 != nil {
		log.WithContext(ctx).Errorf("encode header error: %v", err1)
	}
	dataBytes, err1 := encode(data)
	if err1 != nil {
		log.WithContext(ctx).Errorf("encode data error: %v", err1)
	}
	bodyBytes, err1 := encode(&v1.ResponsePacketBody{Body: dataBytes})
	if err1 != nil {
		log.WithContext(ctx).Errorf("encode body error: %v", err1)
	}
	p := protocol.NewPacket(uint8(enums.Network_RESPONSE), headerBytes, bodyBytes)
	err = ch.Send(p)
	if err != nil {
		log.WithContext(ctx).Errorf("channel.Send channelId: %v error: %v", ch.Id(), err)
	}
}

type validator interface {
	Validate() error
}

func wrapAction[Request any, Response any](fn func(context.Context, *Request) (*Response, error)) HandleFunc {
	return func(ctx context.Context, channel socket.Channel, in []byte) (out any, err error) {
		var req = new(Request)
		err = decode(in, req)
		if err != nil {
			return
		}
		// validate
		var r any = req
		if v, ok := r.(validator); ok {
			if err = v.Validate(); err != nil {
				return nil, errors.BadRequest("VALIDATOR", err.Error())
			}
		}
		return fn(ctx, req)
	}
}

func decode(data []byte, v any) error {
	in, ok := v.(proto.Message)
	if !ok {
		return ErrPacketInvalid
	}
	return proto.Unmarshal(data, in)
}

func encode(v any) ([]byte, error) {
	in, ok := v.(proto.Message)
	if !ok {
		return nil, ErrPacketInvalid
	}
	return proto.Marshal(in)
}
