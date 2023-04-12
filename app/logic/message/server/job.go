package server

import (
	"github.com/soukengo/gopkg/component/queue"
	"github.com/soukengo/gopkg/component/server/job"
)

func NewJobServer(queue queue.Delayed, group *TaskGroup) job.Server {
	js := job.NewQueueServer(queue)
	group.Register(js)
	return js
}
