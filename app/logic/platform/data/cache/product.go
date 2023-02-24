package cache

import (
	"context"
	"rockimserver/apis/rockim/service/platform/v1/types"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
)

type ProductData struct {
	cache cache.ValueCache[types.Product]
}

func NewProductData(redisCli *redis.Client, cfg *cache.Config) *ProductData {
	return &ProductData{
		cache: newValueCache[types.Product](redisCli, cfg.Category(keyProduct)),
	}
}

func (d *ProductData) FindByID(ctx context.Context, productId string) (*types.Product, error) {
	val, err := d.cache.Get(ctx, productId)
	return val, err
}
func (d *ProductData) Save(ctx context.Context, productId string, record *types.Product) error {
	return d.cache.Set(ctx, productId, record)
}

func (d *ProductData) Delete(ctx context.Context, productId string) error {
	return d.cache.Delete(ctx, productId)
}
