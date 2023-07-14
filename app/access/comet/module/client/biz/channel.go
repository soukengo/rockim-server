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
	"rockimserver/app/access/comet/module/client/biz/consts"
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
	uc.setAuthCheckTimer(ch)
	return
}

func (uc *ChannelUseCase) DisConnect(ctx context.Context) (err error) {
	chCtx, ok := socket.FromContext(ctx)
	if !ok {
		err = ErrUnSpecified
		return
	}
	ch := chCtx.Channel()
	opts := &options.ChannelDeleteOptions{
		ProductId: ch.StringAttr(consts.ChannelAttrProductId),
		Uid:       ch.StringAttr(consts.ChannelAttrUid),
		ServerId:  uc.cfg.Global.ID,
		ChannelId: ch.Id(),
	}
	uc.clearAuthCheckTimer(ch)
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
	//  设置自定义属性
	ch.SetAttr(consts.ChannelAttrProductId, opts.ProductId)
	ch.SetAttr(consts.ChannelAttrUid, uid)
	// 标记为已验证
	ch.MarkAuthenticated()
	// 产生随机的心跳间隔
	ch.SetAttr(consts.ChannelAttrServerHeartBeatInterval, uc.cfg.Protocol.RandomServerHeartbeatInterval().Milliseconds())
	// 异步加载房间列表
	runtimes.Async(func() {
		uc.loadRoom(chCtx.Server(), ch)
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
	// 更新客户端心跳时间
	nowts := time.Now().UnixMilli()
	ch.SetAttr(consts.ChannelAttrLastHeartBeatTime, nowts)
	// 获取上一次向后端发起心跳的时间
	lastTime := ch.Int64Attr(consts.ChannelAttrLastServerHeartBeatTime)
	// 获取心跳间隔
	heartBeatInterval := ch.Int64Attr(consts.ChannelAttrServerHeartBeatInterval)
	if heartBeatInterval <= 0 {
		heartBeatInterval = uc.cfg.Protocol.MinServerHeartbeatInterval.Milliseconds()
	}
	// 心跳降频
	if nowts-lastTime < heartBeatInterval {
		log.WithContext(ctx).Debugf("Heartbeat reduce frequency,channelId: %v", ch.Id())
		return
	}
	opts := &options.ChannelRefreshOptions{
		ProductId: ch.StringAttr(consts.ChannelAttrProductId),
		Uid:       ch.StringAttr(consts.ChannelAttrUid),
		ServerId:  uc.cfg.Global.ID,
		ChannelId: ch.Id(),
	}
	err = uc.repo.Refresh(ctx, opts)
	if err != nil {
		return
	}
	// 更新向后端发起心跳的时间
	ch.SetAttr(consts.ChannelAttrLastServerHeartBeatTime, nowts)
	return
}

func (uc *ChannelUseCase) setAuthCheckTimer(ch socket.Channel) {
	// 设置定时器，超时未授权就断开连接
	node := uc.timer.AfterFunc(uc.cfg.Protocol.HandshakeTimeout, func() {
		if !ch.Authenticated() {
			log.Warnf("channel handshake timed out: %v", ch.Id())
			_ = ch.Close()
		}
	})
	ch.SetAttr(consts.ChannelAttrTimerAuthCheck, node)
}
func (uc *ChannelUseCase) clearAuthCheckTimer(ch socket.Channel) {
	value, ok := ch.Attr(consts.ChannelAttrTimerAuthCheck)
	if !ok {
		return
	}
	node, ok := value.(timer.TimeNoder)
	if !ok {
		return
	}
	node.Stop()
	ch.DelAttr(consts.ChannelAttrTimerAuthCheck)
}
func (uc *ChannelUseCase) setHeartBeatCheckTimer(ch socket.Channel) {
	var node timer.TimeNoder
	// 移除上一次的timer
	value, ok := ch.Attr(consts.ChannelAttrTimerHeartBeatCheck)
	if ok {
		node, ok = value.(timer.TimeNoder)
		if ok {
			node.Stop()
		}
	}
	// 设置定时器，超时未心跳就断开连接
	node = uc.timer.AfterFunc(uc.cfg.Protocol.MaxClientHeartbeatInterval, func() {
		lastTime := ch.Int64Attr(consts.ChannelAttrLastHeartBeatTime)
		nowts := time.Now().UnixMilli()
		if nowts-lastTime > uc.cfg.Protocol.MaxClientHeartbeatInterval.Milliseconds() {
			log.Warnf("channel heartbeat timed out: %v", ch.Id())
			_ = ch.Close()
		}
	})
	ch.SetAttr(consts.ChannelAttrTimerHeartBeatCheck, node)
}

// loadRoom 加载房间列表
func (uc *ChannelUseCase) loadRoom(srv socket.Server, ch socket.Channel) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	rooms, err := uc.roomRepo.List(ctx, &options.ListRoomOptions{
		ProductId: ch.StringAttr(consts.ChannelAttrProductId),
		Uid:       ch.StringAttr(consts.ChannelAttrUid),
	})
	if err != nil {
		log.Errorf("ListRoom error: %v", err)
		return
	}
	for _, room := range rooms {
		_ = srv.JoinRoom(protocol.EncodeRoomId(room), ch)
	}
}
