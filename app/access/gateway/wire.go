//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package gateway

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"rockimserver/app/access/gateway/conf"
	"rockimserver/app/access/gateway/infra"
	"rockimserver/app/access/gateway/module"
	"rockimserver/app/access/gateway/server"
	"rockimserver/pkg/component/discovery"
	servercomponent "rockimserver/pkg/component/server"
	"rockimserver/pkg/log"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Config, *discovery.Config, *servercomponent.Config) (*kratos.App, error) {
	panic(wire.Build(infra.ProviderSet, server.ProviderSet, module.ProviderSet, newApp))
}
