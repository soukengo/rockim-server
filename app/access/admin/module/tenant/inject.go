package tenant

import (
	"github.com/google/wire"
	"rockimserver/app/access/admin/module/tenant/biz"
	"rockimserver/app/access/admin/module/tenant/data"
	"rockimserver/app/access/admin/module/tenant/service"
)

var ProviderSet = wire.NewSet(biz.ProviderSet, service.ProviderSet, data.ProviderSet)
