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

func (s *ChatRoomMemberService) IsMember(ctx context.Context, in *v1.ChatRoomMemberCheckRequest) (out *v1.ChatRoomMemberCheckResponse, err error) {
	isMember, err := s.uc.IsMember(ctx, &options.ChatRoomMemberCheckOptions{
		ProductId: in.Base.ProductId,
		GroupId:   in.GroupId,
		Uid:       in.Uid,
	})
	if err != nil {
		return
	}
	out = &v1.ChatRoomMemberCheckResponse{IsMember: isMember}
	return
}

func (s *ChatRoomMemberService) Find(ctx context.Context, in *v1.ChatRoomMemberFindRequest) (out *v1.ChatRoomMemberFindResponse, err error) {
	member, err := s.uc.Find(ctx, &options.ChatRoomMemberFindOptions{
		ProductId: in.Base.ProductId,
		GroupId:   in.GroupId,
		Uid:       in.Uid,
	})
	if err != nil {
		return
	}
	out = &v1.ChatRoomMemberFindResponse{Member: member}
	return
}

func (s *ChatRoomMemberService) List(ctx context.Context, in *v1.ChatRoomMemberListRequest) (out *v1.ChatRoomMemberListResponse, err error) {
	list, err := s.uc.List(ctx, &options.ChatRoomMemberListOptions{
		ProductId: in.Base.ProductId,
		GroupId:   in.GroupId,
		Uids:      in.Uids,
	})
	if err != nil {
		return
	}
	out = &v1.ChatRoomMemberListResponse{Members: list}
	return
}
func (s *ChatRoomMemberService) PaginateUid(ctx context.Context, in *v1.ChatRoomMemberIdPaginateRequest) (out *v1.ChatRoomMemberIdPaginateResponse, err error) {
	list, paginated, err := s.uc.PaginateUid(ctx, &options.ChatRoomMemberIdPaginateOptions{
		ProductId: in.Base.ProductId,
		GroupId:   in.GroupId,
	})
	if err != nil {
		return
	}
	out = &v1.ChatRoomMemberIdPaginateResponse{Members: list, Paginate: paginated}
	return
}
