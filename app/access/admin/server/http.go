package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"rockim/app/access/admin/conf"
	_ "rockim/pkg/encoding/proto"
)

type HttpServiceGroup interface {
	Register(*http.Server)
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, managerGroup *ManagerServiceGroup, tenantGroup *TenantServiceGroup) *http.Server {
	var opts = []http.ServerOption{
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
	managerGroup.Register(srv)
	tenantGroup.Register(srv)
	return srv
}
