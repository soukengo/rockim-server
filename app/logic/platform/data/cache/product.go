package cache

import (
	"context"
	"github.com/soukengo/gopkg/component/cache"
	"rockimserver/apis/rockim/service/platform/v1/types"
)

type ProductData struct {
	cache cache.ValueCache[types.Product]
}

func NewProductData(mgr *cache.Manager, cfg *cache.Config) *ProductData {
	return &ProductData{
		cache: cache.NewValueCache[types.Product](mgr, cfg.Category(keyProduct)),
	}
}

func (d *ProductData) FindByID(ctx context.Context, productId string) (*types.Product, error) {
	val, err := d.cache.Get(ctx, cache.Parts(productId))
	return val, err
}
func (d *ProductData) Save(ctx context.Context, productId string, record *types.Product) error {
	return d.cache.Set(ctx, cache.Parts(productId), record)
}

func (d *ProductData) Delete(ctx context.Context, productId string) error {
	return d.cache.Delete(ctx, cache.Parts(productId))
}
