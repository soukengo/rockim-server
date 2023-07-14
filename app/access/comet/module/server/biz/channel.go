package biz

import (
	"context"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/access/comet/module/server/biz/options"
)

type ChannelRepo interface {
	Push(ctx context.Context, opts *options.PushOptions) error
	PushRoom(ctx context.Context, opts *options.PushRoomOptions) error

	Control(ctx context.Context, opts *options.ControlOptions) error
}

type ChannelUseCase struct {
	repo ChannelRepo
}

func NewChannelUseCase(repo ChannelRepo) *ChannelUseCase {
	return &ChannelUseCase{repo: repo}
}

func (uc *ChannelUseCase) Dispatch(ctx context.Context, opts *options.DispatchOptions) error {
	switch opts.DataType {
	case enums.Comet_PUSH:
		return uc.repo.Push(ctx, &options.PushOptions{
			ChannelIds: opts.ChannelIds,
			Operation:  opts.Push.Operation,
			Body:       opts.Push.Body,
		})
	case enums.Comet_CONTROL:
		return uc.repo.Control(ctx, &options.ControlOptions{
			ChannelIds:  opts.ChannelIds,
			ControlType: opts.Control.ControlType,
			Body:        opts.Control.Body,
		})
	}
	return nil
}

func (uc *ChannelUseCase) DispatchRoom(ctx context.Context, opts *options.DispatchRoomOptions) error {
	return uc.repo.PushRoom(ctx, &options.PushRoomOptions{
		Room:      opts.Room,
		Operation: opts.Push.Operation,
		Body:      opts.Push.Body,
	})
}
