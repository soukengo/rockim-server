package biz

import (
	"context"
	"github.com/samber/lo"
	"rockimserver/apis/rockim/service/session/v1/types"
	"rockimserver/app/logic/session/biz/options"
	biztypes "rockimserver/app/logic/session/biz/types"
)

type ChannelQueryRepo interface {
	Exists(ctx context.Context, productId string, uid string) (bool, error)
	ListByUid(ctx context.Context, productId string, uid string) (list []*biztypes.Channel, err error)
}

type ChannelQueryUseCase struct {
	repo ChannelQueryRepo
}

func NewChannelQueryUseCase(repo ChannelQueryRepo) *ChannelQueryUseCase {
	return &ChannelQueryUseCase{repo: repo}
}

func (uc *ChannelQueryUseCase) CheckOnline(ctx context.Context, opts *options.OnlineCheckOptions) (isOnline bool, err error) {
	return uc.repo.Exists(ctx, opts.ProductId, opts.Uid)
}

func (uc *ChannelQueryUseCase) BatchCheckOnline(ctx context.Context, opts *options.OnlineBatchCheckOptions) (onlineIds []string, err error) {
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
func (uc *ChannelQueryUseCase) ListUser(ctx context.Context, opts *options.UserChannelListOptions) (out []*types.UserChannel, err error) {
	for _, uid := range opts.Uids {
		var channels []*biztypes.Channel
		channels, err = uc.repo.ListByUid(ctx, opts.ProductId, uid)
		if err != nil {
			return
		}
		if len(channels) == 0 {
			continue
		}
		var online = &types.UserChannel{Uid: uid}
		online.Channels = lo.Map(channels, func(item *biztypes.Channel, index int) *types.Channel {
			return &types.Channel{ServerId: item.ServerId, ChannelId: item.ChannelId}
		})
		out = append(out, online)
	}
	return
}
