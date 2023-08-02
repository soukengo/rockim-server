package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/component/transport"
	"github.com/soukengo/gopkg/component/transport/queue"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	discovery.NewRegistrar,
	NewServerGroup,
	NewGRPCServer,
	NewQueueServer,
	NewServiceRegistry,
	NewTaskRegistry,
)

func NewServerGroup(delayed queue.Delayed, grpc *grpc.Server, service *ServiceRegistry, task *TaskRegistry) transport.ServerGroup {
	service.RegisterGrpc(grpc)
	task.RegisterQueue(delayed)
	return transport.NewServerGroup(delayed, grpc)
}
