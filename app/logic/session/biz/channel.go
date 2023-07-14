package biz

import (
	"context"
	"github.com/soukengo/gopkg/errors"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/logic/session/biz/options"
	"rockimserver/app/logic/session/biz/types"
	"time"
)

type ChannelRepo interface {
	Add(ctx context.Context, ch *types.Channel) error
	Exists(ctx context.Context, ch *types.Channel) (bool, error)
	Delete(ctx context.Context, ch *types.Channel) error
}

var (
	ErrUnSpecified = errors.BadRequest(reasons.ErrorReason_UN_SPECIFIED.String(), "未知错误")
)

type ChannelUseCase struct {
	repo ChannelRepo
}

func NewChannelUseCase(repo ChannelRepo) *ChannelUseCase {
	return &ChannelUseCase{repo: repo}
}

func (uc *ChannelUseCase) Add(ctx context.Context, opts *options.ChannelAddOptions) (err error) {
	ch := &types.Channel{
		ProductId:  opts.ProductId,
		Uid:        opts.Uid,
		ServerId:   opts.ServerId,
		ChannelId:  opts.ChannelId,
		UpdateTime: time.Now().UnixMilli(),
	}
	err = uc.repo.Add(ctx, ch)
	return
}

func (uc *ChannelUseCase) Refresh(ctx context.Context, opts *options.ChannelRefreshOptions) (err error) {
	ch := &types.Channel{
		ProductId:  opts.ProductId,
		Uid:        opts.Uid,
		ServerId:   opts.ServerId,
		ChannelId:  opts.ChannelId,
		UpdateTime: time.Now().UnixMilli(),
	}
	exists, err := uc.repo.Exists(ctx, ch)
	if err != nil {
		return
	}
	if !exists {
		err = ErrUnSpecified.NewWithMessage("无效的请求")
		return
	}
	err = uc.repo.Add(ctx, ch)
	return
}

func (uc *ChannelUseCase) Delete(ctx context.Context, opts *options.ChannelDeleteOptions) (err error) {
	ch := &types.Channel{
		ProductId: opts.ProductId,
		Uid:       opts.Uid,
		ServerId:  opts.ServerId,
		ChannelId: opts.ChannelId,
	}
	err = uc.repo.Delete(ctx, ch)
	return
}
