package data

import (
	"github.com/google/wire"
	"rockimserver/app/logic/message/data/cache"
	"rockimserver/app/logic/message/data/database"
	"rockimserver/app/logic/message/data/mq"
)

// ProviderSet is db providers.
var ProviderSet = wire.NewSet(
	database.ProviderSet,
	cache.ProviderSet,
	mq.ProviderSet,
	NewMessageRepo,
	NewUserRepo,
	NewGroupRepo,
	NewPushMessageRepo,
	NewMessageDeliveryRepo,
	NewMessageQueryRepo,
	NewMessageBoxRepo,
)
