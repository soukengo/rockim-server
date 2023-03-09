package data

import (
	"github.com/google/wire"
	"rockimserver/app/access/comet/module/client/data/grpc"
)

var ProviderSet = wire.NewSet(grpc.ProviderSet, NewOnlineRepo)
