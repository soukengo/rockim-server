package infra

import (
	"github.com/google/wire"
	"rockimserver/app/logic/message/infra/grpc"
	"rockimserver/pkg/component/database"
	"rockimserver/pkg/component/discovery"
	"rockimserver/pkg/component/idgen"
	"rockimserver/pkg/component/lock"
	"rockimserver/pkg/component/mq"
	"rockimserver/pkg/component/mq/kafka"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	discovery.NewRegistrar,
	discovery.NewDiscovery,
	grpc.ProviderSet,
	database.NewMongoClient,
	database.NewRedisClient,
	NewKafkaProducer,
	lock.NewRedisBuilder,
	idgen.NewMongoGenerator,
)

func NewKafkaProducer(cfg *mq.Config) (mq.Producer, error) {
	return kafka.NewProducer(cfg.Kafka)
}
