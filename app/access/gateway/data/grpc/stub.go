package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"rockim/api"
	userV1 "rockim/api/logic/user/v1"
	"rockim/pkg/component/discovery"
)

// ProviderSet is grpc providers.
var ProviderSet = wire.NewSet(discovery.NewDiscovery, NewUserClient)

func NewUserClient(r registry.Discovery) (userV1.UserAPIClient, error) {
	conn, err := api.NewGrpcClient(api.AppUser, r)
	if err != nil {
		return nil, err
	}
	return userV1.NewUserAPIClient(conn), nil
}
