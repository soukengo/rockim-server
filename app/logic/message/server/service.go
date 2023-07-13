package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "rockimserver/apis/rockim/service/message/v1"
	"rockimserver/app/logic/message/service"
)

type ServiceGroup struct {
	message *service.MessageService
}

func NewServiceGroup(message *service.MessageService) *ServiceGroup {
	return &ServiceGroup{message: message}
}

func (g *ServiceGroup) RegisterGrpc(srv *grpc.Server) {
	v1.RegisterMessageAPIServer(srv, g.message)
}
