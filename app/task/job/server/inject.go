package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/transport"
	"github.com/soukengo/gopkg/component/transport/queue"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewServerGroup,
	NewServiceRegistry,
	NewQueueServer,
	NewGRPCServer,
)

func NewServerGroup(queue queue.Server, grpc *grpc.Server, group *ServiceRegistry) transport.ServerGroup {
	group.RegisterQueue(queue)
	group.RegisterGrpc(grpc)
	return transport.NewServerGroup(queue, grpc)
}
