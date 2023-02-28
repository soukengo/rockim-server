package cache

import "context"

type Cache[T any] interface {
	Exists(ctx context.Context, parts KeyParts) (bool, error)
	Delete(ctx context.Context, parts KeyParts) error
}

type ValueCache[T any] interface {
	Cache[T]
	Get(ctx context.Context, parts KeyParts) (*T, error)
	Set(ctx context.Context, parts KeyParts, value *T) error
}

type HashCache[T any] interface {
	Cache[T]
	Put(ctx context.Context, parts KeyParts, field string, value *T) error
	PutMap(ctx context.Context, parts KeyParts, hash map[string]*T) error
	Get(ctx context.Context, parts KeyParts, field string) (*T, error)
	GetAll(ctx context.Context, parts KeyParts) (map[string]*T, error)
	ExistsField(ctx context.Context, parts KeyParts, field string) (bool, error)
	DeleteField(ctx context.Context, parts KeyParts, field string) error
}

type SetCache[T any] interface {
	Cache[T]
	Add(ctx context.Context, parts KeyParts, value *T) error
	AddSlice(ctx context.Context, parts KeyParts, values []*T) error
	Paginate(ctx context.Context, parts KeyParts, cursor uint64, count int64) ([]*T, uint64, error)
	DeleteMember(ctx context.Context, parts KeyParts, value *T) error
	ExistsMember(ctx context.Context, parts KeyParts, value *T) (bool, error)
	Count(ctx context.Context, parts KeyParts) (int64, error)
}

type SortedSetCache[T any] interface {
	Cache[T]
	Add(ctx context.Context, parts KeyParts, member *SortedMember[T]) error
	AddSlice(ctx context.Context, parts KeyParts, values []*SortedMember[T]) error
	Paginate(ctx context.Context, parts KeyParts, cursor uint64, count int64) ([]*T, uint64, error)
	DeleteMember(ctx context.Context, parts KeyParts, value *T) error
	ExistsMember(ctx context.Context, parts KeyParts, value *T) (bool, error)
	Count(ctx context.Context, parts KeyParts) (int64, error)
}

type SortedMember[T any] struct {
	Score float64
	Value *T
}
