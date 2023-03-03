package server

import "context"

type HandlerFunc func(context.Context, any) (any, error)

type HandlerRegistrar interface {
	Handle(operation int32, h HandlerFunc)
}
