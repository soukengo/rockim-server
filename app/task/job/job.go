package job

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/soukengo/gopkg/component/transport"
	"github.com/soukengo/gopkg/log"
	"rockimserver"
	"rockimserver/app/task/job/conf"
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
	app, err = wireApp(logger, cfg, cfg.Global.Discovery, cfg.Server)
	if err != nil {
		return
	}
	return
}

func configure(cfg *conf.Config) (err error) {
	err = cfg.Parse()
	if err != nil {
		return
	}
	return
}

func newApp(cfg *conf.Config, group transport.ServerGroup, registrar registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(cfg.Global.AppId),
		kratos.Version(cfg.Global.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Registrar(registrar),
		kratos.Server(
			group.Servers()...,
		),
	)
}
