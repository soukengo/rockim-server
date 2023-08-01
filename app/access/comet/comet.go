package comet

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/soukengo/gopkg/component/transport"
	"github.com/soukengo/gopkg/log"
	"rockimserver"
	"rockimserver/app/access/comet/conf"
)

// New new a new User Application
func New(version string) (app *kratos.App, logger log.Logger, err error) {
	cfg, err := conf.Load()
	if err != nil {
		return
	}
	version = rockimserver.SetVersion(version)
	cfg.Global.Version = version
	logger, err = log.Configure(cfg.Log, log.AppInfo(cfg.Global.AppId, cfg.Global.Version))
	if err != nil {
		return
	}
	app, err = wireApp(logger, cfg, cfg.Global.Discovery, cfg.Server, cfg.Protocol)
	if err != nil {
		//err = errors.New(errors.UnknownCode, "", "wireApp error")
		return
	}
	return
}

func newApp(cfg *conf.Config, group transport.ServerGroup, registrar registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(cfg.Global.ID),
		kratos.Name(cfg.Global.AppId),
		kratos.Version(cfg.Global.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Registrar(registrar),
		kratos.Server(
			group.Servers()...,
		),
	)
}
