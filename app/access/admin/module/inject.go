package module

import (
	"github.com/google/wire"
	"rockimserver/app/access/admin/module/manager"
	"rockimserver/app/access/admin/module/tenant"
)

var ProviderSet = wire.NewSet(manager.ProviderSet, tenant.ProviderSet)
