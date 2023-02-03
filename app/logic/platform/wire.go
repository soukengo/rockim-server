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
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/discovery"
)

// wireApp init kratos application.
func wireApp(*conf.Env, *discovery.Config, *conf.Server, *mongo.Config) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
