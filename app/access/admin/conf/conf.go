package conf

import (
	"flag"
	"rockimserver/apis/rockim/service"
	"rockimserver/pkg/component/config"
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/discovery"
	"rockimserver/pkg/log"
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
	Log       *log.Config
	Server    *Server
	Auth      *Auth
	Database  *Database
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

type Auth struct {
	Jwt *Jwt
}

type Jwt struct {
	AppKey  string
	Expires time.Duration
}

type Database struct {
	Mongodb *mongo.Config
}
