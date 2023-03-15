package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/group/v1"
	"rockimserver/pkg/component/discovery"
)

func NewGroupAPIClient(r registry.Discovery) (v1.GroupAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppGroup, r)
	if err != nil {
		return nil, err
	}
	return v1.NewGroupAPIClient(conn), nil
}
