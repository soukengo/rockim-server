//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package gateway

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"rockim/app/access/gateway/biz"
	"rockim/app/access/gateway/conf"
	"rockim/app/access/gateway/data"
	"rockim/app/access/gateway/server"
	"rockim/app/access/gateway/service"
)

// wireApp init kratos application.
func wireApp(*conf.Env, *conf.Server) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
