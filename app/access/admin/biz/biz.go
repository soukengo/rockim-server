package biz

import (
	"github.com/google/wire"
	"rockim/app/access/admin/biz/manager"
	"rockim/app/access/admin/biz/tenant"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(manager.ProviderSet, tenant.ProviderSet)
