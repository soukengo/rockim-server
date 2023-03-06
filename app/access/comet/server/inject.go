package server

import (
	"github.com/google/wire"
	"rockimserver/app/access/comet/server/grpc"
	"rockimserver/app/access/comet/server/socket"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(socket.ProviderSet, grpc.ProviderSet)
