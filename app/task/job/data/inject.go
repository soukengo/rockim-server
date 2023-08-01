package data

import (
	"github.com/google/wire"
	"rockimserver/app/task/job/data/grpc"
	"rockimserver/app/task/job/data/mq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	grpc.ProviderSet,
	mq.ProviderSet,
	NewCometRepo,
	NewChannelRepo,
)
