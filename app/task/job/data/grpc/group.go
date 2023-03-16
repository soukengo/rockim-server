package grpc

import (
	"errors"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/comet/v1"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/task/job/conf"
	"rockimserver/pkg/log"
	"time"
)

var (
	// ErrComet commet error.
	ErrComet = errors.New("comet rpc is not available")
	// ErrCometFull comet chan full.
	ErrCometFull = errors.New("comet proto chan full")
	// ErrGroupFull group chan full.
	ErrGroupFull = errors.New("group proto chan full")

	groupFinishedProto = new(v1.PushGroupRequest)
)

// Group group.
type Group struct {
	c         *conf.Group
	container *PushManager
	id        string
	proto     chan *v1.PushGroupRequest
}

// newGroup new a group struct, store channel group info.
func newGroup(job *PushManager, id string, c *conf.Group) (r *Group) {
	r = &Group{
		c:         c,
		id:        id,
		container: job,
		proto:     make(chan *v1.PushGroupRequest, c.Batch*2),
	}
	go r.pushproc()
	return
}

// Push push msg to the group, if chan full discard it.
func (r *Group) Push(operation enums.Network_PushOperation, productId string, msg []byte) (err error) {
	var p = &v1.PushGroupRequest{
		Base:      service.GenRequest(productId),
		GroupId:   r.id,
		Operation: operation,
		Body:      msg,
	}
	select {
	case r.proto <- p:
	default:
		err = ErrGroupFull
	}
	return
}

// pushproc merge proto and push msgs in batch.
func (r *Group) pushproc() {
	var (
		p *v1.PushGroupRequest
	)
	td := time.AfterFunc(r.c.Idle, func() {
		select {
		case r.proto <- groupFinishedProto:
		default:
		}
	})
	defer td.Stop()
	for {
		if p = <-r.proto; p == nil {
			break // exit
		} else if p != groupFinishedProto {
			_ = r.container.pushGroup(p)
			if r.c.Idle != 0 {
				td.Reset(r.c.Idle)
			} else {
				td.Reset(time.Minute)
			}
		} else {
			break
		}
	}
	log.Infof("pushproc finished, groupId: %v", r.id)
	r.container.delGroup(r.id)
}
