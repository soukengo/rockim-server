package server

import (
	"github.com/soukengo/gopkg/component/mq"
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/component/server/job"
	"github.com/soukengo/gopkg/log"
)

// NewJobServer new a job server.
func NewJobServer(c *server.Config, mqCfg *mq.Config, group *ServiceGroup, logger log.Logger) (s job.Server, err error) {
	srv, err := job.NewMQServer(c.Job, mqCfg, logger)
	if err != nil {
		return
	}
	group.Register(srv)
	return srv, nil
}
