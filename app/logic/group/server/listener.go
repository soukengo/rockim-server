package server

import (
	"github.com/soukengo/gopkg/component/transport/event"
	"github.com/soukengo/gopkg/component/transport/queue/options"
	bizevent "rockimserver/app/logic/group/biz/event"
	"rockimserver/app/logic/group/listener"
)

type ListenerRegistry struct {
	group *listener.GroupEventListener
}

func NewListenerRegistry(group *listener.GroupEventListener) *ListenerRegistry {
	return &ListenerRegistry{group: group}
}

func (g *ListenerRegistry) RegisterEvent(srv event.Server) {
	srv.Subscribe(bizevent.KeyGroupJoined, event.Handle[bizevent.GroupJoinedEvent](g.group.OnGroupJoined, event.Consumer().SetMode(options.Async)))
}
