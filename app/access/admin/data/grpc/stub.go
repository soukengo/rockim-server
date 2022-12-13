package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"rockim/api/rockim/service"
	"rockim/api/rockim/service/platform/v1"
	"rockim/pkg/component/discovery"
)

// ProviderSet is grpc providers.
var ProviderSet = wire.NewSet(discovery.NewDiscovery, NewPlatUserClient)

func NewPlatUserClient(r registry.Discovery) (v1.PlatUserAPIClient, error) {
	conn, err := service.NewGrpcClient(service.AppPlatform, r)
	if err != nil {
		return nil, err
	}
	return v1.NewPlatUserAPIClient(conn), nil
}
