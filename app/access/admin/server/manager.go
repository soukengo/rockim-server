package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	jwtAuth "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	v1 "rockim/api/rockim/admin/manager/v1"
	"rockim/app/access/admin/conf"
	"rockim/app/access/admin/module/manager/biz"
	"rockim/app/access/admin/module/manager/service"
)

type ManagerServiceGroup struct {
	authCfg           *conf.Auth
	authSrv           *service.AuthService
	sessionSrv        *service.SessionService
	platUserSrv       *service.SysUserService
	platRoleSrv       *service.SysRoleService
	platResourceSrv   *service.SysResourceService
	tenantSrv         *service.TenantService
	tenantResourceSrv *service.SysTenantResourceService
}

func NewManagerServiceGroup(authCfg *conf.Auth, authSrv *service.AuthService, sessionSrv *service.SessionService, platUserSrv *service.SysUserService, platRoleSrv *service.SysRoleService, platResourceSrv *service.SysResourceService, tenantSrv *service.TenantService, tenantResourceSrv *service.SysTenantResourceService) *ManagerServiceGroup {
	return &ManagerServiceGroup{authCfg: authCfg, authSrv: authSrv, sessionSrv: sessionSrv, platUserSrv: platUserSrv, platRoleSrv: platRoleSrv, platResourceSrv: platResourceSrv, tenantSrv: tenantSrv, tenantResourceSrv: tenantResourceSrv}
}

func (g *ManagerServiceGroup) Register(srv *http.Server) {
	srv.Use("/rockim.admin.manager.v1.*", checkManagerAuth(g.authCfg))
	v1.RegisterAuthAPIHTTPServer(srv, g.authSrv)
	v1.RegisterSessionAPIHTTPServer(srv, g.sessionSrv)
	v1.RegisterSysUserAPIHTTPServer(srv, g.platUserSrv)
	v1.RegisterSysRoleAPIHTTPServer(srv, g.platRoleSrv)
	v1.RegisterSysResourceAPIHTTPServer(srv, g.platResourceSrv)
	v1.RegisterTenantAPIHTTPServer(srv, g.tenantSrv)
	v1.RegisterSysTenantResourceAPIHTTPServer(srv, g.tenantResourceSrv)
}

func checkManagerAuth(authCfg *conf.Auth) middleware.Middleware {
	whiteList := make(map[string]struct{})
	whiteList[v1.OperationAuthAPILogin] = struct{}{}
	return selector.Server(jwtAuth.Server(func(token *jwtV4.Token) (interface{}, error) {
		return []byte(authCfg.Jwt.AppKey), nil
	},
		jwtAuth.WithSigningMethod(jwtV4.SigningMethodHS256),
		jwtAuth.WithClaims(func() jwtV4.Claims {
			return &biz.Claims{}
		}),
	)).Match(func(ctx context.Context, operation string) bool {
		_, ok := whiteList[operation]
		return !ok
	}).Build()
}
