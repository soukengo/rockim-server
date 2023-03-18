package cache

import (
	"github.com/google/wire"
	"rockimserver/pkg/component/database"
)

var ProviderSet = wire.NewSet(
	database.NewRedisClient,
	NewUserData,
	NewAuthCodeData,
	NewAccessTokenData,
	NewOnlineData,
)
