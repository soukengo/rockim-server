package data

import (
	"context"
	"github.com/samber/lo"
	"rockimserver/apis/rockim/service"
	"rockimserver/apis/rockim/service/comet/v1/types"
	v1 "rockimserver/apis/rockim/service/group/v1"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/access/comet/module/client/biz"
	"rockimserver/app/access/comet/module/client/biz/options"
)

type roomRepo struct {
	ac v1.GroupMemberAPIClient
}

func NewRoomRepo(ac v1.GroupMemberAPIClient) biz.RoomRepo {
	return &roomRepo{ac: ac}
}

func (r *roomRepo) List(ctx context.Context, opts *options.ListRoomOptions) (out []*types.Room, err error) {
	ret, err := r.ac.ListGroupIdByUid(ctx, &v1.GroupIdListByUidRequest{
		Base: service.GenRequest(opts.ProductId),
		Uid:  opts.Uid,
	})
	if err != nil {
		return
	}
	out = lo.Map(ret.GroupIds, func(item string, index int) *types.Room {
		return &types.Room{RoomType: enums.Comet_GROUP, BizId: item}
	})
	return
}
