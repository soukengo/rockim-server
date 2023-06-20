package data

import (
	"context"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/group/v1"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/app/logic/message/biz"
)

type groupRepo struct {
	ac v1.GroupAPIClient
}

func NewGroupRepo(ac v1.GroupAPIClient) biz.GroupRepo {
	return &groupRepo{ac: ac}
}

func (r *groupRepo) Find(ctx context.Context, productId string, bizId string) (out *types.Group, err error) {
	ret, err := r.ac.Find(ctx, &v1.GroupFindRequest{
		Base:  service.GenRequest(productId),
		BizId: bizId,
	})
	if err != nil {
		return
	}
	out = ret.Group
	return
}
