package client

import (
	"github.com/google/wire"
	"rockimserver/app/access/gateway/module/client/biz"
	"rockimserver/app/access/gateway/module/client/data"
	"rockimserver/app/access/gateway/module/client/service"
)

var ProviderSet = wire.NewSet(data.ProviderSet, biz.ProviderSet, service.ProviderSet)
