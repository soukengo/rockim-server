package conf

import (
	"github.com/soukengo/gopkg/component/auth"
	"github.com/soukengo/gopkg/component/config"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service"
	"rockimserver/conf"
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
