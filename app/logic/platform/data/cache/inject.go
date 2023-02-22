package cache

import (
	"github.com/google/wire"
	"rockimserver/pkg/component/database/redis"
)

var ProviderSet = wire.NewSet(redis.NewClient, NewProductData)
