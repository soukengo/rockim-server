package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/discovery"
	"rockimserver/apis/rockim/service"
	"rockimserver/apis/rockim/service/platform/v1"
)

// ProviderSet is grpc providers.
var ProviderSet = wire.NewSet(
	discovery.NewDiscovery,
	NewTenantAPIClient,
	NewProductAPIClient,
)

func NewTenantAPIClient(r registry.Discovery) (v1.TenantAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppPlatform, r)
	if err != nil {
		return nil, err
	}
	return v1.NewTenantAPIClient(conn), nil
}
func NewProductAPIClient(r registry.Discovery) (v1.ProductAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppPlatform, r)
	if err != nil {
		return nil, err
	}
	return v1.NewProductAPIClient(conn), nil
}
