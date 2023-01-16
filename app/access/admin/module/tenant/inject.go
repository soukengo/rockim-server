package tenant

import (
	"github.com/google/wire"
	"rockim/app/access/admin/module/tenant/biz"
	"rockim/app/access/admin/module/tenant/data"
	"rockim/app/access/admin/module/tenant/service"
)

var ProviderSet = wire.NewSet(biz.ProviderSet, service.ProviderSet, data.ProviderSet)
