//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package user

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"rockimserver/app/logic/user/biz"
	"rockimserver/app/logic/user/conf"
	"rockimserver/app/logic/user/data"
	"rockimserver/app/logic/user/server"
	"rockimserver/app/logic/user/service"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database"
	"rockimserver/pkg/component/discovery"
	servercomponent "rockimserver/pkg/component/server"
	"rockimserver/pkg/log"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Config, *discovery.Config, *servercomponent.Config, *database.Config, *cache.Config) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
