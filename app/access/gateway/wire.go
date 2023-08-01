//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package gateway

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/access/gateway/conf"
	"rockimserver/app/access/gateway/infra"
	"rockimserver/app/access/gateway/module"
	"rockimserver/app/access/gateway/server"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Config, *discovery.Config, *conf.Server) (*kratos.App, error) {
	panic(wire.Build(infra.ProviderSet, server.ProviderSet, module.ProviderSet, newApp))
}
