package comet

import (
	"context"
	"fmt"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/comet/v1"
	comettypes "rockimserver/apis/rockim/service/comet/v1/types"
	"rockimserver/apis/rockim/service/job/v1/types"
)

func (c *Manager) Dispatch(ctx context.Context, message *types.CometMessage) (err error) {
	switch message.Target.TargetType {
	case types.CometMessage_Target_CHANNEL:
		err = c.DispatchChannels(message)
	case types.CometMessage_Target_ROOM:
		err = c.getRoom(message.Target.Room).Dispatch(message)
	default:
		err = fmt.Errorf("no match push type: %s", message.Target)
	}
	return
}

// DispatchChannels push a message to a batch of channels.
func (c *Manager) DispatchChannels(message *types.CometMessage) (err error) {
	p := &v1.DispatchRequest{
		Base:       service.GenRequest(message.ProductId),
		ChannelIds: message.Target.Channels,
		DataType:   message.Message.DataType,
		Push:       message.Message.Push,
		Control:    message.Message.Control,
	}
	serverId := message.Target.ServerId
	if c, ok := c.cometServers[serverId]; ok {
		if err = c.Dispatch(p); err != nil {
			log.Errorf("cfg.Dispatch serverID:%s error(%v)", serverId, err)
		}
	}
	return
}

// dispatchRoom push a message to room
func (c *Manager) dispatchRoom(req *v1.DispatchRoomRequest) (err error) {
	comets := c.cometServers
	for serverID, c := range comets {
		if err = c.DispatchRoom(req); err != nil {
			log.Errorf("cfg.BroadcastRoom roomType:%d,bizId: %s serverID:%s error(%v)", req.Room.RoomType, req.Room.BizId, serverID, err)
		}
	}
	return
}

const (
	roomIdSep = "#"
)

func encodeRoomId(room *comettypes.Room) string {
	return room.RoomType.String() + roomIdSep + room.BizId
}
