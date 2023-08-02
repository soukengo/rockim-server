package cache

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/user/conf"
)

var ProviderSet = wire.NewSet(
	NewCacheManager,
	NewUserData,
	NewAuthCodeData,
	NewAccessTokenData,
)

func NewCacheManager(cfg *conf.Config, logger log.Logger) *cache.Manager {
	return cache.NewManager(cfg.Cache, logger)
}
