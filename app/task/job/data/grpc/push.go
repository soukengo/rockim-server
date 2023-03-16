package grpc

import (
	"context"
	"fmt"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/comet/v1"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/pkg/log"
)

func (c *PushManager) Push(ctx context.Context, message *types.Message) (err error) {
	switch message.Target {
	case types.Message_CHANNEL:
		err = c.pushChannels(message.Operation, message.ProductId, message.ServerId, message.Channels, message.Body)
	case types.Message_GROUP:
		err = c.getGroup(message.GroupId).Push(message.Operation, message.ProductId, message.Body)
	default:
		err = fmt.Errorf("no match push type: %s", message.Target)
	}
	return
}

// pushChannels push a message to a batch of channels.
func (c *PushManager) pushChannels(operation enums.Network_PushOperation, productId string, serverID string, channelIds []string, body []byte) (err error) {
	p := &v1.PushRequest{
		Base:       service.GenRequest(productId),
		Operation:  operation,
		ChannelIds: channelIds,
		Body:       body,
	}
	if c, ok := c.cometServers[serverID]; ok {
		if err = c.Push(p); err != nil {
			log.Errorf("cfg.Push serverID:%s error(%v)", serverID, err)
		}
	}
	return
}

// pushGroup push a message to group
func (c *PushManager) pushGroup(req *v1.PushGroupRequest) (err error) {
	comets := c.cometServers
	for serverID, c := range comets {
		if err = c.PushGroup(req); err != nil {
			log.Errorf("cfg.BroadcastRoom groupID:%s serverID:%s error(%v)", req.GroupId, serverID, err)
		}
	}
	return
}
