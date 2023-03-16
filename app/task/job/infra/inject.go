package infra

import (
	"github.com/google/wire"
	"rockimserver/pkg/component/discovery"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	discovery.NewDiscovery,
)
