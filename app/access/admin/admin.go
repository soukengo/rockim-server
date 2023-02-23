package admin

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"rockimserver"
	"rockimserver/app/access/admin/conf"
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
	app, err = wireApp(cfg, cfg.Global.Discovery, cfg.Server, cfg.Auth, cfg.Database)
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
	return
}
func newApp(env *conf.Config, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.Name(env.Global.AppId),
		kratos.Version(env.Global.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			hs,
		),
	)
}
