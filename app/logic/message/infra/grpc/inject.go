package grpc

import (
	"github.com/google/wire"
)

// ProviderSet is grpc providers.
var ProviderSet = wire.NewSet(
	NewUserAPIClient,
	NewGroupAPIClient,
	NewGroupMemberAPIClient,
	NewCometAPIClient,
)
