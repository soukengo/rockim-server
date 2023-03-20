package biz

import (
	"context"
	"github.com/samber/lo"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/app/logic/user/biz/options"
	biztypes "rockimserver/app/logic/user/biz/types"
)

type OnlineQueryRepo interface {
	Exists(ctx context.Context, productId string, uid string) (bool, error)
	ListByUid(ctx context.Context, productId string, uid string) (list []*biztypes.OnlineChannel, err error)
}

type OnlineQueryUseCase struct {
	repo OnlineQueryRepo
}

func NewOnlineQueryUseCase(repo OnlineQueryRepo) *OnlineQueryUseCase {
	return &OnlineQueryUseCase{repo: repo}
}

func (uc *OnlineQueryUseCase) CheckOnline(ctx context.Context, opts *options.OnlineCheckOptions) (isOnline bool, err error) {
	return uc.repo.Exists(ctx, opts.ProductId, opts.Uid)
}

func (uc *OnlineQueryUseCase) BatchCheckOnline(ctx context.Context, opts *options.OnlineBatchCheckOptions) (onlineIds []string, err error) {
	for _, uid := range opts.Uids {
		var online bool
		online, err = uc.repo.Exists(ctx, opts.ProductId, uid)
		if err != nil {
			return
		}
		if online {
			onlineIds = append(onlineIds, uid)
		}
	}
	return
}
func (uc *OnlineQueryUseCase) ListUser(ctx context.Context, opts *options.OnlineUserListOptions) (out []*types.OnlineUser, err error) {
	for _, uid := range opts.Uids {
		var channels []*biztypes.OnlineChannel
		channels, err = uc.repo.ListByUid(ctx, opts.ProductId, uid)
		if err != nil {
			return
		}
		if len(channels) == 0 {
			continue
		}
		var online = &types.OnlineUser{Uid: uid}
		online.Channels = lo.Map(channels, func(item *biztypes.OnlineChannel, index int) *types.OnlineChannel {
			return &types.OnlineChannel{ServerId: item.ServerId, ChannelId: item.ChannelId}
		})
		out = append(out, online)
	}
	return
}
