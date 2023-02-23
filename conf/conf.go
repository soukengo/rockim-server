package conf

import (
	"flag"
	"os"
	"rockimserver/pkg/component/config"
	"rockimserver/pkg/component/discovery"
)

var (
	configPath = ""
	envKey     = "ROCKIM_ENV"
)

func init() {
	flag.StringVar(&configPath, "conf", "conf/rockim.yaml", "config path, eg: -conf rockim.yaml")

}

type Global struct {
	AppId     string
	Version   string
	Env       string
	Config    *config.Config
	Discovery *discovery.Config
}

func Load(appId string) (cfg *Global, err error) {
	cfg = &Global{AppId: appId}
	source := config.NewFileSource(configPath)
	defer source.Close()
	loader := config.NewLoader(source)
	err = loader.Load(cfg)
	if err != nil {
		return
	}
	// 从环境变量里读取，覆盖配置文件的参数
	var env = os.Getenv(envKey)
	if len(env) > 0 {
		cfg.Env = env
	}
	if cfg.Discovery != nil {
		cfg.Discovery.AppId = appId
	}
	cfg.Config.Env = cfg.Env
	return
}
