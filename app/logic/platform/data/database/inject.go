package database

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/database/mongo"
)

var ProviderSet = wire.NewSet(
	mongo.NewClient,
	NewTenantData,
	NewProductData)
