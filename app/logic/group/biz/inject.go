package biz

import (
	"github.com/google/wire"
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/component/idgen"
	"rockimserver/pkg/component/lock"
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
