package data

import (
	"github.com/google/wire"
	"rockimserver/app/logic/session/data/cache"
)

// ProviderSet is db providers.
var ProviderSet = wire.NewSet(
	cache.ProviderSet,
	NewChannelRepo,
	NewChannelQueryRepo,
)
