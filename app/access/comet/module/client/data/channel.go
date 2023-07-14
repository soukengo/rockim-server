package data

import (
	"context"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/session/v1"
	"rockimserver/app/access/comet/module/client/biz"
	"rockimserver/app/access/comet/module/client/biz/options"
)

type channelRepo struct {
	ac v1.ChannelAPIClient
}

func NewChannelRepo(ac v1.ChannelAPIClient) biz.ChannelRepo {
	return &channelRepo{ac: ac}
}

func (r *channelRepo) Add(ctx context.Context, opts *options.ChannelAddOptions) (err error) {
	_, err = r.ac.Add(ctx, &v1.ChannelAddRequest{
		Base:      service.GenRequest(opts.ProductId),
		ServerId:  opts.ServerId,
		ChannelId: opts.ChannelId,
		Uid:       opts.Uid,
	})
	if err != nil {
		return
	}
	return
}

func (r *channelRepo) Refresh(ctx context.Context, opts *options.ChannelRefreshOptions) (err error) {
	_, err = r.ac.Refresh(ctx, &v1.ChannelRefreshRequest{
		Base:      service.GenRequest(opts.ProductId),
		ServerId:  opts.ServerId,
		ChannelId: opts.ChannelId,
		Uid:       opts.Uid,
	})
	return
}

func (r *channelRepo) Delete(ctx context.Context, opts *options.ChannelDeleteOptions) (err error) {
	_, err = r.ac.Delete(ctx, &v1.ChannelDeleteRequest{
		Base:      service.GenRequest(opts.ProductId),
		ServerId:  opts.ServerId,
		ChannelId: opts.ChannelId,
	})
	return
}
