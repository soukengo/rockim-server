package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"rockim/api/rockim/service"
	"rockim/api/rockim/service/platform/v1"
	"rockim/pkg/component/discovery"
)

// ProviderSet is grpc providers.
var ProviderSet = wire.NewSet(discovery.NewDiscovery, NewPlatUserClient, NewPlatRoleClient, NewPlatResourceClient, NewTenantAPIClient, NewTenantResourceAPIClient)

func NewPlatUserClient(r registry.Discovery) (v1.PlatUserAPIClient, error) {
	conn, err := service.NewGrpcClient(service.AppPlatform, r)
	if err != nil {
		return nil, err
	}
	return v1.NewPlatUserAPIClient(conn), nil
}
func NewPlatRoleClient(r registry.Discovery) (v1.PlatRoleAPIClient, error) {
	conn, err := service.NewGrpcClient(service.AppPlatform, r)
	if err != nil {
		return nil, err
	}
	return v1.NewPlatRoleAPIClient(conn), nil
}
func NewPlatResourceClient(r registry.Discovery) (v1.PlatResourceAPIClient, error) {
	conn, err := service.NewGrpcClient(service.AppPlatform, r)
	if err != nil {
		return nil, err
	}
	return v1.NewPlatResourceAPIClient(conn), nil
}
func NewTenantAPIClient(r registry.Discovery) (v1.TenantAPIClient, error) {
	conn, err := service.NewGrpcClient(service.AppPlatform, r)
	if err != nil {
		return nil, err
	}
	return v1.NewTenantAPIClient(conn), nil
}
func NewTenantResourceAPIClient(r registry.Discovery) (v1.TenantResourceAPIClient, error) {
	conn, err := service.NewGrpcClient(service.AppPlatform, r)
	if err != nil {
		return nil, err
	}
	return v1.NewTenantResourceAPIClient(conn), nil
}
