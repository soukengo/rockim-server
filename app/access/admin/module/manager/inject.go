package manager

import (
	"github.com/google/wire"
	"rockim/app/access/admin/module/manager/biz"
	"rockim/app/access/admin/module/manager/data"
	"rockim/app/access/admin/module/manager/service"
)

var ProviderSet = wire.NewSet(data.ProviderSet, biz.ProviderSet, service.ProviderSet)
