package biz

import (
	"context"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/group/biz/options"
)

type CometRepo interface {
	PushUser(ctx context.Context, opts *options.PushOptions) error
	PushRoom(ctx context.Context, opts *options.PushRoomOptions) error
	SendControl(ctx context.Context, opts *options.ControlOptions) error
}

type CometUseCase struct {
	repo CometRepo
}

func NewCometUseCase(repo CometRepo) *CometUseCase {
	return &CometUseCase{repo: repo}
}

func (uc *CometUseCase) PushUser(ctx context.Context, opts *options.PushOptions) {
	err := uc.repo.PushUser(ctx, opts)
	if err != nil {
		log.Error("PushUser error", log.Pairs{"opts": opts, "err": err})
	}
}

func (uc *CometUseCase) PushRoom(ctx context.Context, opts *options.PushRoomOptions) {
	err := uc.repo.PushRoom(ctx, opts)
	if err != nil {
		log.Error("PushRoom error", log.Pairs{"opts": opts, "err": err})
	}
}

func (uc *CometUseCase) SendControl(ctx context.Context, opts *options.ControlOptions) {
	err := uc.repo.SendControl(ctx, opts)
	if err != nil {
		log.Error("SendControl error", log.Pairs{"opts": opts, "err": err})
	}
}
