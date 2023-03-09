package biz

import (
	"context"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/logic/user/biz/options"
	"rockimserver/app/logic/user/biz/types"
	"rockimserver/pkg/errors"
	"time"
)

type OnlineRepo interface {
	Add(ctx context.Context, ch *types.OnlineChannel) error
	Exists(ctx context.Context, ch *types.OnlineChannel) (bool, error)
	Delete(ctx context.Context, ch *types.OnlineChannel) error
}

var (
	ErrUnSpecified = errors.BadRequest(reasons.ErrorReason_UN_SPECIFIED.String(), "未知错误")
)

type OnlineUseCase struct {
	repo      OnlineRepo
	tokenRepo AccessTokenRepo
}

func NewOnlineUseCase(repo OnlineRepo, tokenRepo AccessTokenRepo) *OnlineUseCase {
	return &OnlineUseCase{repo: repo, tokenRepo: tokenRepo}
}

func (uc *OnlineUseCase) Add(ctx context.Context, opts *options.OnlineAddOptions) (uid string, err error) {
	uid, err = uc.tokenRepo.FindByAccessToken(ctx, opts.ProductId, opts.Token)
	if err != nil {
		return
	}
	ch := &types.OnlineChannel{
		ProductId:  opts.ProductId,
		Uid:        uid,
		ServerId:   opts.ServerId,
		ChannelId:  opts.ChannelId,
		UpdateTime: time.Now().UnixMilli(),
	}
	err = uc.repo.Add(ctx, ch)
	return
}

func (uc *OnlineUseCase) Refresh(ctx context.Context, opts *options.OnlineRefreshOptions) (err error) {
	ch := &types.OnlineChannel{
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

func (uc *OnlineUseCase) Delete(ctx context.Context, opts *options.OnlineDeleteOptions) (err error) {
	ch := &types.OnlineChannel{
		ProductId: opts.ProductId,
		Uid:       opts.Uid,
		ServerId:  opts.ServerId,
		ChannelId: opts.ChannelId,
	}
	err = uc.repo.Delete(ctx, ch)
	return
}
