package log

import (
	"context"
	"github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
)

var (
	defKvs = []any{
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	}
)

type kratosAdapter struct {
	pvd log.Logger
}

func (adp *kratosAdapter) Log(level Level, keyvals ...interface{}) error {
	return adp.pvd.Log(log.Level(level), keyvals...)
}

func (adp *kratosAdapter) WithContext(ctx context.Context) Logger {
	return &kratosAdapter{pvd: log.WithContext(ctx, adp.pvd)}
}

func newAdapter(level Level, appender *AppenderConfig, kvs ...any) Logger {
	var writer = newAppender(appender)
	if len(kvs) == 0 {
		kvs = make([]any, 8)
	}
	kvs = append(kvs, defKvs...)
	return &kratosAdapter{pvd: log.With(zap.NewLogger(newZap(level.String(), writer)), kvs...)}
}

type kratosLogger struct {
	logger Logger
}

func (k kratosLogger) Log(level log.Level, keyvals ...interface{}) error {
	return k.logger.Log(Level(level), keyvals...)
}

func configureKratos(logger Logger) {
	log.SetLogger(&kratosLogger{logger: logger})
}
