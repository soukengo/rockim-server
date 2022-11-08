package conf

import (
	"flag"
	"rockim/api"
	"rockim/pkg/component/config"
	"rockim/pkg/component/database/mongo"
	"rockim/pkg/component/discovery"
	"rockim/pkg/log"
	"time"
)

var (
	configPath = ""
)

func init() {
	flag.StringVar(&configPath, "conf", "configs/logic/user/user.yaml", "config path, eg: -conf config.yaml")
}

func Load() (conf *Config, err error) {
	conf = &Config{
		Env: &Env{AppId: api.AppUser},
		Log: &log.Config{
			LoggerConfig: log.LoggerConfig{
				Level: "info",
			},
			Loggers: []log.LoggerConfig{
				{Name: "mongo", Level: "info"},
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
	Server    *Server
	Discovery *discovery.Config
	Log       *log.Config
	Database  *Database
}

type Env struct {
	AppId   string
	Env     string
	Version string
}

type Server struct {
	Grpc *Grpc
}

type Grpc struct {
	Network string
	Addr    string
	Timeout time.Duration
}

type Database struct {
	Mongodb *mongo.Config
}
