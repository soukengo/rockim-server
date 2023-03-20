package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/user/v1"
	"rockimserver/pkg/component/discovery"
)

func NewOnlineAPIClient(r registry.Discovery) (v1.OnlineAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppUser, r)
	if err != nil {
		return nil, err
	}
	return v1.NewOnlineAPIClient(conn), nil
}
