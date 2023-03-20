package task

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/app/task/job/biz"
	"rockimserver/pkg/log"
)

type MessageTask struct {
	v1.UnimplementedMessageTaskServer
	uc *biz.MessagePushUseCase
}

func NewMessageTask(uc *biz.MessagePushUseCase) *MessageTask {
	return &MessageTask{uc: uc}
}

func (t *MessageTask) Push(ctx context.Context, in *v1.MessagePushRequest) (out *emptypb.Empty, err error) {
	log.WithContext(ctx).Infof("job push :%v", in)
	err = t.uc.Push(ctx, in)
	return
}
