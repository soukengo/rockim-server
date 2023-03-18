package message

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"rockimserver"
	"rockimserver/app/logic/message/conf"
	"rockimserver/pkg/log"
)

// New new a new Message Application
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
	app, err = wireApp(logger, cfg, cfg.Global.Discovery, cfg.Server, cfg.Database, cfg.Cache, cfg.MQ)
	if err != nil {
		//err = errors.New(errors.UnknownCode, "", "wireApp error")
		return
	}
	return
}

func configure(cfg *conf.Config) (err error) {
	if err != nil {
		return
	}
	err = cfg.Cache.Parse()
	if err != nil {
		return
	}
	return
}
func newApp(logger log.Logger, cfg *conf.Config, gs *grpc.Server, registrar registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(cfg.Global.AppId),
		kratos.Version(cfg.Global.Version),
		kratos.Registrar(registrar),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			gs,
		),
	)
}
