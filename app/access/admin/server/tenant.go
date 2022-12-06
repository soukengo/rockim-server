package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "rockim/api/rockim/admin/tenant/v1"
	"rockim/app/access/admin/service/tenant"
)

type TenantServiceGroup struct {
	authSrv *tenant.AuthService
}

func NewTenantServiceGroup(authSrv *tenant.AuthService) *TenantServiceGroup {
	return &TenantServiceGroup{authSrv: authSrv}
}

func (g *TenantServiceGroup) Register(srv *http.Server) {
	v1.RegisterAuthAPIHTTPServer(srv, g.authSrv)
}
