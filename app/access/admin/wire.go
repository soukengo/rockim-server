//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package admin

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"rockim/app/access/admin/biz"
	"rockim/app/access/admin/conf"
	"rockim/app/access/admin/data"
	"rockim/app/access/admin/server"
	"rockim/app/access/admin/service"
	"rockim/pkg/component/discovery"
)

// wireApp init kratos application.
func wireApp(*conf.Env, *discovery.Config, *conf.Server, *conf.Auth) (*kratos.App, error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
}
