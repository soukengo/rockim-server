package conf

import (
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/config"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/transport"
	"github.com/soukengo/gopkg/infra/storage"
	"github.com/soukengo/gopkg/infra/storage/mongo"
	"github.com/soukengo/gopkg/infra/storage/redis"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service"
	"rockimserver/conf"
)

const (
	configName = "platform.yaml"
)

func Load() (cfg *Config, err error) {
	global, err := conf.Load(service.AppPlatform)
	if err != nil {
		return
	}
	cfg = defaultConfig()
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
	Server   *Server
	Log      *log.Config
	Database *database.Config
	Cache    *cache.Config
}

type Server struct {
	Grpc *transport.Grpc
}

func defaultConfig() *Config {
	return &Config{
		Log: log.Default(),
		Server: &Server{
			Grpc: &transport.Grpc{Addr: ":6102"},
		},
		Database: &database.Config{
			Mongodb: &mongo.Reference{Key: storage.DefaultKey},
		},
		Cache: &cache.Config{
			Redis: &redis.Reference{Key: storage.DefaultKey},
		},
	}
}

func (c *Config) Parse() (err error) {
	c.Database.Parse(c.Global.Storage)
	c.Cache.Parse(c.Global.Storage)
	return
}
