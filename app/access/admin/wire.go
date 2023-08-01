//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package admin

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/auth"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/access/admin/conf"
	"rockimserver/app/access/admin/infra"
	"rockimserver/app/access/admin/module"
	"rockimserver/app/access/admin/server"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Config, *discovery.Config, *conf.Server, *auth.Config, *database.Config) (*kratos.App, error) {
	panic(wire.Build(infra.ProviderSet, module.ProviderSet, server.ProviderSet, newApp))
}
