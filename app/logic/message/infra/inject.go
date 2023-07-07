package infra

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/component/idgen"
	"github.com/soukengo/gopkg/component/lock"
	"github.com/soukengo/gopkg/component/queue"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/message/conf"
	"rockimserver/app/logic/message/infra/grpc"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	discovery.NewRegistrar,
	discovery.NewDiscovery,
	grpc.ProviderSet,
	NewLockBuilder,
	NewCacheManager,
	NewDatabaseManager,
	NewGeneralQueueProducer,
	NewDelayedQueueProducer,
	idgen.NewMongoGenerator,
)

func NewLockBuilder(cfg *conf.Config, logger log.Logger) lock.Builder {
	return lock.NewBuilder(cfg.Lock, logger)
}

func NewCacheManager(cfg *conf.Config, logger log.Logger) *cache.Manager {
	return cache.NewManager(cfg.Cache, logger)
}
func NewDatabaseManager(cfg *conf.Config) *database.Manager {
	return database.NewManager(cfg.Database)
}

func NewGeneralQueueProducer(cfg *queue.Config, logger log.Logger) (queue.Producer, error) {
	return queue.NewProducer(cfg.General, logger)
}

func NewDelayedQueueProducer(cfg *queue.Config, logger log.Logger) (producer queue.DelayedProducer, err error) {
	producer, err = queue.NewDelayProducer(cfg.Delayed, logger)
	if err != nil {
		return
	}
	_ = producer.Start()
	return
}
