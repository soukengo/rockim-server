package conf

import (
	"github.com/soukengo/gopkg/component/config"
	"github.com/soukengo/gopkg/component/queue"
	kafkaqueue "github.com/soukengo/gopkg/component/queue/core/kafka"
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/component/server/job"
	"github.com/soukengo/gopkg/infra/storage"
	"github.com/soukengo/gopkg/infra/storage/kafka"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service"
	"rockimserver/conf"
	"time"
)

const (
	configName = "job.yaml"
)

func Load() (cfg *Config, err error) {
	global, err := conf.Load(service.AppJob)
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
	//  这里写死，不使用配置文件的内容
	return
}

func defaultConfig() *Config {
	return &Config{
		Log: log.Default(),
		Comet: &Comet{RoutineSize: 32, RoutineChan: 1024, Room: &Room{
			Batch:  20,
			Signal: time.Second,
			Idle:   time.Minute * 15,
		}},
		Server: &server.Config{
			Job: &job.Config{
				Queue: &queue.Config{
					General: &queue.GeneralConfig{
						Kafka: &kafkaqueue.Config{
							Reference: kafka.Reference{Key: storage.DefaultKey},
							Consumer: &kafkaqueue.ConsumerConfig{
								GroupId: service.AppJob,
								Topics:  []queue.Topic{queue.Topic(service.MQ_MESSAGE_PUSH.String())},
								Workers: 10,
							},
						},
					},
				},
			},
		},
	}
}

type Config struct {
	Global *conf.Global
	Server *server.Config
	Log    *log.Config
	Comet  *Comet
}

// Comet is comet config.
type Comet struct {
	RoutineChan int
	RoutineSize int
	Room        *Room
}

type Room struct {
	Batch  int
	Signal time.Duration
	Idle   time.Duration
}

func (c *Config) Parse() (err error) {
	c.Server.Job.Parse(c.Global.Storage)
	return
}
