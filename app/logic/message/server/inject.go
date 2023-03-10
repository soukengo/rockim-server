package server

import (
	"github.com/google/wire"
	"rockimserver/pkg/component/discovery"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(discovery.NewRegistrar, NewGRPCServer, NewServiceGroup)
