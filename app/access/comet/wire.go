//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package comet

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/access/comet/conf"
	"rockimserver/app/access/comet/infra"
	"rockimserver/app/access/comet/module"
	"rockimserver/app/access/comet/server"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Config, *discovery.Config, *conf.Server, *conf.Protocol) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, module.ProviderSet, infra.ProviderSet, newApp))
}
