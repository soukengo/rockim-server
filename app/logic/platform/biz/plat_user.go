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

// PlatUserRepo is a PlatUser repo.
type PlatUserRepo interface {
	Create(ctx context.Context, user *types.PlatUser) error
	Update(ctx context.Context, user *types.PlatUser) error

	FindById(ctx context.Context, id string) (user *types.PlatUser, err error)
	FindIdByAccount(ctx context.Context, account string) (uid string, err error)
	ListRoleId(ctx context.Context, ids []string) ([]string, error)
}

// PlatUserUseCase is a PlatUser use case.
type PlatUserUseCase struct {
	repo PlatUserRepo
}

// NewPlatUserUseCase new a PlatUser use case.
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
func (uc *PlatUserUseCase) ListRoleId(ctx context.Context, ids []string) (res []string, err error) {
	return uc.repo.ListRoleId(ctx, ids)
}
