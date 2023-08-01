//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package job

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/task/job/biz"
	"rockimserver/app/task/job/conf"
	"rockimserver/app/task/job/data"
	"rockimserver/app/task/job/infra"
	"rockimserver/app/task/job/server"
	"rockimserver/app/task/job/service"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Config, *discovery.Config, *conf.Server) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, infra.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
