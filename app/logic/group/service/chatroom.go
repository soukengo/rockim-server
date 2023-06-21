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
