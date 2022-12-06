package conf

import (
	"flag"
	"rockim/api/rockim/service"
	"rockim/pkg/component/config"
	"rockim/pkg/component/discovery"
	"rockim/pkg/log"
	"time"
)

var (
	configPath = ""
)

func init() {
	flag.StringVar(&configPath, "conf", "config/access/admin.yaml", "config path, eg: -conf config.yaml")
}

func Load() (conf *Config, err error) {
	conf = &Config{
		Env: &Env{AppId: service.AppAdmin},
		Log: &log.Config{
			LoggerConfig: log.LoggerConfig{
				Level: "info",
				Split: true,
			},
		},
	}
	source := config.NewFileSource(configPath)
	defer source.Close()
	loader := config.NewLoader(source)
	err = loader.Load(conf)
	return
}

type Config struct {
	Env       *Env
	Discovery *discovery.Config
	Server    *Server
	Log       *log.Config
}

type Env struct {
	AppId   string
	Env     string
	Version string
}

type Server struct {
	Http *Http
}

type Http struct {
	Network string
	Addr    string
	Timeout time.Duration
}
