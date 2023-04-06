package grpc

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/soukengo/gopkg/component/discovery"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/group/v1"
)

func NewGroupAPIClient(r registry.Discovery) (v1.GroupAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppGroup, r)
	if err != nil {
		return nil, err
	}
	return v1.NewGroupAPIClient(conn), nil
}
func NewChatRoomClient(r registry.Discovery) (v1.ChatRoomAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppGroup, r)
	if err != nil {
		return nil, err
	}
	return v1.NewChatRoomAPIClient(conn), nil
}
func NewChatRoomMemberClient(r registry.Discovery) (v1.ChatRoomMemberAPIClient, error) {
	conn, err := discovery.NewGrpcClient(service.AppGroup, r)
	if err != nil {
		return nil, err
	}
	return v1.NewChatRoomMemberAPIClient(conn), nil
}
