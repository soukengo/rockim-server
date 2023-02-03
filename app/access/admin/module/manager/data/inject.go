package data

import (
	"github.com/google/wire"
	"rockimserver/app/access/admin/module/manager/data/database"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	database.ProviderSet,
	NewSysUserRepo,
	NewSysRoleRepo,
	NewSysResourceRepo,
	NewSysTenantResourceRepo,
	NewTenantRepo,
)
