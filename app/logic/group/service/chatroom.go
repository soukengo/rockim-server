package service

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1"
	"rockimserver/app/logic/group/biz"
	"rockimserver/app/logic/group/biz/options"
)

type ChatRoomService struct {
	uc *biz.ChatRoomUseCase
	v1.UnimplementedChatRoomAPIServer
}

func NewChatRoomService(uc *biz.ChatRoomUseCase) *ChatRoomService {
	return &ChatRoomService{uc: uc}
}

func (s *ChatRoomService) Create(ctx context.Context, in *v1.ChatRoomCreateRequest) (out *v1.ChatRoomCreateResponse, err error) {
	group, err := s.uc.Create(ctx, &options.ChatRoomCreateOptions{
		ProductId:     in.Base.ProductId,
		CustomGroupId: in.CustomGroupId,
		Name:          in.Name,
		IconUrl:       in.IconUrl,
		Fields:        in.Fields,
		Owner:         in.Owner,
	})
	if err != nil {
		return
	}
	out = &v1.ChatRoomCreateResponse{Group: group}
	return
}

func (s *ChatRoomService) Dismiss(ctx context.Context, in *v1.ChatRoomDismissRequest) (out *v1.ChatRoomDismissResponse, err error) {
	err = s.uc.Dismiss(ctx, &options.ChatRoomDismissOptions{
		ProductId: in.Base.ProductId,
		GroupId:   in.GroupId,
	})
	if err != nil {
		return
	}
	out = &v1.ChatRoomDismissResponse{}
	return

}

func (s *ChatRoomService) FindGroupId(ctx context.Context, in *v1.ChatRoomGroupIdFindRequest) (out *v1.ChatRoomGroupIdFindResponse, err error) {
	groupId, err := s.uc.FindByCustomGroupId(ctx, in.Base.ProductId, in.CustomGroupId)
	if err != nil {
		return
	}
	out = &v1.ChatRoomGroupIdFindResponse{GroupId: groupId}
	return
}

func (s *ChatRoomService) FindById(ctx context.Context, in *v1.ChatRoomFindByIdRequest) (out *v1.ChatRoomFindByIdResponse, err error) {
	group, err := s.uc.FindById(ctx, in.Base.ProductId, in.GroupId)
	if err != nil {
		return
	}
	out = &v1.ChatRoomFindByIdResponse{Group: group}
	return
}
