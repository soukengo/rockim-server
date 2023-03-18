package server

import (
	"rockimserver/pkg/component/mq"
	"rockimserver/pkg/component/server"
	"rockimserver/pkg/component/server/job"
	"rockimserver/pkg/log"
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
