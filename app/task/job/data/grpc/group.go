package grpc

import (
	"errors"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/comet/v1"
	comettypes "rockimserver/apis/rockim/service/comet/v1/types"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/task/job/conf"
	"time"
)

var (
	// ErrComet commet error.
	ErrComet = errors.New("comet rpc is not available")
	// ErrCometFull comet chan full.
	ErrCometFull = errors.New("comet proto chan full")
	// ErrRoomFull room chan full.
	ErrRoomFull = errors.New("room proto chan full")

	roomFinishedProto = new(v1.PushRoomRequest)
)

// Room room.
type Room struct {
	c         *conf.Room
	container *PushManager
	id        string
	room      *types.Room
	proto     chan *v1.PushRoomRequest
}

// newRoom new a room struct, store channel room info.
func newRoom(job *PushManager, id string, room *types.Room, c *conf.Room) (r *Room) {
	r = &Room{
		c:         c,
		id:        id,
		room:      room,
		container: job,
		proto:     make(chan *v1.PushRoomRequest, c.Batch*2),
	}
	go r.pushproc()
	return
}

// Push push msg to the room, if chan full discard it.
func (r *Room) Push(operation enums.Network_PushOperation, productId string, msg []byte) (err error) {
	var p = &v1.PushRoomRequest{
		Base:      service.GenRequest(productId),
		Room:      &comettypes.Room{RoomType: r.room.RoomType, BizId: r.room.BizId},
		Operation: operation,
		Body:      msg,
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
		p *v1.PushRoomRequest
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
			_ = r.container.pushRoom(p)
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
