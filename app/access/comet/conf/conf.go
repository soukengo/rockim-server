package conf

import (
	"rockimserver/apis/rockim/service"
	"rockimserver/conf"
	"rockimserver/pkg/component/config"
	"rockimserver/pkg/component/server"
	"rockimserver/pkg/log"
	"time"
)

const (
	configName = "access-comet.yaml"
)

func Load() (cfg *Config, err error) {
	global, err := conf.Load(service.AppComet)
	if err != nil {
		return
	}
	cfg = &Config{
		Log: log.Default,
		Protocol: &Protocol{
			HandshakeTimeout:  time.Second * 10,
			HeartbeatInterval: time.Minute * 3,
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
	Protocol *Protocol
}

type Protocol struct {
	HandshakeTimeout  time.Duration
	HeartbeatInterval time.Duration
}
