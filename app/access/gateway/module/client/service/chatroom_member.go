package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/client/v1/protocol/http"
	"rockimserver/app/access/gateway/module/client/biz"
	"rockimserver/app/access/gateway/module/client/biz/options"
)

type ChatRoomMemberService struct {
	uc         *biz.ChatRoomMemberUseCase
	chatRoomUc *biz.ChatRoomUseCase
}

func NewChatRoomMemberService(uc *biz.ChatRoomMemberUseCase, chatRoomUc *biz.ChatRoomUseCase) *ChatRoomMemberService {
	return &ChatRoomMemberService{uc: uc, chatRoomUc: chatRoomUc}
}

func (s *ChatRoomMemberService) Join(ctx context.Context, in *v1.ChatRoomJoinRequest) (out *v1.ChatRoomJoinResponse, err error) {
	curUid, err := biz.CurrentUidFromContext(ctx)
	if err != nil {
		return
	}
	groupId, err := s.chatRoomUc.FindGroupId(ctx, in.Base.ProductId, in.BizId)
	if err != nil {
		return
	}
	err = s.uc.Join(ctx, &options.ChatRoomJoinOptions{
		ProductId: in.Base.ProductId,
		GroupId:   groupId,
		Uid:       curUid,
	})
	if err != nil {
		return
	}
	out = &v1.ChatRoomJoinResponse{}
	return
}

func (s *ChatRoomMemberService) Quit(ctx context.Context, in *v1.ChatRoomQuitRequest) (out *v1.ChatRoomQuitResponse, err error) {
	curUid, err := biz.CurrentUidFromContext(ctx)
	if err != nil {
		return
	}
	groupId, err := s.chatRoomUc.FindGroupId(ctx, in.Base.ProductId, in.BizId)
	if err != nil {
		return
	}
	err = s.uc.Quit(ctx, &options.ChatRoomQuitOptions{
		ProductId: in.Base.ProductId,
		GroupId:   groupId,
		Uid:       curUid,
	})
	if err != nil {
		return
	}
	out = &v1.ChatRoomQuitResponse{}
	return
}
