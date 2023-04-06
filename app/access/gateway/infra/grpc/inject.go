package grpc

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/discovery"
)

// ProviderSet is grpc providers.
var ProviderSet = wire.NewSet(
	discovery.NewDiscovery,
	NewProductClient,
	NewUserClient,
	NewAuthClient,
	NewGroupAPIClient,
	NewChatRoomClient,
	NewChatRoomMemberClient,
	NewMessageAPIClient,
)
