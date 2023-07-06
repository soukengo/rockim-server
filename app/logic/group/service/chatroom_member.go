package service

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1"
	"rockimserver/app/logic/group/biz"
	"rockimserver/app/logic/group/biz/options"
)

type ChatRoomMemberService struct {
	uc *biz.ChatRoomMemberUseCase
	v1.UnimplementedChatRoomMemberAPIServer
}

func NewChatRoomMemberService(uc *biz.ChatRoomMemberUseCase) *ChatRoomMemberService {
	return &ChatRoomMemberService{uc: uc}
}

func (s *ChatRoomMemberService) Join(ctx context.Context, in *v1.ChatRoomJoinRequest) (out *v1.ChatRoomJoinResponse, err error) {
	err = s.uc.Join(ctx, &options.ChatRoomJoinOptions{
		ProductId: in.Base.ProductId,
		GroupId:   in.GroupId,
		Uid:       in.Uid,
	})
	if err != nil {
		return
	}
	out = &v1.ChatRoomJoinResponse{}
	return
}

func (s *ChatRoomMemberService) Quit(ctx context.Context, in *v1.ChatRoomQuitRequest) (out *v1.ChatRoomQuitResponse, err error) {
	err = s.uc.Quit(ctx, &options.ChatRoomQuitOptions{
		ProductId: in.Base.ProductId,
		GroupId:   in.GroupId,
		Uid:       in.Uid,
	})
	if err != nil {
		return
	}
	out = &v1.ChatRoomQuitResponse{}
	return
}
