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
	configName = "logic-platform.yaml"
)

func Load() (cfg *Config, err error) {
	global, err := conf.Load(service.AppPlatform)
	if err != nil {
		return
	}
	cfg = &Config{
		Log: &log.Config{
			LoggerConfig: log.LoggerConfig{
				Level: "info",
			},
			Loggers: []log.LoggerConfig{
				{Name: "mongo", Level: "info"},
			},
		},
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
