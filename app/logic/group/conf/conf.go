package conf

import (
	"rockimserver/apis/rockim/service"
	"rockimserver/conf"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/config"
	"rockimserver/pkg/component/database"
	"rockimserver/pkg/component/server"
	"rockimserver/pkg/log"
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
