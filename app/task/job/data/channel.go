package data

import (
	"context"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/session/v1"
	"rockimserver/apis/rockim/service/session/v1/types"
	"rockimserver/app/task/job/biz"
)

type channelRepo struct {
	ac v1.ChannelQueryAPIClient
}

func NewChannelRepo(ac v1.ChannelQueryAPIClient) biz.ChannelRepo {
	return &channelRepo{ac: ac}
}

func (r *channelRepo) ListChannels(ctx context.Context, productId string, uids []string) (out []*types.UserChannel, err error) {
	ret, err := r.ac.ListUser(ctx, &v1.UserChannelListRequest{
		Base: service.GenRequest(productId),
		Uids: uids,
	})
	if err != nil {
		return
	}
	out = ret.Users
	return
}
