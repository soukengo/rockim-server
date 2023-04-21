//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package message

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/component/queue"
	servercomponent "github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/message/biz"
	"rockimserver/app/logic/message/conf"
	"rockimserver/app/logic/message/data"
	"rockimserver/app/logic/message/infra"
	"rockimserver/app/logic/message/server"
	"rockimserver/app/logic/message/service"
	"rockimserver/app/logic/message/task"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Config, *discovery.Config, *servercomponent.Config, *database.Config, *cache.Config, *queue.Config) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, infra.ProviderSet, service.ProviderSet, task.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
