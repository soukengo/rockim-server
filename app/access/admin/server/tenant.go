package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	jwtAuth "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	v1 "rockimserver/apis/rockim/api/admin/tenant/v1"
	"rockimserver/app/access/admin/conf"
	"rockimserver/app/access/admin/module/tenant/biz"
	"rockimserver/app/access/admin/module/tenant/service"
)

type TenantServiceGroup struct {
	authCfg    *conf.Auth
	authSrv    *service.AuthService
	sessionSrv *service.SessionService
	productSrv *service.ProductService
}

func NewTenantServiceGroup(authCfg *conf.Auth, authSrv *service.AuthService, sessionSrv *service.SessionService, productSrv *service.ProductService) *TenantServiceGroup {
	return &TenantServiceGroup{authCfg: authCfg, authSrv: authSrv, sessionSrv: sessionSrv, productSrv: productSrv}
}

func (g *TenantServiceGroup) Register(srv *http.Server) {
	srv.Use("/rockim.api.admin.tenant.v1.*", g.checkAuth())
	v1.RegisterAuthAPIHTTPServer(srv, g.authSrv)
	v1.RegisterSessionAPIHTTPServer(srv, g.sessionSrv)
	v1.RegisterProductAPIHTTPServer(srv, g.productSrv)
}

func (g *TenantServiceGroup) checkAuth() middleware.Middleware {
	whiteList := make(map[string]struct{})
	whiteList[v1.OperationAuthAPILogin] = struct{}{}
	return selector.Server(jwtAuth.Server(func(token *jwtV4.Token) (interface{}, error) {
		return []byte(g.authCfg.Jwt.AppKey), nil
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
