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
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/discovery"
)

// wireApp init kratos application.
func wireApp(*conf.Env, *discovery.Config, *conf.Server, *mongo.Config) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
