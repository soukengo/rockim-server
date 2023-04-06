package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/soukengo/gopkg/component/discovery"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/platform/v1"
)

func NewProductClient(r registry.Discovery) (v1.ProductAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppPlatform, r)
	if err != nil {
		return nil, err
	}
	return v1.NewProductAPIClient(conn), nil
}
