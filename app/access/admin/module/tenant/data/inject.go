package data

import (
	"github.com/google/wire"
	"rockim/app/access/admin/module/tenant/data/database"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	database.ProviderSet,
	NewSysTenantResourceRepo,
	NewTenantRepo,
	NewProductRepo,
)
