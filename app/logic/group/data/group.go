package data

import (
	"context"
	"github.com/soukengo/gopkg/errors"
	"github.com/soukengo/gopkg/util/chain"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/app/logic/group/biz"
	"rockimserver/app/logic/group/data/cache"
	"rockimserver/app/logic/group/data/database"
)

type groupRepo struct {
	db    *database.GroupData
	cache *cache.GroupData
}

func NewGroupRepo(db *database.GroupData, cache *cache.GroupData) biz.GroupRepo {
	return &groupRepo{db: db, cache: cache}
}

func (r *groupRepo) FindGroupId(ctx context.Context, productId string, bizId string) (groupId string, err error) {
	// 链式调用，先查缓存，再查数据库，最后回写缓存
	groupId, err = chain.CallWithResult[string](func() (ret string, cont bool, err error) {
		ret, err = r.cache.FindGroupId(ctx, productId, bizId)
		// 正常查下到结果|没有缓存,不再查询数据库
		cont = cache.IsErrNoCache(err)
		return
	}, func() (ret string, cont bool, err error) {
		ret, err = r.db.FindGroupId(ctx, productId, bizId)
		if err == nil || errors.IsNotFound(err) {
			_ = r.cache.SaveBizId(ctx, productId, bizId, ret)
		}
		return
	})
	if errors.IsNotFound(err) {
		err = biz.ErrGroupNotFound
	}
	return
}

func (r *groupRepo) FindById(ctx context.Context, productId string, groupId string) (ret *types.Group, err error) {
	// 链式调用，先查缓存，再查数据库，最后回写缓存
	ret, err = chain.CallWithResult[*types.Group](func() (ret *types.Group, cont bool, err error) {
		ret, err = r.cache.FindByID(ctx, productId, groupId)
		// 正常查下到结果|没有缓存,不再查询数据库
		cont = cache.IsErrNoCache(err)
		return
	}, func() (ret *types.Group, cont bool, err error) {
		ret, err = r.db.FindByID(ctx, productId, groupId)
		if err == nil || errors.IsNotFound(err) {
			_ = r.cache.SaveGroup(ctx, productId, groupId, ret)
		}
		return
	})
	if errors.IsNotFound(err) {
		err = biz.ErrGroupNotFound
	}
	return
}

func (r *groupRepo) Create(ctx context.Context, group *types.Group) (err error) {
	err = r.db.Create(ctx, group)
	if err != nil {
		return
	}
	_ = r.cache.DeleteGroup(ctx, group.ProductId, group.Id)
	_ = r.cache.DeleteBizId(ctx, group.ProductId, group.BizId)
	return
}

func (r *groupRepo) Delete(ctx context.Context, group *types.Group) (err error) {
	err = r.db.Delete(ctx, group.ProductId, group.Id)
	if err != nil {
		return
	}
	_ = r.cache.DeleteGroup(ctx, group.ProductId, group.Id)
	_ = r.cache.DeleteBizId(ctx, group.ProductId, group.BizId)
	return
}
