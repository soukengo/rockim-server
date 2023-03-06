package socket

import (
	"context"
	log "github.com/golang/glog"
	"net"
	"rockimserver/pkg/component/server/socket/network"
	"rockimserver/pkg/component/server/socket/packet"
	"rockimserver/pkg/util/runtimes"
	"sync"
	"sync/atomic"
)

type channel struct {
	clientIP      string
	conn          network.Connection
	done          *sync.Once
	recvQueue     chan packet.IPacket
	sendQueue     chan packet.IPacket
	ctx           context.Context
	cancel        context.CancelFunc
	groups        map[string]struct{}
	glock         sync.RWMutex
	authenticated int32
	Attributes
}

func newChannel(conn network.Connection, recvQueueSize uint32, sendQueueSize uint32) Channel {
	ctx, cancel := context.WithCancel(context.Background())
	ins := &channel{
		ctx:        ctx,
		cancel:     cancel,
		done:       new(sync.Once),
		recvQueue:  make(chan packet.IPacket, recvQueueSize),
		sendQueue:  make(chan packet.IPacket, sendQueueSize),
		conn:       conn,
		groups:     map[string]struct{}{},
		Attributes: newAttributes(),
	}
	ins.clientIP, _, _ = net.SplitHostPort(conn.RemoteAddr().String())
	go runtimes.TryCatch(func() {
		ins.dispatch()
	})
	return ins
}

func (c *channel) Id() string {
	return c.conn.Id()
}

func (c *channel) ClientIP() string {
	return c.clientIP
}

func (c *channel) Send(data packet.IPacket) error {
	select {
	case <-c.ctx.Done():
		return c.ctx.Err()
	default:
		c.sendQueue <- data
	}
	return nil
}

func (c *channel) Receive(p packet.IPacket) {
	select {
	case <-c.ctx.Done():
		return
	default:
		c.recvQueue <- p
	}
}
func (c *channel) AddGroup(groupId string) {
	c.glock.Lock()
	c.groups[groupId] = struct{}{}
	c.glock.Unlock()
}

func (c *channel) DelGroup(groupId string) {
	c.glock.Lock()
	delete(c.groups, groupId)
	c.glock.Unlock()
}
func (c *channel) Groups() (ret []string) {
	c.glock.RLock()
	for id := range c.groups {
		ret = append(ret, id)
	}
	c.glock.RUnlock()
	return
}

func (c *channel) Close() (err error) {
	c.done.Do(func() {
		c.cancel()
		err = c.conn.Close()
	})
	return
}

func (c *channel) MarkAuthenticated() {
	atomic.StoreInt32(&c.authenticated, 1)
}

func (c *channel) Authenticated() bool {
	return atomic.LoadInt32(&c.authenticated) > 0
}

// dispatch send packet to connection
func (c *channel) dispatch() {
	defer func() {
		_ = c.Close()
	}()
	for {
		select {
		case <-c.ctx.Done():
			return
		case p := <-c.sendQueue:
			err := c.conn.Send(p)
			if err != nil {
				log.Errorf("conn.Send err: %v", err)
				return
			}
		}
	}
}
