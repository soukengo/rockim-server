package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/soukengo/gopkg/component/discovery"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/session/v1"
)

func NewChannelQueryAPIClient(r registry.Discovery) (v1.ChannelQueryAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppSession, r)
	if err != nil {
		return nil, err
	}
	return v1.NewChannelQueryAPIClient(conn), nil
}
