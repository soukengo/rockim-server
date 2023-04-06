package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/soukengo/gopkg/component/discovery"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/message/v1"
)

func NewMessageAPIClient(r registry.Discovery) (v1.MessageAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppMessage, r)
	if err != nil {
		return nil, err
	}
	return v1.NewMessageAPIClient(conn), nil
}
