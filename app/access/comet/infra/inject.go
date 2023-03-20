package infra

import (
	"github.com/google/wire"
	"rockimserver/app/access/comet/infra/grpc"
	"rockimserver/pkg/component/discovery"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	discovery.NewRegistrar,
	discovery.NewDiscovery,
	grpc.ProviderSet,
)
