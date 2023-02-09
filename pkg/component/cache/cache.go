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
	Put(ctx context.Context, field string, value *T) error
	PutMap(ctx context.Context, hash map[string]*T) error
	Get(ctx context.Context, field string) (*T, error)
	GetAll(ctx context.Context) (map[string]*T, error)
	ExistsField(ctx context.Context, field string) (bool, error)
	DeleteField(ctx context.Context, field string) error
}

type SetCache[T any] interface {
	Cache[T]
	Add(ctx context.Context, value *T) error
	AddSlice(ctx context.Context, values []*T) error
	Paginate(ctx context.Context, cursor uint64, count int64) ([]*T, uint64, error)
	DeleteMember(ctx context.Context, value *T) error
	ExistsMember(ctx context.Context, value *T) (bool, error)
	Count(ctx context.Context) (int64, error)
}

type SortedSetCache[T any] interface {
	Cache[T]
	Add(ctx context.Context, pairs *SortedMember[T]) error
	AddSlice(ctx context.Context, values []*SortedMember[T]) error
	Paginate(ctx context.Context, cursor uint64, count int64) ([]*T, uint64, error)
	DeleteMember(ctx context.Context, value *T) error
	ExistsMember(ctx context.Context, value *T) (bool, error)
	Count(ctx context.Context) (int64, error)
}

type SortedMember[T any] struct {
	Score float64
	Value *T
}
