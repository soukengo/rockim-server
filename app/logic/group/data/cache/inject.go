package cache

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/group/conf"
)

var ProviderSet = wire.NewSet(
	NewCacheManager,
	NewGroupData,
	NewGroupMemberData,
	NewChatRoomMemberData,
)

func NewCacheManager(cfg *conf.Config, logger log.Logger) *cache.Manager {
	return cache.NewManager(cfg.Cache, logger)
}
