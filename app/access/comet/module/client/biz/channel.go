package biz

import (
	"context"
	"github.com/antlabs/timer"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/access/comet/conf"
	"rockimserver/app/access/comet/module/client/biz/consts"
	"rockimserver/app/access/comet/module/client/biz/options"
	"rockimserver/pkg/component/server/socket"
	"rockimserver/pkg/log"
	"time"
)

type OnlineRepo interface {
	Add(ctx context.Context, opts *options.OnlineAddOptions) (string, error)
	Refresh(ctx context.Context, opts *options.OnlineRefreshOptions) error
	Delete(ctx context.Context, opts *options.OnlineDeleteOptions) error
}

type ChannelUseCase struct {
	cfg    *conf.Config
	server socket.Server
	repo   OnlineRepo
	timer  timer.Timer
}

func NewChannelUseCase(cfg *conf.Config, server socket.Server, repo OnlineRepo) *ChannelUseCase {
	tm := timer.NewTimer()
	tm.Run()
	return &ChannelUseCase{cfg: cfg, server: server, repo: repo, timer: tm}
}

func (uc *ChannelUseCase) Connect(ctx context.Context) (err error) {
	ch, ok := socket.FromContext(ctx)
	if !ok {
		err = reasons.ErrUnSpecified
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

func (uc *ChannelUseCase) DisConnect(ctx context.Context) (err error) {
	ch, ok := socket.FromContext(ctx)
	if !ok {
		err = reasons.ErrUnSpecified
		return
	}
	opts := &options.OnlineDeleteOptions{
		ProductId: ch.StringAttr(consts.ChannelAttrProductId),
		Uid:       ch.StringAttr(consts.ChannelAttrUid),
		ServerId:  uc.cfg.Global.ID,
		ChannelId: ch.Id(),
	}
	return uc.repo.Delete(ctx, opts)
}

func (uc *ChannelUseCase) Auth(ctx context.Context, in *options.ChannelAuthOptions) (err error) {
	ch, ok := socket.FromContext(ctx)
	if !ok {
		err = reasons.ErrUnSpecified
		return
	}
	opts := &options.OnlineAddOptions{
		ProductId: in.ProductId,
		ServerId:  uc.cfg.Global.ID,
		ChannelId: ch.Id(),
		Token:     in.Token,
	}
	uid, err := uc.repo.Add(ctx, opts)
	if err != nil {
		return
	}
	ch.SetAttr(consts.ChannelAttrProductId, opts.ProductId)
	ch.SetAttr(consts.ChannelAttrUid, uid)
	return
}

func (uc *ChannelUseCase) HeartBeat(ctx context.Context) (err error) {
	ch, ok := socket.FromContext(ctx)
	if !ok {
		err = reasons.ErrUnSpecified
		return
	}
	lastTime := ch.Int64Attr(consts.ChannelAttrLastHeartBeatTime)
	nowts := time.Now().UnixMilli()
	// 心跳降频
	if nowts-lastTime < uc.cfg.Protocol.HeartbeatInterval.Milliseconds() {
		log.WithContext(ctx).Debugf("Heartbeat reduce frequency,channelId: %v", ch.Id())
		return
	}
	opts := &options.OnlineRefreshOptions{
		ProductId: ch.StringAttr(consts.ChannelAttrProductId),
		Uid:       ch.StringAttr(consts.ChannelAttrUid),
		ServerId:  uc.cfg.Global.ID,
		ChannelId: ch.Id(),
	}
	err = uc.repo.Refresh(ctx, opts)
	if err != nil {
		return
	}
	ch.SetAttr(consts.ChannelAttrLastHeartBeatTime, nowts)
	return
}
