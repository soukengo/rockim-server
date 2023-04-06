package biz

import (
	"context"
	"github.com/soukengo/gopkg/component/idgen"
	"github.com/soukengo/gopkg/component/lock"
	"github.com/soukengo/gopkg/errors"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/logic/user/biz/consts"
	"rockimserver/app/logic/user/biz/options"
	"time"

	"rockimserver/apis/rockim/service/user/v1/types"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound      = errors.NotFound(reasons.User_USER_NOT_FOUND.String(), "user not found")
	ErrAccountRegistered = errors.Conflict(reasons.User_ACCOUNT_REGISTERED.String(), "account already registered")
)

// UserRepo is a User repo.
type UserRepo interface {
	FindByID(ctx context.Context, productId string, id string) (*types.User, error)
	FindUidByAccount(ctx context.Context, productId string, account string) (uid string, err error)
	Create(context.Context, *types.User) error
	Update(context.Context, *types.User) error
}

// UserUseCase is a User use case.
type UserUseCase struct {
	repo    UserRepo
	lockBdr lock.Builder
	idgen   idgen.Generator
}

// NewUserUseCase new a User use case.
func NewUserUseCase(repo UserRepo, idgen idgen.Generator, lockBdr lock.Builder) *UserUseCase {
	return &UserUseCase{repo: repo, lockBdr: lockBdr, idgen: idgen}
}

// Register register a new user
func (uc *UserUseCase) Register(ctx context.Context, opts *options.UserCreateOptions) (out *types.User, err error) {
	productId := opts.ProductId
	account := opts.Account
	distributedLock := uc.lockBdr.Build(consts.LockKeyUserCreate, productId, account)
	locked := distributedLock.TryLock(ctx)
	defer distributedLock.UnLock()
	if !locked {
		return
	}
	// 检查是否已经注册
	exists, err := uc.existsAccount(ctx, productId, account)
	if err != nil {
		return nil, err
	}
	// 账号已注册
	if exists {
		return nil, ErrAccountRegistered
	}
	uid, err := uc.idgen.GenID()
	if err != nil {
		return nil, err
	}
	u := &types.User{
		Id:         uid,
		CreateTime: time.Now().UnixMilli(),
		ProductId:  opts.ProductId,
		Account:    opts.Account,
		Name:       opts.Name,
		AvatarUrl:  opts.AvatarUrl,
		Fields:     opts.Fields,
		Status:     enums.User_NORMAL,
	}
	err = uc.repo.Create(ctx, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *UserUseCase) Find(ctx context.Context, productId string, uid string) (ret *types.User, err error) {
	return uc.repo.FindByID(ctx, productId, uid)
}

func (uc *UserUseCase) FindByAccount(ctx context.Context, productId string, account string) (ret *types.User, err error) {
	uid, err := uc.repo.FindUidByAccount(ctx, productId, account)
	if err != nil {
		return
	}
	return uc.repo.FindByID(ctx, productId, uid)
}
func (uc *UserUseCase) FindUid(ctx context.Context, productId string, account string) (ret string, err error) {
	return uc.repo.FindUidByAccount(ctx, productId, account)
}

func (uc *UserUseCase) existsAccount(ctx context.Context, productId string, account string) (exists bool, err error) {
	uid, err := uc.repo.FindUidByAccount(ctx, productId, account)
	if err != nil {
		if errors.IsNotFound(err) {
			err = nil
		}
		return
	}
	exists = len(uid) > 0
	return
}
