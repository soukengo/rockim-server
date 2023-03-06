package biz

import (
	"context"
	"github.com/antlabs/timer"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/access/comet/conf"
	"rockimserver/app/access/comet/module/client/biz/consts"
	"rockimserver/app/access/comet/module/client/biz/options"
	"rockimserver/pkg/component/server/socket"
	"rockimserver/pkg/errors"
	"rockimserver/pkg/log"
	"time"
)

var (
	ErrRequestInvalid = errors.BadRequest(reasons.ErrorReason_UN_SPECIFIED.String(), "无效的请求")
)

type OnlineRepo interface {
	Connect(ctx context.Context, opts *options.ClientAuthOptions) (string, error)
	HeartBeat(ctx context.Context, opts *options.ClientHeartBeatOptions) error
	DisConnect(ctx context.Context, opts *options.ClientDisConnectOptions) error
}

type ClientUseCase struct {
	cfg    *conf.Config
	server socket.Server
	repo   OnlineRepo
	timer  timer.Timer
}

func NewClientUseCase(cfg *conf.Config, server socket.Server, repo OnlineRepo) *ClientUseCase {
	tm := timer.NewTimer()
	tm.Run()
	return &ClientUseCase{cfg: cfg, server: server, repo: repo, timer: tm}
}

func (uc *ClientUseCase) Connect(ctx context.Context) (err error) {
	ch, ok := socket.FromContext(ctx)
	if !ok {
		err = ErrRequestInvalid
		return
	}
	channelId := ch.Id()
	// 设置定时器，超时未授权就断开连接
	uc.timer.AfterFunc(uc.cfg.Protocol.HandshakeTimeout, func() {
		ch, ok = uc.server.Channel(channelId)
		if !ok {
			return
		}
		if !ch.Authenticated() {
			log.WithContext(ctx).Warnf("channel handshake timed out: %v", channelId)
			_ = ch.Close()
		}
	})
	return
}

func (uc *ClientUseCase) DisConnect(ctx context.Context) (err error) {
	ch, ok := socket.FromContext(ctx)
	if !ok {
		err = ErrRequestInvalid
		return
	}
	opts := &options.ClientDisConnectOptions{
		ProductId: ch.StringAttr(consts.ChannelAttrProductId),
		Uid:       ch.StringAttr(consts.ChannelAttrUid),
		ServerId:  uc.cfg.Global.ID,
		ChannelId: ch.Id(),
	}
	return uc.repo.DisConnect(ctx, opts)
}

func (uc *ClientUseCase) Auth(ctx context.Context, opts *options.ClientAuthOptions) (err error) {
	ch, ok := socket.FromContext(ctx)
	if !ok {
		err = ErrRequestInvalid
		return
	}
	opts.ServerId = uc.cfg.Global.ID
	opts.ChannelId = ch.Id()
	uid, err := uc.repo.Connect(ctx, opts)
	if err != nil {
		return
	}
	ch.SetAttr(consts.ChannelAttrProductId, opts.ProductId)
	ch.SetAttr(consts.ChannelAttrUid, uid)
	return
}

func (uc *ClientUseCase) HeartBeat(ctx context.Context) (err error) {
	ch, ok := socket.FromContext(ctx)
	if !ok {
		err = ErrRequestInvalid
		return
	}
	lastTime := ch.Int64Attr(consts.ChannelAttrLastHeartBeatTime)
	nowts := time.Now().UnixMilli()
	// 心跳降频
	if nowts-lastTime < uc.cfg.Protocol.HeartbeatInterval.Milliseconds() {
		log.WithContext(ctx).Debugf("Heartbeat reduce frequency,channelId: %v", ch.Id())
		return
	}
	opts := &options.ClientHeartBeatOptions{
		ProductId: ch.StringAttr(consts.ChannelAttrProductId),
		Uid:       ch.StringAttr(consts.ChannelAttrUid),
		ServerId:  uc.cfg.Global.ID,
		ChannelId: ch.Id(),
	}
	err = uc.repo.HeartBeat(ctx, opts)
	if err != nil {
		return
	}
	ch.SetAttr(consts.ChannelAttrLastHeartBeatTime, nowts)
	return
}
