package data

import (
	"github.com/google/wire"
	"rockim/app/logic/platform/data/database"
)

// ProviderSet is db providers.
var ProviderSet = wire.NewSet(database.ProviderSet, NewUserRepo)
