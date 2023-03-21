package infra

import (
	"github.com/google/wire"
	"rockimserver/app/logic/message/biz/consts"
	"rockimserver/app/logic/message/infra/grpc"
	"rockimserver/pkg/component/database"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/component/discovery"
	"rockimserver/pkg/component/idgen"
	"rockimserver/pkg/component/lock"
	"rockimserver/pkg/component/mq"
	"rockimserver/pkg/component/mq/kafka"
	"rockimserver/pkg/component/queue"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	discovery.NewRegistrar,
	discovery.NewDiscovery,
	grpc.ProviderSet,
	database.NewMongoClient,
	database.NewRedisClient,
	lock.NewRedisBuilder,
	idgen.NewMongoGenerator,
	NewKafkaProducer,
	NewRedisDelayQueue,
)

func NewKafkaProducer(cfg *mq.Config) (mq.Producer, error) {
	return kafka.NewProducer(cfg.Kafka)
}
func NewRedisDelayQueue(cli *redis.Client) queue.Delayed {
	return queue.NewRedisDelayed(cli, []queue.Topic{consts.QueueTopicMessageDelivery}, 10)
}
