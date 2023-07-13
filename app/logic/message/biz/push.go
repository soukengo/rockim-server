package biz

import (
	"context"
	"github.com/samber/lo"
	comettypes "rockimserver/apis/rockim/service/comet/v1/types"
	"rockimserver/apis/rockim/service/job/v1/types"
	usertypes "rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/apis/rockim/shared/enums"
)

type PushManager interface {
	PushUsers(ctx context.Context, operation enums.Comet_PushOperation, productId string, uids []string, body []byte) error
	PushGroup(ctx context.Context, operation enums.Comet_PushOperation, productId string, groupId string, body []byte) error
}

type PushMessageRepo interface {
	SavePushMessage(ctx context.Context, messages []*types.CometMessage) error
}

type pushManager struct {
	repo       PushMessageRepo
	onlineRepo OnlineRepo
}

func NewPushManager(repo PushMessageRepo, onlineRepo OnlineRepo) PushManager {
	return &pushManager{repo: repo, onlineRepo: onlineRepo}
}

func (m *pushManager) PushUsers(ctx context.Context, operation enums.Comet_PushOperation, productId string, uids []string, body []byte) (err error) {
	users, err := m.onlineRepo.ListUser(ctx, productId, uids)
	if err != nil {
		return
	}
	channels := lo.FlatMap(users, func(item *usertypes.OnlineUser, index int) []*usertypes.OnlineChannel {
		return item.Channels
	})
	channelsMap := lo.GroupBy(channels, func(item *usertypes.OnlineChannel) string {
		return item.ServerId
	})
	var messages []*types.CometMessage
	for serverId, chs := range channelsMap {
		channelIds := lo.Map(chs, func(item *usertypes.OnlineChannel, index int) string {
			return item.ChannelId
		})
		messages = append(messages, &types.CometMessage{
			ProductId: productId,
			Target: &types.CometMessage_Target{
				TargetType: types.CometMessage_Target_CHANNEL,
				ServerId:   serverId,
				Channels:   channelIds,
			},
			Message: &comettypes.Message{
				Push: &comettypes.PushMessage{
					Operation: operation,
					Body:      body,
				},
			},
		})
	}
	if len(messages) == 0 {
		return
	}
	err = m.repo.SavePushMessage(ctx, messages)
	return
}

func (m *pushManager) PushGroup(ctx context.Context, operation enums.Comet_PushOperation, productId string, groupId string, body []byte) (err error) {
	messages := []*types.CometMessage{
		{
			ProductId: productId,
			Target: &types.CometMessage_Target{
				TargetType: types.CometMessage_Target_ROOM,
				Room:       &comettypes.Room{RoomType: enums.Comet_Group, BizId: groupId},
			},
			Message: &comettypes.Message{
				Push: &comettypes.PushMessage{
					Operation: operation,
					Body:      body,
				},
			},
		},
	}
	err = m.repo.SavePushMessage(ctx, messages)
	return
}
