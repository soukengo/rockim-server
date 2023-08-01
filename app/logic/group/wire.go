//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package group

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/group/biz"
	"rockimserver/app/logic/group/conf"
	"rockimserver/app/logic/group/data"
	"rockimserver/app/logic/group/infra"
	"rockimserver/app/logic/group/listener"
	"rockimserver/app/logic/group/server"
	"rockimserver/app/logic/group/service"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Config, *discovery.Config, *conf.Server, *database.Config, *cache.Config) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, listener.ProviderSet, infra.ProviderSet, newApp))
}
