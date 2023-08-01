package server

import (
	"github.com/soukengo/gopkg/component/transport/queue"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/logic/message/biz/consts"
	"rockimserver/app/logic/message/conf"
)

func NewQueueServer(cfg *conf.Server, logger log.Logger) (q queue.Delayed, err error) {
	topics := []queue.Topic{
		consts.QueueMessageDelivery,
	}
	return queue.NewDelayServer(cfg.DelayQueue, topics, logger)
}
