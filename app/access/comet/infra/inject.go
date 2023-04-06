package infra

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/discovery"
	"rockimserver/app/access/comet/infra/grpc"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	discovery.NewRegistrar,
	discovery.NewDiscovery,
	grpc.ProviderSet,
)
