package server

import (
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/component/server/job"
	"github.com/soukengo/gopkg/log"
)

// NewJobServer new a job server.
func NewJobServer(c *server.Config, room *ServiceGroup, logger log.Logger) (s job.Server, err error) {
	srv, err := job.NewServer(c.Job, logger)
	if err != nil {
		return
	}
	room.RegisterJob(srv)
	return srv, nil
}
