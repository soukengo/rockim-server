package cache

import "context"

type Cache[T any] interface {
	Exists(ctx context.Context, keySuffix string) (bool, error)
	Delete(ctx context.Context, keySuffix string) error
}

type ValueCache[T any] interface {
	Cache[T]
	Get(ctx context.Context, keySuffix string) (*T, error)
	Set(ctx context.Context, keySuffix string, value *T) error
}

type HashCache[T any] interface {
	Cache[T]
	Put(ctx context.Context, keySuffix string, field string, value *T) error
	PutMap(ctx context.Context, keySuffix string, hash map[string]*T) error
	Get(ctx context.Context, keySuffix string, field string) (*T, error)
	GetAll(ctx context.Context, keySuffix string) (map[string]*T, error)
	ExistsField(ctx context.Context, keySuffix string, field string) (bool, error)
	DeleteField(ctx context.Context, keySuffix string, field string) error
}

type SetCache[T any] interface {
	Cache[T]
	Add(ctx context.Context, keySuffix string, value *T) error
	AddSlice(ctx context.Context, keySuffix string, values []*T) error
	Paginate(ctx context.Context, keySuffix string, cursor uint64, count int64) ([]*T, uint64, error)
	DeleteMember(ctx context.Context, keySuffix string, value *T) error
	ExistsMember(ctx context.Context, keySuffix string, value *T) (bool, error)
	Count(ctx context.Context, keySuffix string) (int64, error)
}

type SortedSetCache[T any] interface {
	Cache[T]
	Add(ctx context.Context, keySuffix string, pairs *SortedMember[T]) error
	AddSlice(ctx context.Context, keySuffix string, values []*SortedMember[T]) error
	Paginate(ctx context.Context, keySuffix string, cursor uint64, count int64) ([]*T, uint64, error)
	DeleteMember(ctx context.Context, keySuffix string, value *T) error
	ExistsMember(ctx context.Context, keySuffix string, value *T) (bool, error)
	Count(ctx context.Context, keySuffix string) (int64, error)
}

type SortedMember[T any] struct {
	Score float64
	Value *T
}
