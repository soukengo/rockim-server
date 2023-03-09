package server

import (
	"github.com/google/wire"
	"rockimserver/app/access/comet/module/server/biz"
	"rockimserver/app/access/comet/module/server/service"
)

var ProviderSet = wire.NewSet(service.ProviderSet, biz.ProviderSet)
