package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"go.opentelemetry.io/otel/trace"
	"rockimserver"
)

const (
	ServerHeaderTraceID = "RockIM-Server-TraceID"
	ServerHeaderVersion = "RockIM-Server-Version"
)

// Trace 调用链路追踪中间件
func Trace() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tr.ReplyHeader().Set(ServerHeaderTraceID, trace.SpanContextFromContext(ctx).TraceID().String())
				tr.ReplyHeader().Set(ServerHeaderVersion, rockimserver.Version())
			}
			return handler(ctx, req)
		}
	}
}
