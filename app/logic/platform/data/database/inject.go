package database

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewTenantData,
	NewProductData)
