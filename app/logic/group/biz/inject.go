package biz

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/database/mongo"
	"github.com/soukengo/gopkg/component/database/redis"
	"github.com/soukengo/gopkg/component/idgen"
	"github.com/soukengo/gopkg/component/lock"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(componentSet,
	NewGroupUseCase,
	NewChatRoomMemberManager,
	NewChatRoomUseCase,
	NewChatRoomMemberUseCase,
)

// 注入需要使用的组件
var componentSet = wire.NewSet(
	mongo.NewClient,
	redis.NewClient,
	lock.NewRedisBuilder,
	idgen.NewMongoGenerator,
)
