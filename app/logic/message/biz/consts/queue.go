package consts

import (
	"github.com/soukengo/gopkg/component/queue"
	"rockimserver/apis/rockim/service"
)

var (
	QueueMessageDelivery = queue.Topic("message.delivery")
)

var (
	DistributedQueueMessagePush = queue.Topic(service.MQ_MESSAGE_PUSH.String())
)
