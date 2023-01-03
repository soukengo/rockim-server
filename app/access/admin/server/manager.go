package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	jwtAuth "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	v1 "rockim/api/rockim/admin/manager/v1"
	bizManager "rockim/app/access/admin/biz/manager"
	"rockim/app/access/admin/conf"
	"rockim/app/access/admin/service/manager"
)

type ManagerServiceGroup struct {
	authCfg           *conf.Auth
	authSrv           *manager.AuthService
	sessionSrv        *manager.SessionService
	platUserSrv       *manager.PlatUserService
	platRoleSrv       *manager.PlatRoleService
	platResourceSrv   *manager.PlatResourceService
	tenantSrv         *manager.TenantService
	tenantResourceSrv *manager.TenantResourceService
}

func NewManagerServiceGroup(authCfg *conf.Auth, authSrv *manager.AuthService, sessionSrv *manager.SessionService, platUserSrv *manager.PlatUserService, platRoleSrv *manager.PlatRoleService, platResourceSrv *manager.PlatResourceService, tenantSrv *manager.TenantService, tenantResourceSrv *manager.TenantResourceService) *ManagerServiceGroup {
	return &ManagerServiceGroup{authCfg: authCfg, authSrv: authSrv, sessionSrv: sessionSrv, platUserSrv: platUserSrv, platRoleSrv: platRoleSrv, platResourceSrv: platResourceSrv, tenantSrv: tenantSrv, tenantResourceSrv: tenantResourceSrv}
}

func (g *ManagerServiceGroup) Register(srv *http.Server) {
	srv.Use("/rockim.admin.manager.v1.*", checkManagerAuth(g.authCfg))
	v1.RegisterAuthAPIHTTPServer(srv, g.authSrv)
	v1.RegisterSessionAPIHTTPServer(srv, g.sessionSrv)
	v1.RegisterPlatUserAPIHTTPServer(srv, g.platUserSrv)
	v1.RegisterPlatRoleAPIHTTPServer(srv, g.platRoleSrv)
	v1.RegisterPlatResourceAPIHTTPServer(srv, g.platResourceSrv)
	v1.RegisterTenantAPIHTTPServer(srv, g.tenantSrv)
	v1.RegisterTenantResourceAPIHTTPServer(srv, g.tenantResourceSrv)
}

func checkManagerAuth(authCfg *conf.Auth) middleware.Middleware {
	whiteList := make(map[string]struct{})
	whiteList[v1.OperationAuthAPILogin] = struct{}{}
	return selector.Server(jwtAuth.Server(func(token *jwtV4.Token) (interface{}, error) {
		return []byte(authCfg.Jwt.AppKey), nil
	},
		jwtAuth.WithSigningMethod(jwtV4.SigningMethodHS256),
		jwtAuth.WithClaims(func() jwtV4.Claims {
			return &bizManager.Claims{}
		}),
	)).Match(func(ctx context.Context, operation string) bool {
		_, ok := whiteList[operation]
		return !ok
	}).Build()
}
