package server

import (
	"rockimserver/pkg/component/queue"
	"rockimserver/pkg/component/server/job"
)

func NewJobServer(queue queue.Delayed, group *TaskGroup) job.Server {
	js := job.NewDelayQueueServer(queue)
	group.Register(js)
	return js
}
