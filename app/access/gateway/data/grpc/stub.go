package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"rockimserver/apis/rockim/service"
	"rockimserver/apis/rockim/service/user/v1"
	"rockimserver/pkg/component/discovery"
)

// ProviderSet is grpc providers.
var ProviderSet = wire.NewSet(discovery.NewDiscovery, NewUserClient, NewAuthClient)

func NewUserClient(r registry.Discovery) (v1.UserAPIClient, error) {
	conn, err := service.NewGrpcClient(service.AppUser, r)
	if err != nil {
		return nil, err
	}
	return v1.NewUserAPIClient(conn), nil
}
func NewAuthClient(r registry.Discovery) (v1.AuthAPIClient, error) {
	conn, err := service.NewGrpcClient(service.AppUser, r)
	if err != nil {
		return nil, err
	}
	return v1.NewAuthAPIClient(conn), nil
}
