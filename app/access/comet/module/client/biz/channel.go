package biz

import (
	"context"
	"github.com/antlabs/timer"
	"github.com/soukengo/gopkg/component/server/socket"
	"github.com/soukengo/gopkg/errors"
	"github.com/soukengo/gopkg/log"
	"github.com/soukengo/gopkg/util/runtimes"
	"rockimserver/apis/rockim/service/comet/v1/types"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/access/comet/conf"
	"rockimserver/app/access/comet/module/client/biz/options"
	"rockimserver/app/access/comet/module/protocol"
	"time"
)

type ChannelRepo interface {
	Add(ctx context.Context, opts *options.ChannelAddOptions) error
	Refresh(ctx context.Context, opts *options.ChannelRefreshOptions) error
	Delete(ctx context.Context, opts *options.ChannelDeleteOptions) error
}

type TokenRepo interface {
	CheckAuth(ctx context.Context, opts *options.ChannelAuthOptions) (string, error)
}

type RoomRepo interface {
	List(ctx context.Context, opts *options.ListRoomOptions) ([]*types.Room, error)
}

var (
	ErrUnSpecified = errors.BadRequest(reasons.ErrorReason_UN_SPECIFIED.String(), "未知错误")
)

type ChannelUseCase struct {
	cfg       *conf.Config
	timer     timer.Timer
	tokenRepo TokenRepo
	repo      ChannelRepo
	roomRepo  RoomRepo
}

func NewChannelUseCase(cfg *conf.Config, tokenRepo TokenRepo, repo ChannelRepo, roomRepo RoomRepo) *ChannelUseCase {
	tm := timer.NewTimer()
	runtimes.Async(tm.Run)
	return &ChannelUseCase{cfg: cfg, timer: tm, tokenRepo: tokenRepo, repo: repo, roomRepo: roomRepo}
}

func (uc *ChannelUseCase) Connect(ctx context.Context) (err error) {
	chCtx, ok := socket.FromContext(ctx)
	if !ok {
		err = ErrUnSpecified
		return
	}
	ch := chCtx.Channel()
	session, err := protocol.SessionFromChannel(ch)
	if err != nil {
		return
	}
	uc.setAuthCheckTimer(ch, session)
	return
}

func (uc *ChannelUseCase) DisConnect(ctx context.Context) (err error) {
	chCtx, ok := socket.FromContext(ctx)
	if !ok {
		err = ErrUnSpecified
		return
	}
	ch := chCtx.Channel()
	session, err := protocol.SessionFromChannel(ch)
	if err != nil {
		return
	}
	opts := &options.ChannelDeleteOptions{
		ProductId: session.ProductId,
		Uid:       session.Uid,
		ServerId:  uc.cfg.Global.ID,
		ChannelId: ch.Id(),
	}
	uc.clearAuthCheckTimer(ch, session)
	uc.clearHeartBeatCheckTimer(ch, session)
	return uc.repo.Delete(ctx, opts)
}

func (uc *ChannelUseCase) Auth(ctx context.Context, in *options.ChannelAuthOptions) (err error) {
	chCtx, ok := socket.FromContext(ctx)
	if !ok {
		err = ErrUnSpecified
		return
	}
	ch := chCtx.Channel()

	uid, err := uc.tokenRepo.CheckAuth(ctx, in)
	if err != nil {
		return
	}
	opts := &options.ChannelAddOptions{
		ProductId: in.ProductId,
		ServerId:  uc.cfg.Global.ID,
		ChannelId: ch.Id(),
		Uid:       uid,
	}
	err = uc.repo.Add(ctx, opts)
	if err != nil {
		return
	}
	session, err := protocol.SessionFromChannel(ch)
	if err != nil {
		return
	}
	session.ProductId = opts.ProductId
	session.Uid = uid
	// 标记为已验证
	session.MarkAuthenticated()
	// 产生随机的心跳间隔
	session.ServerHeartBeatInterval = uc.cfg.Protocol.RandomServerHeartbeatInterval().Milliseconds()
	// 开启心跳检测
	uc.setHeartBeatCheckTimer(ch, session)

	// 异步加载房间列表
	runtimes.Async(func() {
		uc.loadRoom(chCtx.Server(), ch, session)
	})

	return
}

func (uc *ChannelUseCase) HeartBeat(ctx context.Context) (err error) {
	chCtx, ok := socket.FromContext(ctx)
	if !ok {
		err = ErrUnSpecified
		return
	}
	ch := chCtx.Channel()
	session, err := protocol.SessionFromChannel(ch)
	if err != nil {
		return
	}
	// 更新客户端心跳时间
	nowts := time.Now().UnixMilli()
	session.LastHeartBeatTime = nowts
	// 获取上一次向后端发起心跳的时间
	lastTime := session.LastServerHeartBeatTime
	// 获取心跳间隔
	heartBeatInterval := session.ServerHeartBeatInterval
	if heartBeatInterval <= 0 {
		heartBeatInterval = uc.cfg.Protocol.MinServerHeartbeatInterval.Milliseconds()
	}
	// 心跳降频
	if nowts-lastTime < heartBeatInterval {
		log.WithContext(ctx).Debugf("Heartbeat reduce frequency,channelId: %v", ch.Id())
		return
	}
	opts := &options.ChannelRefreshOptions{
		ProductId: session.ProductId,
		Uid:       session.Uid,
		ServerId:  uc.cfg.Global.ID,
		ChannelId: ch.Id(),
	}
	err = uc.repo.Refresh(ctx, opts)
	if err != nil {
		return
	}
	// 更新向后端发起心跳的时间
	session.LastServerHeartBeatTime = nowts
	return
}

func (uc *ChannelUseCase) setAuthCheckTimer(ch socket.Channel, session *protocol.Session) {
	// 设置定时器，超时未授权就断开连接
	node := uc.timer.AfterFunc(uc.cfg.Protocol.HandshakeTimeout, func() {
		if !session.Authenticated() {
			log.Warnf("channel handshake timed out: %v", ch.Id())
			_ = ch.Close()
		}
	})
	session.AuthCheckTimer = node
}
func (uc *ChannelUseCase) clearAuthCheckTimer(ch socket.Channel, session *protocol.Session) {
	if session.AuthCheckTimer != nil {
		session.AuthCheckTimer.Stop()
		session.AuthCheckTimer = nil
	}
}

func (uc *ChannelUseCase) setHeartBeatCheckTimer(ch socket.Channel, session *protocol.Session) {
	// 移除上一次的timer
	uc.clearHeartBeatCheckTimer(ch, session)
	// 设置定时器，超时未心跳就断开连接
	node := uc.timer.AfterFunc(uc.cfg.Protocol.MaxClientHeartbeatInterval, func() {
		lastTime := session.LastHeartBeatTime
		nowts := time.Now().UnixMilli()
		if nowts-lastTime > uc.cfg.Protocol.MaxClientHeartbeatInterval.Milliseconds() {
			log.Warnf("channel heartbeat timed out: %v", ch.Id())
			_ = ch.Close()
		}
	})
	session.HeartBeatCheckTimer = node
}

func (uc *ChannelUseCase) clearHeartBeatCheckTimer(ch socket.Channel, session *protocol.Session) {
	if session.HeartBeatCheckTimer != nil {
		session.HeartBeatCheckTimer.Stop()
		session.HeartBeatCheckTimer = nil
	}
}

// loadRoom 加载房间列表
func (uc *ChannelUseCase) loadRoom(srv socket.Server, ch socket.Channel, session *protocol.Session) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	rooms, err := uc.roomRepo.List(ctx, &options.ListRoomOptions{
		ProductId: session.ProductId,
		Uid:       session.Uid,
	})
	if err != nil {
		log.Errorf("ListRoom error: %v", err)
		return
	}
	for _, room := range rooms {
		_ = srv.JoinRoom(protocol.EncodeRoomId(room), ch)
	}
}
