package conf

import (
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/config"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service"
	"rockimserver/conf"
)

const (
	configName = "group.yaml"
)

func Load() (cfg *Config, err error) {
	global, err := conf.Load(service.AppGroup)
	if err != nil {
		return
	}
	cfg = &Config{
		Log: log.Default,
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
}
