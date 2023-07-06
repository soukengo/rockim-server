package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/soukengo/gopkg/component/discovery"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/group/v1"
)

func NewGroupMemberAPIClient(r registry.Discovery) (v1.GroupMemberAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppGroup, r)
	if err != nil {
		return nil, err
	}
	return v1.NewGroupMemberAPIClient(conn), nil
}
