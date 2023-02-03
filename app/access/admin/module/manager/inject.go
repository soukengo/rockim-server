package manager

import (
	"github.com/google/wire"
	"rockimserver/app/access/admin/module/manager/biz"
	"rockimserver/app/access/admin/module/manager/data"
	"rockimserver/app/access/admin/module/manager/service"
)

var ProviderSet = wire.NewSet(data.ProviderSet, biz.ProviderSet, service.ProviderSet)
