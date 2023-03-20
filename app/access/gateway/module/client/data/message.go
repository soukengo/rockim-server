package data

import (
	"context"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/message/v1"
	"rockimserver/app/access/gateway/module/client/biz"
	"rockimserver/app/access/gateway/module/client/biz/options"
)

type messageRepo struct {
	ac v1.MessageAPIClient
}

func NewMessageRepo(ac v1.MessageAPIClient) biz.MessageRepo {
	return &messageRepo{ac: ac}
}

func (r *messageRepo) Send(ctx context.Context, in *options.MessageSendOptions) (err error) {
	_, err = r.ac.Send(ctx, &v1.MessageSendRequest{
		Base:        service.GenRequest(in.ProductId),
		Uid:         in.Uid,
		Target:      in.Target,
		ClientMsgId: in.ClientMsgId,
		MessageType: in.MessageType,
		Content:     in.Content,
		Payload:     in.Payload,
	})
	return
}
