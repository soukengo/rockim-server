package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/openapi/v1"
	"rockimserver/app/access/gateway/module/openapi/biz"
	"rockimserver/app/access/gateway/module/openapi/biz/options"
)

type ChatRoomService struct {
	uc     *biz.ChatRoomUseCase
	userUc *biz.UserUseCase
}

func NewChatRoomService(uc *biz.ChatRoomUseCase, userUc *biz.UserUseCase) *ChatRoomService {
	return &ChatRoomService{uc: uc, userUc: userUc}
}

func (s *ChatRoomService) Create(ctx context.Context, in *v1.ChatRoomCreateRequest) (out *v1.ChatRoomCreateResponse, err error) {
	var ownerUid string
	if len(in.OwnerAccount) > 0 {
		ownerUid, err = s.userUc.FindUid(ctx, in.Base.ProductId, in.OwnerAccount)
		if err != nil {
			return
		}
	}
	group, err := s.uc.Create(ctx, &options.ChatRoomCreateOptions{
		ProductId:     in.Base.ProductId,
		CustomGroupId: in.CustomGroupId,
		Name:          in.Name,
		IconUrl:       in.IconUrl,
		Fields:        in.Fields,
		Owner:         ownerUid,
	})
	if err != nil {
		return
	}
	out = &v1.ChatRoomCreateResponse{CustomGroupId: group.CustomGroupId}
	return
}
