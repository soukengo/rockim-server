package grpc

import (
	"context"
	"fmt"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/comet/v1"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/apis/rockim/shared/enums"
)

func (c *PushManager) Push(ctx context.Context, message *types.Message) (err error) {
	switch message.Target {
	case types.Message_CHANNEL:
		err = c.pushChannels(message.Operation, message.ProductId, message.ServerId, message.Channels, message.Body)
	case types.Message_ROOM:
		err = c.getRoom(message.Room).Push(message.Operation, message.ProductId, message.Body)
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

// pushRoom push a message to room
func (c *PushManager) pushRoom(req *v1.PushRoomRequest) (err error) {
	comets := c.cometServers
	for serverID, c := range comets {
		if err = c.PushRoom(req); err != nil {

			log.Errorf("cfg.BroadcastRoom roomType:%d,bizId: %s serverID:%s error(%v)", req.Room.RoomType, req.Room.BizId, serverID, err)
		}
	}
	return
}

const (
	roomIdSep = "#"
)

func encodeRoomId(room *types.Room) string {
	return room.RoomType.String() + roomIdSep + room.BizId
}
