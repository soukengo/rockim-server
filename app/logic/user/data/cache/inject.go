package cache

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/database"
)

var ProviderSet = wire.NewSet(
	database.NewRedisClient,
	NewUserData,
	NewAuthCodeData,
	NewAccessTokenData,
	NewOnlineData,
)
