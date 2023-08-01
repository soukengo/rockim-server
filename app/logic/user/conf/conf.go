package conf

import (
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/config"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/lock"
	"github.com/soukengo/gopkg/component/transport"
	"github.com/soukengo/gopkg/infra/storage"
	"github.com/soukengo/gopkg/infra/storage/mongo"
	"github.com/soukengo/gopkg/infra/storage/redis"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service"
	"rockimserver/conf"
)

const (
	configName = "user.yaml"
)

func Load() (cfg *Config, err error) {
	global, err := conf.Load(service.AppUser)
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
	cfg.Log.Split = false
	return
}

func defaultConfig() *Config {
	return &Config{
		Log: log.Default(),
		Server: &Server{
			Grpc: &transport.Grpc{Addr: ":6101"},
		},
		Database: &database.Config{
			Mongodb: &mongo.Reference{Key: storage.DefaultKey},
		},
		Cache: &cache.Config{
			Redis: &redis.Reference{Key: storage.DefaultKey},
		},
		Lock: &lock.Config{
			Redis: &redis.Reference{Key: storage.DefaultKey},
		},
	}
}

type Config struct {
	Global   *conf.Global
	Server   *Server
	Log      *log.Config
	Database *database.Config
	Cache    *cache.Config
	Lock     *lock.Config
}

type Server struct {
	Grpc *transport.Grpc
}

func (c *Config) Parse() (err error) {
	c.Database.Parse(c.Global.Storage)
	c.Cache.Parse(c.Global.Storage)
	c.Lock.Parse(c.Global.Storage)
	return
}
