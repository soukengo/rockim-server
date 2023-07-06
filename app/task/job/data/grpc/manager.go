package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/app/task/job/conf"
	"sync"
)

type PushManager struct {
	config       *conf.Config
	cometServers map[string]*Comet
	rooms        map[string]*Room
	roomsMutex   sync.RWMutex
	watcher      *Watcher
}

func NewPushManager(config *conf.Config, discovery registry.Discovery) (*PushManager, error) {
	m := &PushManager{
		config:       config,
		cometServers: make(map[string]*Comet),
		rooms:        make(map[string]*Room),
	}
	w, err := newWatcher(m, discovery)
	if err != nil {
		return nil, err
	}
	m.watcher = w
	return m, nil
}

func (c *PushManager) delRoom(roomId string) {
	c.roomsMutex.Lock()
	delete(c.rooms, roomId)
	c.roomsMutex.Unlock()
}

func (c *PushManager) getRoom(r *types.Room) *Room {
	roomId := encodeRoomId(r)
	c.roomsMutex.RLock()
	room, ok := c.rooms[roomId]
	c.roomsMutex.RUnlock()
	if !ok {
		c.roomsMutex.Lock()
		if room, ok = c.rooms[roomId]; !ok {
			room = newRoom(c, roomId, r, c.config.Comet.Room)
			c.rooms[roomId] = room
		}
		c.roomsMutex.Unlock()
	}
	return room
}
