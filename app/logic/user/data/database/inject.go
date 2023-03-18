package database

import (
	"github.com/google/wire"
	"rockimserver/pkg/component/database"
)

var ProviderSet = wire.NewSet(
	database.NewMongoClient,
	NewUserData,
)
