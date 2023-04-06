package database

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/database"
)

var ProviderSet = wire.NewSet(
	database.NewMongoClient,
	NewUserData,
)
