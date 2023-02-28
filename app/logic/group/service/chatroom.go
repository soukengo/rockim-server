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
	})
	if err != nil {
		return
	}
	out = &v1.ChatRoomCreateResponse{GroupId: group.Id}
	return
}

func (s *ChatRoomService) Dismiss(ctx context.Context, request *v1.ChatRoomDismissRequest) (*v1.ChatRoomDismissResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChatRoomService) Find(ctx context.Context, request *v1.ChatRoomFindRequest) (*v1.ChatRoomFindResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChatRoomService) FindById(ctx context.Context, request *v1.ChatRoomFindRequest) (*v1.ChatRoomFindResponse, error) {
	//TODO implement me
	panic("implement me")
}
