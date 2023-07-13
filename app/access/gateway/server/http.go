package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/log"
	_ "github.com/soukengo/gopkg/util/encoding/proto"
	"rockimserver/app/access/gateway/server/middleware"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *server.Config, clientGroup *ClientServiceGroup, openapiGroup *OpenApiServiceGroup, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			middleware.Logging(logger),
			middleware.Trace(),
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
	clientGroup.RegisterHttp(srv)
	openapiGroup.RegisterHttp(srv)
	return srv
}
