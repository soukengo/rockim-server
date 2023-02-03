package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"rockimserver/app/access/admin/conf"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, managerGroup *ManagerServiceGroup, tenantGroup *TenantServiceGroup) *http.Server {
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
