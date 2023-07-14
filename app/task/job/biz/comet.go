package biz

import (
	"context"
	"github.com/samber/lo"
	"github.com/soukengo/gopkg/log"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/apis/rockim/service/job/v1/types"
	sessiontypes "rockimserver/apis/rockim/service/session/v1/types"
)

type CometRepo interface {
	Dispatch(ctx context.Context, message *types.CometMessage) error
}

type ChannelRepo interface {
	ListChannels(ctx context.Context, productId string, uids []string) ([]*sessiontypes.UserChannel, error)
}

type CometUseCase struct {
	repo        CometRepo
	sessionRepo ChannelRepo
}

func NewCometUseCase(repo CometRepo, sessionRepo ChannelRepo) *CometUseCase {
	return &CometUseCase{repo: repo, sessionRepo: sessionRepo}
}

func (uc *CometUseCase) Dispatch(ctx context.Context, in *v1.CometDispatchRequest) (err error) {
	for _, message := range in.List {
		err1 := uc.dispatch(ctx, message)
		if err1 != nil {
			log.WithContext(ctx).Error("dispatch message error", log.Pairs{"err": err1})
			return
		}
	}
	return
}
func (uc *CometUseCase) dispatch(ctx context.Context, message *types.CometMessage) (err error) {
	if message.Target.TargetType == types.CometMessage_Target_USER {
		return uc.dispatchUser(ctx, message)
	}
	return uc.repo.Dispatch(ctx, message)
}

func (uc *CometUseCase) dispatchUser(ctx context.Context, message *types.CometMessage) (err error) {
	userChannels, err := uc.sessionRepo.ListChannels(ctx, message.ProductId, message.Target.Uids)
	if err != nil {
		return
	}
	channels := lo.FlatMap(userChannels, func(item *sessiontypes.UserChannel, index int) []*sessiontypes.Channel {
		return item.Channels
	})
	channelsMap := lo.GroupBy(channels, func(item *sessiontypes.Channel) string {
		return item.ServerId
	})

	for sid, chs := range channelsMap {
		cids := lo.Map(chs, func(item *sessiontypes.Channel, index int) string {
			return item.ChannelId
		})
		newMessage := &types.CometMessage{
			ProductId: message.ProductId,
			Target: &types.CometMessage_Target{
				TargetType: types.CometMessage_Target_CHANNEL,
				ServerId:   sid,
				Channels:   cids,
			},
			Message: message.Message,
		}
		err1 := uc.repo.Dispatch(ctx, newMessage)
		if err1 != nil {
			log.WithContext(ctx).Error("dispatch message error", log.Pairs{"err": err1})
			return
		}
	}
	return
}
