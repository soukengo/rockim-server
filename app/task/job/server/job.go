package server

import (
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/component/server/job"
	"github.com/soukengo/gopkg/log"
)

// NewJobServer new a job server.
func NewJobServer(c *server.Config, group *ServiceGroup, logger log.Logger) (s job.Server, err error) {
	srv, err := job.NewServer(c.Job, logger)
	if err != nil {
		return
	}
	group.Register(srv)
	return srv, nil
}
