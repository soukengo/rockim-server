package data

import (
	"context"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/user/v1"
	"rockimserver/app/access/comet/module/client/biz"
	"rockimserver/app/access/comet/module/client/biz/options"
)

type onlineRepo struct {
	ac v1.OnlineAPIClient
}

func NewOnlineRepo(ac v1.OnlineAPIClient) biz.OnlineRepo {
	return &onlineRepo{ac: ac}
}

func (r *onlineRepo) Add(ctx context.Context, opts *options.OnlineAddOptions) (uid string, err error) {
	ret, err := r.ac.Add(ctx, &v1.OnlineAddRequest{
		Base:      service.GenRequest(opts.ProductId),
		ServerId:  opts.ServerId,
		ChannelId: opts.ChannelId,
		Token:     opts.Token,
	})
	if err != nil {
		return
	}
	uid = ret.Uid
	return
}

func (r *onlineRepo) Refresh(ctx context.Context, opts *options.OnlineRefreshOptions) (err error) {
	_, err = r.ac.Refresh(ctx, &v1.OnlineRefreshRequest{
		Base:      service.GenRequest(opts.ProductId),
		ServerId:  opts.ServerId,
		ChannelId: opts.ChannelId,
	})
	return
}

func (r *onlineRepo) Delete(ctx context.Context, opts *options.OnlineDeleteOptions) (err error) {
	_, err = r.ac.Delete(ctx, &v1.OnlineDeleteRequest{
		Base:      service.GenRequest(opts.ProductId),
		ServerId:  opts.ServerId,
		ChannelId: opts.ChannelId,
	})
	return
}
