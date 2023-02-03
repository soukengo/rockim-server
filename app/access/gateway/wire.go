//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package gateway

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"rockimserver/app/access/gateway/biz"
	"rockimserver/app/access/gateway/conf"
	"rockimserver/app/access/gateway/data"
	"rockimserver/app/access/gateway/server"
	"rockimserver/app/access/gateway/service"
	"rockimserver/pkg/component/discovery"
)

// wireApp init kratos application.
func wireApp(*conf.Env, *discovery.Config, *conf.Server) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
