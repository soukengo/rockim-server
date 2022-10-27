package biz

import (
	"context"
	"time"

	v1 "rockim/api/logic/user/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound      = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
	ErrAccountRegistered = errors.NotFound(v1.ErrorReason_ACCOUNT_REGISTERED.String(), "account already registered")
)

// UserRepo is a User repo.
type UserRepo interface {
	FindByID(ctx context.Context, appId string, id string) (*User, error)
	FindByAccount(ctx context.Context, appId string, account string) (uid string, err error)
	GenID() (string, error)
	Save(context.Context, *User) (*User, error)
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
func (uc *UserUseCase) Register(ctx context.Context, u *User) (*User, error) {
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
	uid, err := uc.repo.GenID()
	if err != nil {
		return nil, err
	}
	u.Id = uid
	u.CreateTime = time.Now().UnixMilli()
	u.Status = UserStatus_USER_STATUS_NORMAL
	return uc.repo.Save(ctx, u)
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
