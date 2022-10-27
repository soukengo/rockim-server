package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
)

func V2Auth(h middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req any) (any, error) {
		return h(ctx, req)
	}
}
