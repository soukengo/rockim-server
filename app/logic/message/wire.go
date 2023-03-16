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
	"rockimserver/app/logic/message/infra"
	"rockimserver/app/logic/message/server"
	"rockimserver/app/logic/message/service"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database"
	"rockimserver/pkg/component/discovery"
	"rockimserver/pkg/component/mq"
	servercomponent "rockimserver/pkg/component/server"
)

// wireApp init kratos application.
func wireApp(*conf.Config, *discovery.Config, *servercomponent.Config, *database.Config, *cache.Config, *mq.Config) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, infra.ProviderSet, biz.ProviderSet, service.ProviderSet, data.ProviderSet, newApp))
}