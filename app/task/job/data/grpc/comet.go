package grpc

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/soukengo/gopkg/log"
	"github.com/soukengo/gopkg/util/runtimes"
	ggrpc "google.golang.org/grpc"
	"rockimserver/apis/rockim/service/comet/v1"
	"rockimserver/app/task/job/conf"
	"sync/atomic"
	"time"
)

const (
	timeout = time.Second * 10
)

func newGrpcClient(addr string) (*ggrpc.ClientConn, error) {
	return grpc.DialInsecure(
		context.Background(),
		grpc.WithTimeout(timeout),
		grpc.WithEndpoint(addr),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
	)
}

func newCometClient(addr string) (v1.ChannelAPIClient, error) {
	conn, err := newGrpcClient(addr)
	if err != nil {
		return nil, err
	}
	return v1.NewChannelAPIClient(conn), err
}

// Comet is a comet.
type Comet struct {
	serverID    string
	client      v1.ChannelAPIClient
	pushChan    []chan *v1.PushRequest
	roomChan    []chan *v1.PushRoomRequest
	pushChanNum uint64
	roomChanNum uint64
	routineSize uint64

	ctx    context.Context
	cancel context.CancelFunc
}

// newComet new a comet.
func newComet(hostname string, grpcAddr string, c *conf.Comet) (*Comet, error) {
	cmt := &Comet{
		serverID:    hostname,
		pushChan:    make([]chan *v1.PushRequest, c.RoutineSize),
		roomChan:    make([]chan *v1.PushRoomRequest, c.RoutineSize),
		routineSize: uint64(c.RoutineSize),
	}
	var err error
	if cmt.client, err = newCometClient(grpcAddr); err != nil {
		return nil, err
	}
	cmt.ctx, cmt.cancel = context.WithCancel(context.Background())

	for i := 0; i < c.RoutineSize; i++ {
		cmt.pushChan[i] = make(chan *v1.PushRequest, c.RoutineChan)
		cmt.roomChan[i] = make(chan *v1.PushRoomRequest, c.RoutineChan)
		go cmt.process(cmt.pushChan[i], cmt.roomChan[i])
	}
	return cmt, nil
}

// Push push a channel message.
func (c *Comet) Push(arg *v1.PushRequest) (err error) {
	idx := atomic.AddUint64(&c.pushChanNum, 1) % c.routineSize
	c.pushChan[idx] <- arg
	return
}

// PushRoom broadcast a room message.
func (c *Comet) PushRoom(arg *v1.PushRoomRequest) (err error) {
	idx := atomic.AddUint64(&c.roomChanNum, 1) % c.routineSize
	c.roomChan[idx] <- arg
	return
}

func (c *Comet) process(pushChan chan *v1.PushRequest, roomChan chan *v1.PushRoomRequest) {
	for {
		select {
		case args := <-roomChan:
			_, err := c.client.PushRoom(context.Background(), &v1.PushRoomRequest{
				Base:      args.Base,
				Operation: args.Operation,
				Room:      args.Room,
				Body:      args.Body,
			})
			if err != nil {
				log.Errorf("cfg.client.PushRoom(%s, reply) serverId:%s error(%v)", args, c.serverID, err)
			}
		case pushArg := <-pushChan:
			_, err := c.client.Push(context.Background(), &v1.PushRequest{
				Base:       pushArg.Base,
				Operation:  pushArg.Operation,
				ChannelIds: pushArg.ChannelIds,
				Body:       pushArg.Body,
			})
			if err != nil {
				log.Errorf("cfg.client.PushMsg(%s, reply) serverId:%s error(%v)", pushArg, c.serverID, err)
			}
		case <-c.ctx.Done():
			return
		}
	}
}

// Close close the resources.
func (c *Comet) Close() (err error) {
	finish := make(chan bool)
	runtimes.Async(func() {
		for {
			n := 0
			for _, ch := range c.pushChan {
				n += len(ch)
			}
			for _, ch := range c.roomChan {
				n += len(ch)
			}
			if n == 0 {
				finish <- true
				return
			}
			time.Sleep(time.Second)
		}
	})
	select {
	case <-finish:
		log.Infof("close comet finish")
	case <-time.After(5 * time.Second):
		err = fmt.Errorf("close comet(server:%s push:%d room:%d) timeout", c.serverID, len(c.pushChan), len(c.roomChan))
	}
	c.cancel()
	return
}
