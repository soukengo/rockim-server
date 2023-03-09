package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "rockimserver/apis/rockim/service/user/v1"
	"rockimserver/app/logic/user/service"
)

type ServiceGroup struct {
	user   *service.UserService
	auth   *service.AuthService
	online *service.OnlineService
}

func NewServiceGroup(user *service.UserService, auth *service.AuthService, online *service.OnlineService) *ServiceGroup {
	return &ServiceGroup{user: user, auth: auth, online: online}
}

func (g *ServiceGroup) Register(srv *grpc.Server) {
	v1.RegisterUserAPIServer(srv, g.user)
	v1.RegisterAuthAPIServer(srv, g.auth)
	v1.RegisterOnlineAPIServer(srv, g.online)
}