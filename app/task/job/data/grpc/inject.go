package grpc

import (
	"github.com/google/wire"
	"rockimserver/app/task/job/data/grpc/comet"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	comet.NewCometManager,
	NewChannelQueryAPIClient,
)
