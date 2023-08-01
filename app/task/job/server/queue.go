package server

import (
	"github.com/soukengo/gopkg/component/transport/queue"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/task/job/conf"
)

// NewQueueServer new a job server.
func NewQueueServer(c *conf.Server, logger log.Logger) (s queue.Server, err error) {
	srv, err := queue.NewGeneralServer(c.Queue, logger)
	if err != nil {
		return
	}
	return srv, nil
}
