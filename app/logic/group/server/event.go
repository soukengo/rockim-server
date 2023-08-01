package server

import (
	"github.com/soukengo/gopkg/component/transport/event"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/group/conf"
)

func NewEventServer(cfg *conf.Server, logger log.Logger) event.Server {
	return event.NewMemoryServer(cfg.Event.Memory, logger)
}
