package data

import (
	"context"
	"github.com/soukengo/gopkg/errors"
	"github.com/soukengo/gopkg/util/chain"
	"rockimserver/apis/rockim/service/platform/v1/types"
	"rockimserver/app/logic/platform/biz"
	"rockimserver/app/logic/platform/data/cache"
	"rockimserver/app/logic/platform/data/database"
)

type productRepo struct {
	db    *database.ProductData
	cache *cache.ProductData
}

func NewProductRepo(db *database.ProductData, cache *cache.ProductData) biz.ProductRepo {
	return &productRepo{db: db, cache: cache}
}

func (r *productRepo) GenID(ctx context.Context) (string, error) {
	return r.db.GenID(ctx)
}

func (r *productRepo) Create(ctx context.Context, record *types.Product) (err error) {
	err = r.db.Create(ctx, record)
	if err != nil {
		return
	}
	_ = r.cache.Delete(ctx, record.Id)
	return
}

func (r *productRepo) Update(ctx context.Context, record *types.Product) error {
	return r.db.Update(ctx, record)
}

func (r *productRepo) FindById(ctx context.Context, id string) (out *types.Product, err error) {
	// 链式调用，先查缓存，再查数据库，最后回写缓存
	out, err = chain.CallWithResult[*types.Product](func() (ret *types.Product, cont bool, err error) {
		ret, err = r.cache.FindByID(ctx, id)
		// 正常查下到结果|没有缓存,不再查询数据库
		cont = cache.IsErrNoCache(err)
		return
	}, func() (ret *types.Product, cont bool, err error) {
		ret, err = r.db.FindByID(ctx, id)
		if err == nil || errors.IsNotFound(err) {
			_ = r.cache.Save(ctx, id, ret)
		}
		return
	})
	if errors.IsNotFound(err) {
		err = biz.ErrProductNotFound
	}
	return
}

func (r *productRepo) ListByIds(ctx context.Context, ids []string) ([]*types.Product, error) {
	return r.db.ListByIds(ctx, ids)
}

func (r *productRepo) ListByTenant(ctx context.Context, tenantId string) (res []*types.Product, err error) {
	return r.db.ListByTenant(ctx, tenantId)
}
