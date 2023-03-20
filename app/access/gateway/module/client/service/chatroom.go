package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/client/v1/http"
	"rockimserver/app/access/gateway/module/client/biz"
	"rockimserver/app/access/gateway/module/client/service/converter"
)

type ChatRoomService struct {
	uc *biz.ChatRoomUseCase
}

func NewChatRoomService(uc *biz.ChatRoomUseCase) *ChatRoomService {
	return &ChatRoomService{uc: uc}
}

func (s *ChatRoomService) Find(ctx context.Context, in *v1.ChatRoomFindRequest) (out *v1.ChatRoomFindResponse, err error) {
	group, err := s.uc.Find(ctx, in.Base.ProductId, in.CustomGroupId)
	if err != nil {
		return nil, err
	}
	return &v1.ChatRoomFindResponse{Group: converter.ToClientGroup(group)}, nil
}
