package conf

import (
	"math/rand"
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
			HandshakeTimeout:           time.Second * 10,
			MinServerHeartbeatInterval: time.Minute * 2,
			MaxServerHeartbeatInterval: time.Minute * 4,
			MaxClientHeartbeatInterval: time.Minute * 2,
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
	HandshakeTimeout           time.Duration // 长连接握手超时时间
	MinServerHeartbeatInterval time.Duration // comet服务向后端服务发起心跳的最小时间间隔（心跳降频）
	MaxServerHeartbeatInterval time.Duration // comet服务向后端服务发起心跳的最大时间间隔（心跳降频）
	MaxClientHeartbeatInterval time.Duration // 客户端向comet服务发起心跳的最大时间间隔（超时则断开）
}

func (p *Protocol) RandomServerHeartbeatInterval() time.Duration {
	return p.MinServerHeartbeatInterval + time.Duration(rand.Int63n(int64(p.MaxServerHeartbeatInterval)))
}
