//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package user

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"rockim/internal/business/user/biz"
	"rockim/internal/business/user/conf"
	"rockim/internal/business/user/data"
	"rockim/internal/business/user/server"
	"rockim/internal/business/user/service"
)

// wireApp init kratos application.
func wireApp(*conf.Env, *conf.Server) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
