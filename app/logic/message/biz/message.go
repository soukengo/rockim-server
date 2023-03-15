package biz

import (
	"context"
	"github.com/golang/protobuf/proto"
	"rockimserver/apis/rockim/service/message/v1/types"
	usertypes "rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/message/biz/options"
	biztypes "rockimserver/app/logic/message/biz/types"
	"rockimserver/pkg/component/idgen"
	"rockimserver/pkg/component/lock"
	"time"
)

type MessageRepo interface {
	GenSequence(ctx context.Context, productId string, conversation *types.ConversationID) (int64, error)
	Save(ctx context.Context, message *types.IMMessage) error
	SaveRelations(ctx context.Context, relations []*types.IMMessageRelation) error
}

type MessageUseCase struct {
	repo      MessageRepo
	userRepo  UserRepo
	groupRepo GroupRepo
	lockBdr   lock.Builder
	idgen     idgen.Generator

	pushMgr PushManager
}

func NewMessageUseCase(repo MessageRepo, userRepo UserRepo, groupRepo GroupRepo, lockBdr lock.Builder, idgen idgen.Generator, pushMgr PushManager) *MessageUseCase {
	return &MessageUseCase{repo: repo, userRepo: userRepo, groupRepo: groupRepo, lockBdr: lockBdr, idgen: idgen, pushMgr: pushMgr}
}

func (uc *MessageUseCase) Send(ctx context.Context, opts *options.MessageSendOptions) (err error) {
	var meta *biztypes.MessageMeta
	switch opts.Target.Category {
	case enums.Conversation_PERSON:
		meta, err = uc.completePersonMessage(ctx, opts)
	case enums.Conversation_GROUP:
		meta, err = uc.completeGroupMessage(ctx, opts)
	default:
		err = ErrTargetNotSupported
	}
	if err != nil {
		return
	}
	return uc.confirmMessage(ctx, meta)
}

// completePersonMessage 发送私聊消息
func (uc *MessageUseCase) completePersonMessage(ctx context.Context, opts *options.MessageSendOptions) (meta *biztypes.MessageMeta, err error) {
	recvUser, err := uc.userRepo.FindByAccount(ctx, opts.ProductId, opts.Target.Value)
	if err != nil {
		return
	}
	// todo: check
	body, sender, err := uc.completeMessageBody(ctx, opts)
	if err != nil {
		return
	}
	meta = biztypes.NewPersonMessageMeta(opts.ProductId, sender, opts.Target, recvUser, body)
	meta.Receivers = []string{meta.Sender.Id, meta.RecvUser.Id}
	return
}

// completeGroupMessage 发送群组消息
func (uc *MessageUseCase) completeGroupMessage(ctx context.Context, opts *options.MessageSendOptions) (meta *biztypes.MessageMeta, err error) {
	group, err := uc.groupRepo.Find(ctx, opts.ProductId, opts.Target.Value)
	if err != nil {
		return
	}
	// todo: check
	body, sender, err := uc.completeMessageBody(ctx, opts)
	if err != nil {
		return
	}
	meta = biztypes.NewGroupMessageMeta(opts.ProductId, sender, opts.Target, group, body)
	// todo: 按照群成员扩散写
	//meta.Receivers = []string{}
	return
}

// confirmMessage 确认消息
func (uc *MessageUseCase) confirmMessage(ctx context.Context, meta *biztypes.MessageMeta) (err error) {
	productId := meta.ProductId
	msgId, err := uc.idgen.GenID()
	if err != nil {
		return
	}
	conversation := meta.ConversationId()
	sequence, err := uc.repo.GenSequence(ctx, productId, conversation)
	if err != nil {
		return
	}
	message := &types.IMMessage{
		ProductId:      meta.ProductId,
		MsgId:          msgId,
		ConversationId: meta.ConversationId(),
		Body:           meta.Body,
		Sequence:       sequence,
		Status:         enums.Message_SUCCESS,
		Version:        0,
	}
	// 同步存储消息
	err = uc.repo.Save(ctx, message)
	if err != nil {
		return
	}
	// 异步扩散写
	err = uc.diffusionWrite(ctx, message, meta.Receivers)
	if err != nil {
		return
	}
	err = uc.dispatchMessage(ctx, message, meta.Receivers)
	return
}

// diffusionWrite 扩散写
func (uc *MessageUseCase) diffusionWrite(ctx context.Context, message *types.IMMessage, receivers []string) (err error) {
	relations := make([]*types.IMMessageRelation, len(receivers))
	for i, receiver := range receivers {
		relations[i] = &types.IMMessageRelation{
			ProductId:      message.ProductId,
			ConversationId: message.ConversationId,
			MsgId:          message.MsgId,
			Uid:            receiver,
		}
	}
	err = uc.repo.SaveRelations(ctx, relations)
	return
}

// dispatchMessage 分发消息
func (uc *MessageUseCase) dispatchMessage(ctx context.Context, message *types.IMMessage, receivers []string) (err error) {
	body, err := proto.Marshal(message)
	if err != nil {
		return
	}
	operation := enums.Network_MESSAGES
	switch message.ConversationId.Category {
	case enums.Conversation_PERSON:
		err = uc.pushMgr.PushUsers(ctx, operation, message.ProductId, receivers, body)
	case enums.Conversation_GROUP:
		err = uc.pushMgr.PushGroup(ctx, operation, message.ProductId, message.ConversationId.Value, body)
	}
	return
}

// completeMessageBody 补全消息数据
func (uc *MessageUseCase) completeMessageBody(ctx context.Context, opts *options.MessageSendOptions) (body *types.IMMessageBody, sender *usertypes.User, err error) {
	sender, err = uc.userRepo.FindById(ctx, opts.ProductId, opts.Uid)
	if err != nil {
		return
	}
	body = &types.IMMessageBody{
		Timestamp: time.Now().UnixMilli(),
		Sender: &types.IMMessageSender{
			Uid:       sender.Id,
			Account:   sender.Account,
			Name:      sender.Name,
			AvatarUrl: sender.AvatarUrl,
		},
		ClientMsgId: opts.ClientMsgId,
		MessageType: opts.MessageType,
		Content:     opts.Content,
		Payload:     opts.Payload,
	}
	return
}
