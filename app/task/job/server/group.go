package server

import (
	"github.com/soukengo/gopkg/component/server/job"
	service "rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/app/task/job/task"
)

type ServiceRoom struct {
	messageTask *task.MessageTask
}

func NewServiceRoom(messageTask *task.MessageTask) *ServiceRoom {
	return &ServiceRoom{messageTask: messageTask}
}

func (g *ServiceRoom) Register(srv job.Server) {
	srv.Register(service.MQ_MESSAGE_PUSH.String(), wrapAction[v1.MessagePushRequest](g.messageTask.Push))
}
