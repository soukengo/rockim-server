package data

import (
	"github.com/google/wire"
	"rockimserver/app/logic/group/data/cache"
	"rockimserver/app/logic/group/data/database"
	"rockimserver/app/logic/group/data/grpc"
)

// ProviderSet is db providers.
var ProviderSet = wire.NewSet(
	database.ProviderSet,
	cache.ProviderSet,
	grpc.ProviderSet,
	NewGroupRepo,
	NewGroupMemberRepo,
	NewChatRoomMemberRepo,
	NewCometRepo,
)
