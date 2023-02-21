package biz

import (
	"context"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/pkg/errors"
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
	GenID(ctx context.Context) (string, error)
	Create(context.Context, *types.User) error
	Update(context.Context, *types.User) error
}

// UserUseCase is a User use case.
type UserUseCase struct {
	repo UserRepo
}

// NewUserUseCase new a User use case.
func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

// Register register a new user
func (uc *UserUseCase) Register(ctx context.Context, u *types.User) (*types.User, error) {
	// TODO: 分布式锁
	// 检查是否已经注册
	exists, err := uc.existsAccount(ctx, u.ProductId, u.Account)
	if err != nil {
		return nil, err
	}
	// 账号已注册
	if exists {
		return nil, ErrAccountRegistered
	}
	uid, err := uc.repo.GenID(ctx)
	if err != nil {
		return nil, err
	}
	u.Id = uid
	u.CreateTime = time.Now().UnixMilli()
	u.Status = enums.UserStatus_USER_STATUS_NORMAL
	err = uc.repo.Create(ctx, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *UserUseCase) Find(ctx context.Context, productId string, account string) (ret *types.User, err error) {
	uid, err := uc.repo.FindUidByAccount(ctx, productId, account)
	if err != nil {
		if errors.IsNotFound(err) {
			err = ErrUserNotFound
		}
		return
	}
	return uc.repo.FindByID(ctx, productId, uid)
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
