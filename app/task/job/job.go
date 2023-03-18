package job

import (
	"github.com/go-kratos/kratos/v2"
	"rockimserver"
	"rockimserver/app/task/job/conf"
	"rockimserver/pkg/component/server/job"
	"rockimserver/pkg/log"
)

// New new a new Job Application
func New(version string) (app *kratos.App, logger log.Logger, err error) {
	cfg, err := conf.Load()
	if err != nil {
		return
	}
	version = rockimserver.SetVersion(version)
	cfg.Global.Version = version
	err = configure(cfg)
	if err != nil {
		return
	}
	logger, err = log.Configure(cfg.Log, log.AppInfo(cfg.Global.AppId, cfg.Global.Version))
	if err != nil {
		return
	}
	app, err = wireApp(logger, cfg, cfg.Global.Discovery, cfg.Server, cfg.MQ)
	if err != nil {
		//err = errors.New(errors.UnknownCode, "", "wireApp error")
		return
	}
	return
}

func configure(cfg *conf.Config) (err error) {
	return
}
func newApp(logger log.Logger, cfg *conf.Config, js job.Server) *kratos.App {
	return kratos.New(
		kratos.Name(cfg.Global.AppId),
		kratos.Version(cfg.Global.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			js,
		),
	)
}
