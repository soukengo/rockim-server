package conf

import (
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/config"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/lock"
	"github.com/soukengo/gopkg/component/queue"
	kafkaqueue "github.com/soukengo/gopkg/component/queue/core/kafka"
	redisqueue "github.com/soukengo/gopkg/component/queue/core/redis"
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/component/server/job"
	"github.com/soukengo/gopkg/infra/storage"
	"github.com/soukengo/gopkg/infra/storage/kafka"
	"github.com/soukengo/gopkg/infra/storage/mongo"
	"github.com/soukengo/gopkg/infra/storage/redis"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service"
	"rockimserver/app/logic/message/biz/consts"
	"rockimserver/conf"
)

const (
	configName = "message.yaml"
)

type Config struct {
	Global   *conf.Global
	Log      *log.Config
	Server   *server.Config
	Database *database.Config
	Cache    *cache.Config
	Queue    *queue.Config
	Lock     *lock.Config
}

func Load() (cfg *Config, err error) {
	global, err := conf.Load(service.AppMessage)
	if err != nil {
		return
	}
	cfg = defaultConfig()
	source := config.NewEnvSource(global.Config, configName)
	defer source.Close()
	loader := config.NewLoader(source)
	err = loader.Load(cfg)
	if err != nil {
		return
	}
	cfg.Global = global
	return
}

func defaultConfig() *Config {
	return &Config{
		Log: log.Default(),
		Server: &server.Config{
			Job: &job.Config{
				Queue: &queue.Config{
					Delayed: &queue.DelayedConfig{
						Redis: &redisqueue.Config{
							Reference: redis.Reference{Key: storage.DefaultKey},
							Consumer: &redisqueue.ConsumerConfig{
								Topics:  []queue.Topic{consts.QueueMessageDelivery},
								Workers: 10,
							},
						},
					},
				},
			},
		},
		Database: &database.Config{
			Mongodb: &mongo.Reference{Key: storage.DefaultKey},
		},
		Cache: &cache.Config{
			Redis: &redis.Reference{Key: storage.DefaultKey},
		},
		Queue: &queue.Config{
			General: &queue.GeneralConfig{
				Kafka: &kafkaqueue.Config{
					Reference: kafka.Reference{Key: storage.DefaultKey},
				},
			},
			Delayed: &queue.DelayedConfig{
				Redis: &redisqueue.Config{
					Reference: redis.Reference{Key: storage.DefaultKey},
					Consumer: &redisqueue.ConsumerConfig{
						Topics: []queue.Topic{consts.QueueMessageDelivery},
					},
				},
			},
		},
		Lock: &lock.Config{
			Redis: &redis.Reference{Key: storage.DefaultKey},
		},
	}
}

func (c *Config) Parse() (err error) {
	c.Server.Job.Parse(c.Global.Storage)
	c.Database.Parse(c.Global.Storage)
	c.Cache.Parse(c.Global.Storage)
	c.Queue.Parse(c.Global.Storage)
	c.Lock.Parse(c.Global.Storage)
	return
}
