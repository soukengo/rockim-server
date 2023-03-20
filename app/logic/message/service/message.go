package service

import (
	"context"
	v1 "rockimserver/apis/rockim/service/message/v1"
	"rockimserver/app/logic/message/biz"
	"rockimserver/app/logic/message/biz/options"
)

type MessageService struct {
	uc *biz.MessageUseCase
	v1.UnimplementedMessageAPIServer
}

func NewMessageService(uc *biz.MessageUseCase) *MessageService {
	return &MessageService{uc: uc}
}

func (s *MessageService) Send(ctx context.Context, in *v1.MessageSendRequest) (out *v1.MessageSendResponse, err error) {
	err = s.uc.Send(ctx, &options.MessageSendOptions{
		ProductId:   in.Base.ProductId,
		Uid:         in.Uid,
		Target:      in.Target,
		ClientMsgId: in.ClientMsgId,
		MessageType: in.MessageType,
		Content:     in.Content,
		Payload:     in.Payload,
	})
	out = &v1.MessageSendResponse{}
	return
}
