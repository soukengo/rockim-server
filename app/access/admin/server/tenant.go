package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	jwtAuth "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	v1 "rockim/api/rockim/admin/tenant/v1"
	bizTenant "rockim/app/access/admin/biz/tenant"
	"rockim/app/access/admin/conf"
	"rockim/app/access/admin/service/tenant"
)

type TenantServiceGroup struct {
	authCfg    *conf.Auth
	authSrv    *tenant.AuthService
	sessionSrv *tenant.SessionService
}

func NewTenantServiceGroup(authCfg *conf.Auth, authSrv *tenant.AuthService, sessionSrv *tenant.SessionService) *TenantServiceGroup {
	return &TenantServiceGroup{authCfg: authCfg, authSrv: authSrv, sessionSrv: sessionSrv}
}

func (g *TenantServiceGroup) Register(srv *http.Server) {
	srv.Use("/rockim.admin.tenant.v1.*", checkTenantAuth(g.authCfg))
	v1.RegisterAuthAPIHTTPServer(srv, g.authSrv)
	v1.RegisterSessionAPIHTTPServer(srv, g.sessionSrv)
}

func checkTenantAuth(authCfg *conf.Auth) middleware.Middleware {
	whiteList := make(map[string]struct{})
	whiteList[v1.OperationAuthAPILogin] = struct{}{}
	return selector.Server(jwtAuth.Server(func(token *jwtV4.Token) (interface{}, error) {
		return []byte(authCfg.Jwt.AppKey), nil
	},
		jwtAuth.WithSigningMethod(jwtV4.SigningMethodHS256),
		jwtAuth.WithClaims(func() jwtV4.Claims {
			return &bizTenant.Claims{}
		}),
	)).Match(func(ctx context.Context, operation string) bool {
		_, ok := whiteList[operation]
		return !ok
	}).Build()
}
