package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"rockim/api/rockim/service/user/v1"
	"rockim/api/rockim/service/user/v1/types"
	v1enums "rockim/api/rockim/shared/enums/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound      = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
	ErrAccountRegistered = errors.Conflict(v1.ErrorReason_ACCOUNT_REGISTERED.String(), "account already registered")
)

// UserRepo is a User repo.
type UserRepo interface {
	FindByID(ctx context.Context, appId string, id string) (*types.User, error)
	FindByAccount(ctx context.Context, appId string, account string) (uid string, err error)
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
	exists, err := uc.existsAccount(ctx, u.AppId, u.Account)
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
	u.Status = v1enums.UserStatus_USER_STATUS_NORMAL
	err = uc.repo.Create(ctx, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *UserUseCase) existsAccount(ctx context.Context, appId string, account string) (exists bool, err error) {
	uid, err := uc.repo.FindByAccount(ctx, appId, account)
	if err != nil {
		if errors.IsNotFound(err) {
			err = nil
		}
		return
	}
	exists = len(uid) > 0
	return
}
