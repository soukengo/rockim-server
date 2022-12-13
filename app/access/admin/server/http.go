package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	"rockim/app/access/admin/biz/manager"
	"rockim/app/access/admin/conf"
)

type HttpServiceGroup interface {
	Register(*http.Server)
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, authCfg *conf.Auth, managerGroup *ManagerServiceGroup, tenantGroup *TenantServiceGroup) *http.Server {
	var opts = []http.ServerOption{
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "AccessToken", "X-Auth-Token", "X-Token", "Accept"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			jwt.Server(func(token *jwtV4.Token) (interface{}, error) {
				return []byte(authCfg.Jwt.AppKey), nil
			},
				jwt.WithSigningMethod(jwtV4.SigningMethodHS256),
				jwt.WithClaims(func() jwtV4.Claims {
					return manager.Claims{}
				})),
			selector.Server().Match(func(ctx context.Context, operation string) bool {
				return true
			}).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout > 0 {
		opts = append(opts, http.Timeout(c.Http.Timeout))
	}
	srv := http.NewServer(opts...)
	srv.Use("")
	managerGroup.Register(srv)
	tenantGroup.Register(srv)
	return srv
}
