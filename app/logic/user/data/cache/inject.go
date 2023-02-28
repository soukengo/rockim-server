package cache

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserData, NewAuthCodeData, NewAccessTokenData)
