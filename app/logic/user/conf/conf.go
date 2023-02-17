package conf

import (
	"flag"
	"rockimserver/apis/rockim/service"
	"rockimserver/pkg/component/config"
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/component/discovery"
	"rockimserver/pkg/log"
	"time"
)

var (
	configPath = ""
)

func init() {
	flag.StringVar(&configPath, "conf", "config/logic/user.yaml", "config path, eg: -conf config.yaml")
}

func Load() (conf *Config, err error) {
	conf = &Config{
		Env: &Env{AppId: service.AppUser},
		Log: &log.Config{
			LoggerConfig: log.LoggerConfig{
				Level: "info",
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
	Redis   *redis.Config
}
