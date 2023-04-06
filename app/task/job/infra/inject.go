package infra

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/discovery"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	discovery.NewDiscovery,
)
