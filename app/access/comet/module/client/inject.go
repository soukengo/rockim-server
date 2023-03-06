package client

import (
	"github.com/google/wire"
	"rockimserver/app/access/comet/module/client/biz"
	"rockimserver/app/access/comet/module/client/data"
	"rockimserver/app/access/comet/module/client/service"
)

var ProviderSet = wire.NewSet(data.ProviderSet, biz.ProviderSet, service.ProviderSet)
