package server

import (
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/component/server/job"
	"github.com/soukengo/gopkg/log"
)

func NewJobServer(cfg *server.Config, group *TaskGroup, logger log.Logger) (js job.Server, err error) {
	js, err = job.NewServer(cfg.Job, logger)
	if err != nil {
		return
	}
	group.Register(js)
	return
}
