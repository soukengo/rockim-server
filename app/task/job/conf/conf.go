package conf

import (
	"github.com/soukengo/gopkg/component/config"
	"github.com/soukengo/gopkg/component/transport"
	"github.com/soukengo/gopkg/component/transport/queue"
	kafkaqueue "github.com/soukengo/gopkg/component/transport/queue/provider/kafka"
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
		Server: &Server{
			Queue: &queue.GeneralConfig{
				Kafka: &kafkaqueue.Config{
					Reference: kafka.Reference{Key: storage.DefaultKey},
					Consumer: &kafka.ConsumerConfig{
						GroupId: service.AppJob,
						Workers: 10,
					},
				},
			},
			Grpc: &transport.Grpc{
				Addr:    ":6108",
				Timeout: time.Second * 10,
			},
		},
	}
}

type Config struct {
	Global *conf.Global
	Server *Server
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

type Server struct {
	Queue *queue.GeneralConfig
	Grpc  *transport.Grpc
}

func (c *Config) Parse() (err error) {
	c.Server.Queue.Parse(c.Global.Storage)
	return
}
