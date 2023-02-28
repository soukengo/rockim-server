package service

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1"
	"rockimserver/app/logic/group/biz"
)

type ChatRoomMemberService struct {
	uc *biz.ChatRoomMemberUseCase
	v1.UnimplementedChatRoomMemberAPIServer
}

func NewChatRoomMemberService(uc *biz.ChatRoomMemberUseCase) *ChatRoomMemberService {
	return &ChatRoomMemberService{uc: uc}
}

func (s *ChatRoomMemberService) Join(ctx context.Context, request *v1.ChatRoomJoinRequest) (*v1.ChatRoomJoinResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChatRoomMemberService) Quit(ctx context.Context, request *v1.ChatRoomQuitRequest) (*v1.ChatRoomQuitResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChatRoomMemberService) IsMember(ctx context.Context, request *v1.ChatRoomMemberCheckRequest) (*v1.ChatRoomMemberCheckResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChatRoomMemberService) Find(ctx context.Context, request *v1.ChatRoomMemberFindRequest) (*v1.ChatRoomMemberFindResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChatRoomMemberService) List(ctx context.Context, request *v1.ChatRoomMemberCheckRequest) (*v1.ChatRoomMemberCheckResponse, error) {
	//TODO implement me
	panic("implement me")
}
