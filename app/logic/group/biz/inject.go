package biz

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/idgen"
	"github.com/soukengo/gopkg/component/lock"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/group/conf"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewLockBuilder,
	idgen.NewMongoGenerator,

	NewGroupUseCase,
	NewGroupMemberUseCase,
	NewChatRoomMemberManager,
	NewChatRoomUseCase,
	NewChatRoomMemberUseCase,
	NewCometUseCase,
)

func NewLockBuilder(cfg *conf.Config, logger log.Logger) lock.Builder {
	return lock.NewBuilder(cfg.Lock, logger)
}
