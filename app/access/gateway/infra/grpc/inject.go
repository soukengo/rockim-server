package grpc

import (
	"github.com/google/wire"
	"rockimserver/pkg/component/discovery"
)

// ProviderSet is grpc providers.
var ProviderSet = wire.NewSet(discovery.NewDiscovery, NewProductClient, NewUserClient, NewAuthClient)
