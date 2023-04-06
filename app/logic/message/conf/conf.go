package conf

import (
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/config"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/mq"
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service"
	"rockimserver/conf"
)

const (
	configName = "message.yaml"
)

func Load() (cfg *Config, err error) {
	global, err := conf.Load(service.AppMessage)
	if err != nil {
		return
	}
	cfg = &Config{
		Log:   log.Default,
		Queue: &Queue{Workers: 10},
	}
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

type Config struct {
	Global   *conf.Global
	Server   *server.Config
	Log      *log.Config
	Database *database.Config
	Cache    *cache.Config
	MQ       *mq.Config
	Queue    *Queue
}

type Queue struct {
	Workers int
}
