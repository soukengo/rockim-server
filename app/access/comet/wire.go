//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package comet

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"rockimserver/app/access/comet/conf"
	"rockimserver/app/access/comet/module"
	"rockimserver/app/access/comet/server"
	"rockimserver/pkg/component/discovery"
	servercomponent "rockimserver/pkg/component/server"
)

// wireApp init kratos application.
func wireApp(*conf.Config, *discovery.Config, *servercomponent.Config, *conf.Protocol) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, module.ProviderSet, newApp))
}
