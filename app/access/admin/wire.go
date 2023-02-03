//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package admin

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"rockimserver/app/access/admin/conf"
	"rockimserver/app/access/admin/infra"
	"rockimserver/app/access/admin/module/manager"
	"rockimserver/app/access/admin/module/tenant"
	"rockimserver/app/access/admin/server"
	"rockimserver/pkg/component/discovery"
)

// wireApp init kratos application.
func wireApp(*conf.Env, *discovery.Config, *conf.Server, *conf.Auth, *conf.Database) (*kratos.App, error) {
	panic(wire.Build(infra.ProviderSet, manager.ProviderSet, tenant.ProviderSet, server.ProviderSet, newApp))
}
