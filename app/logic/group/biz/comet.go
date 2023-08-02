package biz

import (
	"context"
	"rockimserver/app/logic/group/biz/options"
)

type CometRepo interface {
	PushUser(ctx context.Context, opts *options.PushOptions) error
	PushRoom(ctx context.Context, opts *options.PushRoomOptions)
	SendControl(ctx context.Context, opts *options.ControlOptions)
}
