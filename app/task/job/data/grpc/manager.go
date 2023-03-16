package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"rockimserver/app/task/job/conf"
	"sync"
)

type PushManager struct {
	config       *conf.Config
	cometServers map[string]*Comet
	groups       map[string]*Group
	groupsMutex  sync.RWMutex
	watcher      *Watcher
}

func NewPushManager(config *conf.Config, discovery registry.Discovery) (*PushManager, error) {
	m := &PushManager{
		config:       config,
		cometServers: make(map[string]*Comet),
		groups:       make(map[string]*Group),
	}
	w, err := newWatcher(m, discovery)
	if err != nil {
		return nil, err
	}
	m.watcher = w
	return m, nil
}

func (c *PushManager) delGroup(groupID string) {
	c.groupsMutex.Lock()
	delete(c.groups, groupID)
	c.groupsMutex.Unlock()
}

func (c *PushManager) getGroup(groupID string) *Group {
	c.groupsMutex.RLock()
	group, ok := c.groups[groupID]
	c.groupsMutex.RUnlock()
	if !ok {
		c.groupsMutex.Lock()
		if group, ok = c.groups[groupID]; !ok {
			group = newGroup(c, groupID, c.config.Comet.Group)
			c.groups[groupID] = group
		}
		c.groupsMutex.Unlock()
	}
	return group
}
