package data

import (
	"context"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/app/logic/user/biz"
	"rockimserver/app/logic/user/data/cache"
	"rockimserver/app/logic/user/data/database"
	"rockimserver/pkg/errors"
	"rockimserver/pkg/util/chain"
)

type userRepo struct {
	db    *database.UserData
	cache *cache.UserData
}

// NewUserRepo .
func NewUserRepo(db *database.UserData, cache *cache.UserData) biz.UserRepo {
	return &userRepo{
		db:    db,
		cache: cache,
	}
}

func (r *userRepo) FindByID(ctx context.Context, productId string, uid string) (ret *types.User, err error) {
	// 链式调用，先查缓存，再查数据库，最后回写缓存
	ret, err = chain.Call[*types.User](func() (ret *types.User, cont bool, err error) {
		ret, err = r.cache.FindUserByID(ctx, productId, uid)
		// 正常查下到结果|没有缓存,不再查询数据库
		cont = cache.IsErrNoCache(err)
		return
	}, func() (ret *types.User, cont bool, err error) {
		ret, err = r.db.FindByID(ctx, productId, uid)
		if err == nil || errors.IsNotFound(err) {
			_ = r.cache.SaveUser(ctx, productId, uid, ret)
		}
		return
	})
	if errors.IsNotFound(err) {
		err = biz.ErrUserNotFound
	}
	return
}

func (r *userRepo) FindUidByAccount(ctx context.Context, productId string, account string) (uid string, err error) {
	// 链式调用，先查缓存，再查数据库，最后回写缓存
	uid, err = chain.Call[string](func() (ret string, cont bool, err error) {
		ret, err = r.cache.FindUidByAccount(ctx, productId, account)
		// 正常查下到结果|没有缓存,不再查询数据库
		cont = cache.IsErrNoCache(err)
		return
	}, func() (ret string, cont bool, err error) {
		ret, err = r.db.FindUidByAccount(ctx, productId, account)
		if err == nil || errors.IsNotFound(err) {
			_ = r.cache.SaveAccountUid(ctx, productId, account, uid)
		}
		return
	})
	if errors.IsNotFound(err) {
		err = biz.ErrUserNotFound
	}
	return
}

func (r *userRepo) GenID(ctx context.Context) (string, error) {
	return r.db.GenID(ctx)
}

func (r *userRepo) Create(ctx context.Context, user *types.User) (err error) {
	err = r.db.Create(ctx, user)
	if err != nil {
		return
	}
	_ = r.cache.DeleteAccountUid(ctx, user.ProductId, user.Account)
	return
}
func (r *userRepo) Update(ctx context.Context, user *types.User) (err error) {
	err = r.db.Update(ctx, user)
	if err != nil {
		return
	}
	_ = r.cache.DeleteAccountUid(ctx, user.ProductId, user.Account)
	return
}
