package conf

import (
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/config"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/lock"
	"github.com/soukengo/gopkg/component/server"
	"github.com/soukengo/gopkg/infra/storage"
	"github.com/soukengo/gopkg/infra/storage/mongo"
	"github.com/soukengo/gopkg/infra/storage/redis"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service"
	"rockimserver/conf"
)

const (
	configName = "group.yaml"
)

func Load() (cfg *Config, err error) {
	global, err := conf.Load(service.AppGroup)
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
	Server   *server.Config
	Log      *log.Config
	Database *database.Config
	Cache    *cache.Config
	Lock     *lock.Config
}

func defaultConfig() *Config {
	return &Config{
		Log:    log.Default(),
		Server: &server.Config{},
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

func (c *Config) Parse() (err error) {
	c.Database.Parse(c.Global.Storage)
	c.Cache.Parse(c.Global.Storage)
	c.Lock.Parse(c.Global.Storage)
	return
}
