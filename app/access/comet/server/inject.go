package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/transport"
	"github.com/soukengo/gopkg/component/transport/socket"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewServerGroup,
	NewGRPCServer,
	NewSocketServer,
	NewServiceRegistry,
	NewSocketRegistry,
)

func NewServerGroup(socket socket.Server, grpc *grpc.Server, socketRegistry *SocketRegistry, grpcRegistry *GrpcRegistry) transport.ServerGroup {
	grpcRegistry.RegisterGrpc(grpc)
	socketRegistry.RegisterSocket(socket)
	return transport.NewServerGroup(socket, grpc)
}
