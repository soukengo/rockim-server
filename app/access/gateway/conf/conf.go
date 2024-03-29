package conf

import (
	"github.com/soukengo/gopkg/component/config"
	"github.com/soukengo/gopkg/component/transport"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service"
	"rockimserver/conf"
)

const (
	configName = "gateway.yaml"
)

func Load() (cfg *Config, err error) {
	global, err := conf.Load(service.AppGateway)
	if err != nil {
		return
	}
	cfg = &Config{
		Log: log.Default(),
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
	Global *conf.Global
	Server *Server
	Log    *log.Config
	Comet  Comet
}

type Server struct {
	Http *transport.Http
}

type Comet struct {
	TCP       TCP
	WebSocket WebSocket
}

type TCP struct {
	Addr string
}

type WebSocket struct {
	Addr string
}
