package cache

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/database/redis"
)

var ProviderSet = wire.NewSet(redis.NewClient, NewProductData)
