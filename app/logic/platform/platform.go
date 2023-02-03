package platform

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"rockimserver"
	"rockimserver/app/logic/platform/conf"
	"rockimserver/pkg/log"
)

// New a new Platform Application
func New(version string) (app *kratos.App, err error) {
	cfg, err := conf.Load()
	if err != nil {
		return
	}
	version = rockim.SetVersion(version)
	cfg.Env.Version = version
	err = configure(cfg)
	if err != nil {
		return
	}
	app, err = wireApp(cfg.Env, cfg.Discovery, cfg.Server, cfg.Database.Mongodb)
	if err != nil {
		//err = errors.New(errors.UnknownCode, "", "wireApp error")
		return
	}
	return
}

func configure(cfg *conf.Config) (err error) {
	cfg.Log.AppId = cfg.Env.AppId
	cfg.Log.AppVersion = cfg.Env.Version
	err = log.Configure(cfg.Log)
	if err != nil {
		return
	}
	return
}
func newApp(env *conf.Env, gs *grpc.Server, registrar registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(env.AppId),
		kratos.Version(env.Version),
		kratos.Registrar(registrar),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			gs,
		),
	)
}
