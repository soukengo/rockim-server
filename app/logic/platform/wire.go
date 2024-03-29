//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package platform

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/platform/biz"
	"rockimserver/app/logic/platform/conf"
	"rockimserver/app/logic/platform/data"
	"rockimserver/app/logic/platform/server"
	"rockimserver/app/logic/platform/service"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Config, *discovery.Config, *conf.Server, *database.Config, *cache.Config) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
