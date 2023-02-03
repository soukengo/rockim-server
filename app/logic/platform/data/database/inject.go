package database

import (
	"github.com/google/wire"
	"rockimserver/pkg/component/database/mongo"
)

var ProviderSet = wire.NewSet(
	mongo.NewClient,
	NewTenantData,
	NewProductData)
