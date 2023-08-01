package consts

import (
	"github.com/soukengo/gopkg/component/transport/queue"
	v1 "rockimserver/apis/rockim/service/job/v1"
)

var (
	QueueCometDispatch = queue.Topic(v1.TaskType_COMET_DISPATCH.String())
)
