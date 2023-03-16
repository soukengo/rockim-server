package conf

import (
	"rockimserver/apis/rockim/service"
	"rockimserver/conf"
	"rockimserver/pkg/component/auth"
	"rockimserver/pkg/component/config"
	"rockimserver/pkg/component/database"
	"rockimserver/pkg/component/server"
	"rockimserver/pkg/log"
)

const (
	configName = "admin.yaml"
)

func Load() (cfg *Config, err error) {
	global, err := conf.Load(service.AppAdmin)
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
	Log      *log.Config
	Server   *server.Config
	Auth     *auth.Config
	Database *database.Config
}
