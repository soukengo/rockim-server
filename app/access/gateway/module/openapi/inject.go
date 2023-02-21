package openapi

import (
	"github.com/google/wire"
	"rockimserver/app/access/gateway/module/openapi/biz"
	"rockimserver/app/access/gateway/module/openapi/data"
	"rockimserver/app/access/gateway/module/openapi/service"
)

var ProviderSet = wire.NewSet(data.ProviderSet, biz.ProviderSet, service.ProviderSet)
