//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package user

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/discovery"
	servercomponent "github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/user/biz"
	"rockimserver/app/logic/user/conf"
	"rockimserver/app/logic/user/data"
	"rockimserver/app/logic/user/server"
	"rockimserver/app/logic/user/service"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Config, *discovery.Config, *servercomponent.Config, *database.Config, *cache.Config) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
