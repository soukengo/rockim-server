package cache

import "context"

type Cache[T any] interface {
	Exists(ctx context.Context) (bool, error)
	Delete(ctx context.Context) error
}

type ValueCache[T any] interface {
	Cache[T]
	Get(ctx context.Context) (*T, error)
	Set(ctx context.Context, value *T) error
}

type HashCache[T any] interface {
	Cache[T]
	Get(ctx context.Context) (*T, error)
	GetAll(ctx context.Context) (map[string]*T, error)
	Put(ctx context.Context, value *T) error
	PutAll(ctx context.Context, hash map[string]*T) error
	ExistsKey(ctx context.Context) (bool, error)
	DeleteOne(ctx context.Context) error
}

type SetCache[T any] interface {
	Cache[T]
	Add(ctx context.Context, value *T) error
	DeleteOne(ctx context.Context, value *T) error
	ExistsMember(ctx context.Context, value *T) (bool, error)
	Count(ctx context.Context) (int64, error)
}

type SortedSetCache[T Codec] interface {
	Cache[T]
	Add(ctx context.Context, score float64, value *T) error
	DeleteOne(ctx context.Context, value *T) error
	ExistsMember(ctx context.Context, value *T) (bool, error)
	Count(ctx context.Context) (int64, error)
}
