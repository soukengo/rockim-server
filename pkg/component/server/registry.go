package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type HttpServerRegistry interface {
	Register(*http.Server)
}

type GrpcServerRegistry interface {
	Register(*grpc.Server)
}
