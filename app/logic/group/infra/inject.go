package infra

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/component/idgen"
	"github.com/soukengo/gopkg/component/lock"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/group/conf"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	discovery.NewRegistrar,
	NewCacheManager,
	NewDatabaseManager,
	idgen.NewMongoGenerator,
	NewLockBuilder,
)

func NewCacheManager(cfg *conf.Config, logger log.Logger) *cache.Manager {
	return cache.NewManager(cfg.Cache, logger)
}
func NewDatabaseManager(cfg *conf.Config) *database.Manager {
	return database.NewManager(cfg.Database)
}

func NewLockBuilder(cfg *conf.Config, logger log.Logger) lock.Builder {
	return lock.NewBuilder(cfg.Lock, logger)
}
