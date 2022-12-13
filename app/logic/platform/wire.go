//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package platform

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"rockim/app/logic/platform/biz"
	"rockim/app/logic/platform/conf"
	"rockim/app/logic/platform/data"
	"rockim/app/logic/platform/server"
	"rockim/app/logic/platform/service"
	"rockim/pkg/component/database/mongo"
	"rockim/pkg/component/discovery"
)

// wireApp init kratos application.
func wireApp(*conf.Env, *discovery.Config, *conf.Server, *mongo.Config) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
