package user

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"rockimserver"
	"rockimserver/app/logic/user/conf"
	"rockimserver/pkg/log"
)

// New new a new User Application
func New(version string) (app *kratos.App, err error) {
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
	app, err = wireApp(cfg, cfg.Global.Discovery, cfg.Server, cfg.Database.Mongodb, cfg.Database.Redis, cfg.Cache)
	if err != nil {
		//err = errors.New(errors.UnknownCode, "", "wireApp error")
		return
	}
	return
}

func configure(cfg *conf.Config) (err error) {
	cfg.Log.AppId = cfg.Global.AppId
	cfg.Log.AppVersion = cfg.Global.Version
	err = log.Configure(cfg.Log)
	if err != nil {
		return
	}
	err = cfg.Cache.Parse()
	if err != nil {
		return
	}
	return
}
func newApp(cfg *conf.Config, gs *grpc.Server, registrar registry.Registrar) *kratos.App {
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
