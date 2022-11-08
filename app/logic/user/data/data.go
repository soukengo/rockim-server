package data

import (
	"github.com/google/wire"
	"rockim/app/logic/user/data/database"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(database.ProviderSet, NewUserRepo)
