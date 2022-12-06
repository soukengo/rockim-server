package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "rockim/api/rockim/admin/manager/v1"
	"rockim/app/access/admin/service/manager"
)

type ManagerServiceGroup struct {
	authSrv *manager.AuthService
}

func NewManagerServiceGroup(authSrv *manager.AuthService) *ManagerServiceGroup {
	return &ManagerServiceGroup{authSrv: authSrv}
}

func (g *ManagerServiceGroup) Register(srv *http.Server) {
	v1.RegisterAuthAPIHTTPServer(srv, g.authSrv)
}
