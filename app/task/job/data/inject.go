package data

import (
	"github.com/google/wire"
	"rockimserver/app/task/job/data/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	grpc.ProviderSet,
	NewPushRepo,
)
