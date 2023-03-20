package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/message/v1"
	"rockimserver/pkg/component/discovery"
)

func NewMessageAPIClient(r registry.Discovery) (v1.MessageAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppMessage, r)
	if err != nil {
		return nil, err
	}
	return v1.NewMessageAPIClient(conn), nil
}
