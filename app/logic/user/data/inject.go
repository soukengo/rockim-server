package data

import (
	"github.com/google/wire"
	"rockimserver/app/logic/user/data/cache"
	"rockimserver/app/logic/user/data/database"
)

// ProviderSet is db providers.
var ProviderSet = wire.NewSet(database.ProviderSet, cache.ProviderSet, NewUserRepo, NewAuthCodeRepo, NewAccessTokenRepo)
