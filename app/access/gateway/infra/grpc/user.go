package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/soukengo/gopkg/component/discovery"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/user/v1"
)

func NewUserClient(r registry.Discovery) (v1.UserAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppUser, r)
	if err != nil {
		return nil, err
	}
	return v1.NewUserAPIClient(conn), nil
}
func NewAuthClient(r registry.Discovery) (v1.AuthAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppUser, r)
	if err != nil {
		return nil, err
	}
	return v1.NewAuthAPIClient(conn), nil
}
