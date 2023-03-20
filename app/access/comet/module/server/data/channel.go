package data

import (
	"context"
	"rockimserver/app/access/comet/module/server/biz"
	"rockimserver/app/access/comet/module/server/biz/options"
	"rockimserver/app/access/comet/module/server/data/socket"
)

type channelRepo struct {
	socket *socket.ChannelManager
}

func NewChannelRepo(socket *socket.ChannelManager) biz.ChannelRepo {
	return &channelRepo{socket: socket}
}

func (r *channelRepo) Push(ctx context.Context, opts *options.PushOptions) error {
	return r.socket.Push(ctx, opts)
}

func (r *channelRepo) PushGroup(ctx context.Context, opts *options.PushGroupOptions) error {
	return r.socket.PushGroup(ctx, opts)
}
