package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/pkg/log"
)

var (
	ErrPlatUserNotFound = errors.NotFound(v1.ErrorReason_PLAT_USER_NOT_FOUND.String(), "plat account not found")
)

// PlatUserRepo is a User repo.
type PlatUserRepo interface {
	FindById(ctx context.Context, id string) (user *types.PlatUser, err error)
	FindIdByAccount(ctx context.Context, account string) (uid string, err error)
	Create(ctx context.Context, user *types.PlatUser) error
}

// PlatUserUseCase is a User use case.
type PlatUserUseCase struct {
	repo PlatUserRepo
}

// NewPlatUserUseCase new a User use case.
func NewPlatUserUseCase(repo PlatUserRepo) *PlatUserUseCase {
	return &PlatUserUseCase{repo: repo}
}

func (uc *PlatUserUseCase) Find(ctx context.Context, account string) (user *types.PlatUser, err error) {
	uid, err := uc.repo.FindIdByAccount(ctx, account)
	if err != nil {
		log.Errorf("FindIdByAccount account: %v error: %v", account, err)
		err = ErrPlatUserNotFound
		return
	}
	return uc.repo.FindById(ctx, uid)
}
