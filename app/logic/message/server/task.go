package server

import (
	"rockimserver/app/logic/message/biz/consts"
	"rockimserver/app/logic/message/task"
	"rockimserver/pkg/component/server/job"
)

type TaskGroup struct {
	messageTask *task.MessageTask
}

func NewTaskGroup(messageTask *task.MessageTask) *TaskGroup {
	return &TaskGroup{messageTask: messageTask}
}

func (g *TaskGroup) Register(srv job.Server) {
	srv.Register(string(consts.QueueTopicMessageDelivery), g.messageTask.Delivery)
}
