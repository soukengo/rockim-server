package data

import (
	"context"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/group/v1"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/app/access/gateway/module/openapi/biz"
	"rockimserver/app/access/gateway/module/openapi/biz/options"
)

type chatRoomRepo struct {
	ac v1.ChatRoomAPIClient
}

func NewChatRoomRepo(ac v1.ChatRoomAPIClient) biz.ChatRoomRepo {
	return &chatRoomRepo{ac: ac}
}

func (r *chatRoomRepo) Create(ctx context.Context, opts *options.ChatRoomCreateOptions) (out *types.Group, err error) {
	ret, err := r.ac.Create(ctx, &v1.ChatRoomCreateRequest{
		Base:          service.GenRequest(opts.ProductId),
		CustomGroupId: opts.CustomGroupId,
		Name:          opts.Name,
		IconUrl:       opts.IconUrl,
		Fields:        opts.Fields,
		Owner:         opts.Owner,
	})
	if err != nil {
		return
	}
	out = ret.Group
	return
}
