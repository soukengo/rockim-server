package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/client/v1/protocol/http"
	"rockimserver/apis/rockim/service/message/v1/types"
	"rockimserver/app/access/gateway/module/client/biz"
	"rockimserver/app/access/gateway/module/client/biz/options"
)

type MessageService struct {
	uc *biz.MessageUseCase
}

func NewMessageService(uc *biz.MessageUseCase) *MessageService {
	return &MessageService{uc: uc}
}

func (s *MessageService) Send(ctx context.Context, in *v1.MessageSendRequest) (out *v1.MessageSendResponse, err error) {
	curUid, err := biz.CurrentUidFromContext(ctx)
	if err != nil {
		return
	}
	err = s.uc.Send(ctx, &options.MessageSendOptions{
		ProductId:   in.Base.ProductId,
		Uid:         curUid,
		Target:      &types.TargetID{Category: in.Target.Category, Value: in.Target.Value},
		ClientMsgId: in.ClientMsgId,
		MessageType: in.MessageType,
		Content:     in.Content,
		Payload:     in.Payload,
	})
	if err != nil {
		return
	}
	out = &v1.MessageSendResponse{}
	return
}
