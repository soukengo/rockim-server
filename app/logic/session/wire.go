//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package session

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/discovery"
	servercomponent "github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/session/biz"
	"rockimserver/app/logic/session/conf"
	"rockimserver/app/logic/session/data"
	"rockimserver/app/logic/session/infra"
	"rockimserver/app/logic/session/server"
	"rockimserver/app/logic/session/service"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Config, *discovery.Config, *servercomponent.Config, *cache.Config) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, infra.ProviderSet, newApp))
}
