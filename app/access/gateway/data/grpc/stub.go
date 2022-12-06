package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"rockim/api/rockim/service"
	"rockim/api/rockim/service/user/v1"
	"rockim/pkg/component/discovery"
)

// ProviderSet is grpc providers.
var ProviderSet = wire.NewSet(discovery.NewDiscovery, NewUserClient)

func NewUserClient(r registry.Discovery) (v1.UserAPIClient, error) {
	conn, err := service.NewGrpcClient(service.AppUser, r)
	if err != nil {
		return nil, err
	}
	return v1.NewUserAPIClient(conn), nil
}
