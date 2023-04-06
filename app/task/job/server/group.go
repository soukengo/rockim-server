package server

import (
	"github.com/soukengo/gopkg/component/server/job"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/task/job/task"
)

type ServiceGroup struct {
	messageTask *task.MessageTask
}

func NewServiceGroup(messageTask *task.MessageTask) *ServiceGroup {
	return &ServiceGroup{messageTask: messageTask}
}

func (g *ServiceGroup) Register(srv job.Server) {
	srv.Register(enums.MQ_MESSAGE_PUSH.String(), wrapAction[v1.MessagePushRequest](g.messageTask.Push))
}
