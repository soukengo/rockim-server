package biz

import (
	"context"
	"github.com/soukengo/gopkg/component/idgen"
	"github.com/soukengo/gopkg/component/queue"
	"github.com/soukengo/gopkg/log"
	"rockimserver/apis/rockim/service/message/v1/types"
	usertypes "rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/message/biz/consts"
	"rockimserver/app/logic/message/biz/options"
	biztypes "rockimserver/app/logic/message/biz/types"
	"time"
)

type MessageUseCase struct {
	repo         MessageRepo
	deliveryRepo MessageDeliveryRepo
	boxRepo      MessageBoxRepo
	userRepo     UserRepo
	groupRepo    GroupRepo
	idgen        idgen.Generator
	delayed      queue.DelayedProducer
}

func NewMessageUseCase(repo MessageRepo, deliveryRepo MessageDeliveryRepo, boxRepo MessageBoxRepo, userRepo UserRepo, groupRepo GroupRepo, idgen idgen.Generator, delayed queue.DelayedProducer) *MessageUseCase {
	return &MessageUseCase{repo: repo, deliveryRepo: deliveryRepo, boxRepo: boxRepo, userRepo: userRepo, groupRepo: groupRepo, idgen: idgen, delayed: delayed}
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
	meta.MsgId = msgId
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
	err = uc.diffusionWrite(ctx, meta)
	if err != nil {
		return
	}
	err = uc.dispatchMessage(ctx, meta, message)
	return
}

// diffusionWrite 扩散写
func (uc *MessageUseCase) diffusionWrite(ctx context.Context, meta *biztypes.MessageMeta) (err error) {
	if len(meta.Receivers) == 0 {
		return
	}
	relations := make([]*types.IMMessageRelation, len(meta.Receivers))
	letters := make([]*biztypes.IMMessageLetter, len(meta.Receivers))
	for i, receiver := range meta.Receivers {
		relations[i] = &types.IMMessageRelation{
			ProductId:      meta.ProductId,
			ConversationId: meta.ConversationId(),
			MsgId:          meta.MsgId,
			Uid:            receiver,
		}
		letters[i] = &biztypes.IMMessageLetter{
			ConversationId: meta.ConversationId(),
			MsgId:          meta.MsgId,
			Uid:            receiver,
			Timestamp:      meta.Body.Timestamp,
		}
	}
	err = uc.repo.SaveRelations(ctx, meta.ProductId, relations)
	if err != nil {
		return
	}
	err1 := uc.boxRepo.SaveLetters(ctx, meta.ProductId, letters)
	if err1 != nil {
		log.WithContext(ctx).Errorf("SaveLetters error: %v", err1)
	}
	return
}

const (
	groupMessageDelay = time.Millisecond * 100
)

// dispatchMessage 分发消息
func (uc *MessageUseCase) dispatchMessage(ctx context.Context, meta *biztypes.MessageMeta, message *types.IMMessage) (err error) {
	delivery := biztypes.NewIMMessageDelivery(message.ConversationId, message.MsgId, meta.FromTarget(), meta.Target)
	err = uc.deliveryRepo.Save(ctx, message.ProductId, delivery)
	if err != nil {
		return
	}
	conversationId := biztypes.EncodeConversationID(message.ProductId, message.ConversationId)
	// 群消息延时投递
	if message.ConversationId.Category == enums.Conversation_GROUP {
		return uc.delayed.SubmitDelay(ctx, consts.QueueMessageDelivery, []queue.Value{queue.Value(conversationId)}, groupMessageDelay.Milliseconds(), false)
	}
	err = uc.delayed.Submit(ctx, consts.QueueMessageDelivery, []queue.Value{queue.Value(conversationId)})
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
