package comet

import (
	"errors"
	"github.com/soukengo/gopkg/log"
	"github.com/soukengo/gopkg/util/runtimes"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/comet/v1"
	comettypes "rockimserver/apis/rockim/service/comet/v1/types"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/app/task/job/conf"
	"time"
)

var (
	// ErrRoomFull room chan full.
	ErrRoomFull = errors.New("room proto chan full")

	roomFinishedProto = new(v1.DispatchRoomRequest)
)

// Room room.
type Room struct {
	c         *conf.Room
	container *Manager
	id        string
	room      *comettypes.Room
	proto     chan *v1.DispatchRoomRequest
}

// newRoom new a room struct, store channel room info.
func newRoom(job *Manager, id string, room *comettypes.Room, c *conf.Room) (r *Room) {
	r = &Room{
		c:         c,
		id:        id,
		room:      room,
		container: job,
		proto:     make(chan *v1.DispatchRoomRequest, c.Batch*2),
	}
	runtimes.Async(r.pushproc)
	return
}

// Dispatch push msg to the room, if chan full discard it.
func (r *Room) Dispatch(message *types.CometMessage) (err error) {
	var p = &v1.DispatchRoomRequest{
		Base: service.GenRequest(message.ProductId),
		Room: &comettypes.Room{RoomType: r.room.RoomType, BizId: r.room.BizId},
		Push: message.Message.Push,
	}
	select {
	case r.proto <- p:
	default:
		err = ErrRoomFull
	}
	return
}

// pushproc merge proto and push msgs in batch.
func (r *Room) pushproc() {
	var (
		p *v1.DispatchRoomRequest
	)
	td := time.AfterFunc(r.c.Idle, func() {
		select {
		case r.proto <- roomFinishedProto:
		default:
		}
	})
	defer td.Stop()
	for {
		if p = <-r.proto; p == nil {
			break // exit
		} else if p != roomFinishedProto {
			_ = r.container.dispatchRoom(p)
			if r.c.Idle != 0 {
				td.Reset(r.c.Idle)
			} else {
				td.Reset(time.Minute)
			}
		} else {
			break
		}
	}
	log.Infof("pushproc finished, roomId: %v", r.id)
	r.container.delRoom(r.id)
}
