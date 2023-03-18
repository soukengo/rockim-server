//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package group

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"rockimserver/app/logic/group/biz"
	"rockimserver/app/logic/group/conf"
	"rockimserver/app/logic/group/data"
	"rockimserver/app/logic/group/server"
	"rockimserver/app/logic/group/service"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/component/discovery"
	servercomponent "rockimserver/pkg/component/server"
	"rockimserver/pkg/log"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Config, *discovery.Config, *servercomponent.Config, *mongo.Config, *redis.Config, *cache.Config) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
