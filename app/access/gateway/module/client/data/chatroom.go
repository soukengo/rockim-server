package data

import (
	"context"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/group/v1"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/app/access/gateway/module/client/biz"
)

type chatRoomRepo struct {
	ac      v1.ChatRoomAPIClient
	groupAc v1.GroupAPIClient
}

func NewChatRoomRepo(ac v1.ChatRoomAPIClient, groupAc v1.GroupAPIClient) biz.ChatRoomRepo {
	return &chatRoomRepo{ac: ac, groupAc: groupAc}
}

func (r *chatRoomRepo) FindById(ctx context.Context, productId string, groupId string) (out *types.Group, err error) {
	ret, err := r.groupAc.FindById(ctx, &v1.GroupFindByIdRequest{
		Base:    service.GenRequest(productId),
		GroupId: groupId,
	})
	if err != nil {
		return
	}
	return ret.Group, nil
}

func (r *chatRoomRepo) FindGroupId(ctx context.Context, productId string, customGroupId string) (out string, err error) {
	ret, err := r.groupAc.FindGroupId(ctx, &v1.GroupIdFindRequest{
		Base:          service.GenRequest(productId),
		CustomGroupId: customGroupId,
	})
	if err != nil {
		return
	}
	return ret.GroupId, nil
}
