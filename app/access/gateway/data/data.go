package data

import (
	"github.com/google/wire"
	"rockimserver/app/access/gateway/data/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(grpc.ProviderSet, NewUserRepo)
