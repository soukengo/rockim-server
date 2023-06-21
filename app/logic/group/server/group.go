package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "rockimserver/apis/rockim/service/group/v1"
	"rockimserver/app/logic/group/service"
)

type ServiceGroup struct {
	groupSrv          *service.GroupService
	chatRoomSrv       *service.ChatRoomService
	chatRoomMemberSrv *service.ChatRoomMemberService
}

func NewServiceGroup(groupSrv *service.GroupService, chatRoomSrv *service.ChatRoomService, chatRoomMemberSrv *service.ChatRoomMemberService) *ServiceGroup {
	return &ServiceGroup{groupSrv: groupSrv, chatRoomSrv: chatRoomSrv, chatRoomMemberSrv: chatRoomMemberSrv}
}

func (g *ServiceGroup) Register(srv *grpc.Server) {
	v1.RegisterGroupAPIServer(srv, g.groupSrv)
	v1.RegisterChatRoomAPIServer(srv, g.chatRoomSrv)
	v1.RegisterChatRoomMemberAPIServer(srv, g.chatRoomMemberSrv)
}
