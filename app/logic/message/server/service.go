package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "rockimserver/apis/rockim/service/message/v1"
	"rockimserver/app/logic/message/service"
)

type ServiceRegistry struct {
	message *service.MessageService
}

func NewServiceRegistry(message *service.MessageService) *ServiceRegistry {
	return &ServiceRegistry{message: message}
}

func (g *ServiceRegistry) RegisterGrpc(srv *grpc.Server) {
	v1.RegisterMessageAPIServer(srv, g.message)
}
