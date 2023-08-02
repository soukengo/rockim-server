package server

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/discovery"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	discovery.NewRegistrar,
	NewGRPCServer,
	NewServiceRegistry,
)
