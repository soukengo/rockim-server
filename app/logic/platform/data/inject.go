package data

import (
	"github.com/google/wire"
	"rockimserver/app/logic/platform/data/cache"
	"rockimserver/app/logic/platform/data/database"
)

// ProviderSet is db providers.
var ProviderSet = wire.NewSet(
	database.ProviderSet,
	cache.ProviderSet,
	NewTenantRepo,
	NewProductRepo,
)
