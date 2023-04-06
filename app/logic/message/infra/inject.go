package infra

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/database/redis"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/component/idgen"
	"github.com/soukengo/gopkg/component/lock"
	"github.com/soukengo/gopkg/component/mq"
	"github.com/soukengo/gopkg/component/mq/kafka"
	"github.com/soukengo/gopkg/component/queue"
	"rockimserver/app/logic/message/biz/consts"
	"rockimserver/app/logic/message/infra/grpc"
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
