package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/soukengo/gopkg/component/discovery"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/job/v1"
)

func NewCometAPIClient(r registry.Discovery) (v1.CometAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppJob, r)
	if err != nil {
		return nil, err
	}
	return v1.NewCometAPIClient(conn), nil
}
