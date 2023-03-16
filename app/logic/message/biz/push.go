package biz

import (
	"context"
	"github.com/samber/lo"
	"rockimserver/apis/rockim/service/job/v1/types"
	usertypes "rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/apis/rockim/shared/enums"
)

type PushManager interface {
	PushUsers(ctx context.Context, operation enums.Network_PushOperation, productId string, uids []string, body []byte) error
	PushGroup(ctx context.Context, operation enums.Network_PushOperation, productId string, groupId string, body []byte) error
}

type PushMessageRepo interface {
	SavePushMessage(ctx context.Context, messages []*types.Message) error
}

type pushManager struct {
	repo       PushMessageRepo
	onlineRepo OnlineRepo
}

func NewPushManager(repo PushMessageRepo, onlineRepo OnlineRepo) PushManager {
	return &pushManager{repo: repo, onlineRepo: onlineRepo}
}

func (m *pushManager) PushUsers(ctx context.Context, operation enums.Network_PushOperation, productId string, uids []string, body []byte) (err error) {
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
	var messages []*types.Message
	for serverId, chs := range channelsMap {
		channelIds := lo.Map(chs, func(item *usertypes.OnlineChannel, index int) string {
			return item.ChannelId
		})
		messages = append(messages, &types.Message{
			Target:    types.Message_CHANNEL,
			Operation: operation,
			ServerId:  serverId,
			Channels:  channelIds,
			Body:      body,
		})
	}
	if len(messages) == 0 {
		return
	}
	err = m.repo.SavePushMessage(ctx, messages)
	return
}

func (m *pushManager) PushGroup(ctx context.Context, operation enums.Network_PushOperation, productId string, groupId string, body []byte) (err error) {
	messages := []*types.Message{
		{
			Target:    types.Message_GROUP,
			Operation: operation,
			GroupId:   groupId,
			Body:      body,
		},
	}
	err = m.repo.SavePushMessage(ctx, messages)
	return
}
