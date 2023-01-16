package data

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewSysUserRepo,
	NewSysRoleRepo,
	NewSysResourceRepo,
	NewSysTenantResourceRepo,
	NewTenantRepo,
)
