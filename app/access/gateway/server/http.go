package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "rockim/api/access/gateway/client/v1"
	"rockim/app/access/gateway/conf"
	"rockim/app/access/gateway/server/middleware"
	"rockim/app/access/gateway/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, userSrv *service.UserService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
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
	srv.Use("/api/v1", middleware.V1Auth)
	v1.RegisterUserAPIHTTPServer(srv, userSrv)
	srv.Use("/api/v2", middleware.V2Auth)
	return srv
}
