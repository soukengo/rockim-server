package infra

import (
	"github.com/google/wire"
	"rockimserver/app/access/gateway/infra/grpc"
)

var ProviderSet = wire.NewSet(grpc.ProviderSet)
