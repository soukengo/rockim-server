package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/discovery"
	"github.com/soukengo/gopkg/component/transport"
	"github.com/soukengo/gopkg/component/transport/event"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	discovery.NewRegistrar,
	NewServiceRegistry,
	NewListenerRegistry,
	NewGRPCServer,
	NewEventServer,
	NewServerGroup,
)

func NewServerGroup(ev event.Server, grpc *grpc.Server, service *ServiceRegistry, listener *ListenerRegistry) transport.ServerGroup {
	service.RegisterGrpc(grpc)
	listener.RegisterEvent(ev)
	return transport.NewServerGroup(ev, grpc)
}
