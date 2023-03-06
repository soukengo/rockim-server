package data

import (
	"context"
	"rockimserver/app/access/comet/module/client/biz/options"
)

type onlineRepo struct {
}

func (r *onlineRepo) Connect(ctx context.Context, opts *options.ClientAuthOptions) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *onlineRepo) HeartBeat(ctx context.Context, opts *options.ClientHeartBeatOptions) error {
	//TODO implement me
	panic("implement me")
}

func (r *onlineRepo) DisConnect(ctx context.Context, opts *options.ClientDisConnectOptions) error {
	//TODO implement me
	panic("implement me")
}
