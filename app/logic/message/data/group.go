package data

import (
	"context"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/group/v1"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/message/biz"
)

type groupRepo struct {
	ac   v1.GroupAPIClient
	gmac v1.GroupMemberAPIClient
}

func NewGroupRepo(ac v1.GroupAPIClient, gmac v1.GroupMemberAPIClient) biz.GroupRepo {
	return &groupRepo{ac: ac, gmac: gmac}
}

func (r *groupRepo) Find(ctx context.Context, productId string, customGroupId string) (out *types.Group, err error) {
	ret, err := r.ac.Find(ctx, &v1.GroupFindRequest{
		Base:          service.GenRequest(productId),
		CustomGroupId: customGroupId,
	})
	if err != nil {
		return
	}
	out = ret.Group
	return
}

func (r *groupRepo) IsMember(ctx context.Context, productId string, category enums.Group_Category, groupId string, uid string) (isMember bool, err error) {
	ret, err := r.gmac.IsMember(ctx, &v1.GroupMemberCheckRequest{
		Base:     service.GenRequest(productId),
		Category: category,
		GroupId:  groupId,
		Uid:      uid,
	})
	if err != nil {
		return
	}
	return ret.IsMember, nil
}
