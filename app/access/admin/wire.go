//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package admin

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"rockimserver/app/access/admin/conf"
	"rockimserver/app/access/admin/infra"
	"rockimserver/app/access/admin/module"
	"rockimserver/app/access/admin/server"
	"rockimserver/pkg/component/auth"
	"rockimserver/pkg/component/database"
	"rockimserver/pkg/component/discovery"
	servercomponent "rockimserver/pkg/component/server"
	"rockimserver/pkg/log"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Config, *discovery.Config, *servercomponent.Config, *auth.Config, *database.Config) (*kratos.App, error) {
	panic(wire.Build(infra.ProviderSet, module.ProviderSet, server.ProviderSet, newApp))
}
