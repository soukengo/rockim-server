package gateway

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"rockimserver"
	"rockimserver/app/access/gateway/conf"
	"rockimserver/pkg/log"
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
	app, err = wireApp(logger, cfg, cfg.Global.Discovery, cfg.Server)
	if err != nil {
		//err = errors.New(errors.UnknownCode, "", "wireApp error")
		return
	}
	return
}

func newApp(logger log.Logger, env *conf.Config, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.Name(env.Global.AppId),
		kratos.Version(env.Global.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			hs,
		),
	)
}
