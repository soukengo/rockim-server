package biz

import (
	"context"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/message/biz/options"
	"rockimserver/pkg/component/idgen"
	"rockimserver/pkg/component/lock"
)

type MessageRepo interface {
	Save(ctx context.Context) error
}

type MessageUseCase struct {
	repo     MessageRepo
	userRepo UserRepo
	lockBdr  lock.Builder
	idgen    idgen.Generator
}

func NewMessageUseCase(repo MessageRepo, lockBdr lock.Builder, idgen idgen.Generator) *MessageUseCase {
	return &MessageUseCase{repo: repo, lockBdr: lockBdr, idgen: idgen}
}

func (uc *MessageUseCase) Send(ctx context.Context, opts *options.MessageSendOptions) (err error) {
	switch opts.ConversationId.Category {
	case enums.Conversation_PERSON:
		return uc.sendPersonMessage(ctx, opts)
	default:
		err = ErrTargetNotSupported
	}
	return
}

func (uc *MessageUseCase) sendPersonMessage(ctx context.Context, opts *options.MessageSendOptions) (err error) {
	_, err = uc.userRepo.FindById(ctx, opts.ProductId, opts.Uid)
	if err != nil {
		return
	}
	return
}
