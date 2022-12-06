package service

import (
	"github.com/google/wire"
	"rockim/app/access/admin/service/manager"
	"rockim/app/access/admin/service/tenant"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(tenant.ProviderSet, manager.ProviderSet)
