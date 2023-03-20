package conf

import (
	"rockimserver/apis/rockim/service"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/conf"
	"rockimserver/pkg/component/config"
	"rockimserver/pkg/component/mq"
	"rockimserver/pkg/component/server"
	"rockimserver/pkg/component/server/job"
	"rockimserver/pkg/log"
	"time"
)

const (
	configName = "job.yaml"
)

func Load() (cfg *Config, err error) {
	global, err := conf.Load(service.AppJob)
	if err != nil {
		return
	}
	cfg = &Config{
		Log:   log.Default,
		Comet: &Comet{RoutineSize: 32, RoutineChan: 1024},
	}
	source := config.NewEnvSource(global.Config, configName)
	defer source.Close()
	loader := config.NewLoader(source)
	err = loader.Load(cfg)
	if err != nil {
		return
	}
	cfg.Global = global
	//  这里写死，不使用配置文件的内容
	cfg.Server = &server.Config{
		Job: &job.Config{
			GroupId: service.AppJob,
			Topics:  []string{enums.MQ_MESSAGE_PUSH.String()},
		},
	}
	return
}

type Config struct {
	Global *conf.Global
	Server *server.Config
	Log    *log.Config
	Comet  *Comet
	MQ     *mq.Config
}

// Comet is comet config.
type Comet struct {
	RoutineChan int
	RoutineSize int
	Group       *Group
}

type Group struct {
	Batch  int
	Signal time.Duration
	Idle   time.Duration
}
