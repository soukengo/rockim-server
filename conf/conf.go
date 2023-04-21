package conf

import (
	"flag"
	"fmt"
	"github.com/soukengo/gopkg/component/config"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/infra/storage"
	"github.com/soukengo/gopkg/util/ip"
	"os"
	"time"
)

var (
	configPath = ""
	envKey     = "ROCKIM_ENV"
)

func init() {
	flag.StringVar(&configPath, "conf", "conf/rockim.yaml", "config path, eg: -conf rockim.yaml")

}

type Global struct {
	ID        string
	AppId     string
	Version   string
	Env       string
	Config    *config.Config
	Discovery *discovery.Config
	Storage   *storage.Config
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
	cfg.ID = fmt.Sprintf("%s:%d", ip.InternalIP(), time.Now().UnixMilli())
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
