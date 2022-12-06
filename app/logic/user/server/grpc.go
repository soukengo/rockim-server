package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "rockim/api/rockim/service/user/v1"
	"rockim/app/logic/user/conf"
	"rockim/app/logic/user/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, user *service.UserService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout > 0 {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserAPIServer(srv, user)
	return srv
}
