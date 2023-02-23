//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package platform

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"rockimserver/app/logic/platform/biz"
	"rockimserver/app/logic/platform/conf"
	"rockimserver/app/logic/platform/data"
	"rockimserver/app/logic/platform/server"
	"rockimserver/app/logic/platform/service"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/component/discovery"
	servercomponent "rockimserver/pkg/component/server"
)

// wireApp init kratos application.
func wireApp(*conf.Config, *discovery.Config, *servercomponent.Config, *mongo.Config, *redis.Config, *cache.Config) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
