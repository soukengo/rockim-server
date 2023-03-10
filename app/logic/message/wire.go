//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package message

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"rockimserver/app/logic/message/biz"
	"rockimserver/app/logic/message/conf"
	"rockimserver/app/logic/message/data"
	"rockimserver/app/logic/message/server"
	"rockimserver/app/logic/message/service"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/component/discovery"
	servercomponent "rockimserver/pkg/component/server"
)

// wireApp init kratos application.
func wireApp(*conf.Config, *discovery.Config, *servercomponent.Config, *mongo.Config, *redis.Config, *cache.Config) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, service.ProviderSet, data.ProviderSet, newApp))
}
