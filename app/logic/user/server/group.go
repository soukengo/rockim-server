package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "rockimserver/apis/rockim/service/user/v1"
	"rockimserver/app/logic/user/service"
)

type ServiceGroup struct {
	user *service.UserService
	auth *service.AuthService
}

func NewServiceGroup(user *service.UserService, auth *service.AuthService) *ServiceGroup {
	return &ServiceGroup{user: user, auth: auth}
}

func (g *ServiceGroup) RegisterGrpc(srv *grpc.Server) {
	v1.RegisterUserAPIServer(srv, g.user)
	v1.RegisterAuthAPIServer(srv, g.auth)
}
