package middleware

import (
	"context"
	kratoserrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/soukengo/gopkg/errors"
	"github.com/soukengo/gopkg/log"
	"github.com/soukengo/gopkg/util/encoding/json"
	"go.opentelemetry.io/otel/trace"
	"rockimserver"
	"time"
)

const (
	ServerHeaderTraceID = "RockIM-Server-TraceID"
	ServerHeaderVersion = "RockIM-Server-Version"
)

// Trace 调用链路追踪中间件
func Trace() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tr.ReplyHeader().Set(ServerHeaderTraceID, trace.SpanContextFromContext(ctx).TraceID().String())
				tr.ReplyHeader().Set(ServerHeaderVersion, rockimserver.Version())
			}
			return handler(ctx, req)
		}
	}
}

func Logging(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			var (
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				operation = info.Operation()
			}
			reply, err = handler(ctx, req)
			if err != nil {
				e := kratoserrors.InternalServer("", err.Error())
				if se := errors.FromError(err); se != nil {
					e.Code = se.Code()
					e.Reason = se.Reason()
					e.Message = se.Message()
					e.Metadata = se.Metadata()
				}
				reply = e

			}
			logger.Helper().WithContext(ctx).Info("requestLog", log.Pairs{
				"operation": operation,
				"req":       json.TryToString(req),
				"reply":     json.TryToString(reply),
				"cost":      time.Since(startTime).Milliseconds()})
			return
		}
	}
}
