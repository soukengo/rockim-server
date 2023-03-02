package data

import (
	"context"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/group/v1"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/app/access/gateway/module/client/biz"
)

type chatRoomRepo struct {
	ac v1.ChatRoomAPIClient
}

func NewChatRoomRepo(ac v1.ChatRoomAPIClient) biz.ChatRoomRepo {
	return &chatRoomRepo{ac: ac}
}

func (r *chatRoomRepo) Find(ctx context.Context, productId string, groupId string) (out *types.Group, err error) {
	ret, err := r.ac.FindById(ctx, &v1.ChatRoomFindByIdRequest{
		Base:    service.GenServiceRequest(productId),
		GroupId: groupId,
	})
	if err != nil {
		return
	}
	return ret.Group, nil
}

func (r *chatRoomRepo) FindGroupId(ctx context.Context, productId string, customGroupId string) (out string, err error) {
	ret, err := r.ac.FindGroupId(ctx, &v1.ChatRoomGroupIdFindRequest{
		Base:          service.GenServiceRequest(productId),
		CustomGroupId: customGroupId,
	})
	if err != nil {
		return
	}
	return ret.GroupId, nil
}
