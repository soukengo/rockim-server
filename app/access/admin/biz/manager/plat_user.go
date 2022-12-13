package manager

import (
	"context"
	"rockim/api/rockim/service/platform/v1/types"
)

type PlatUserRepo interface {
	FindByAccount(ctx context.Context, account string) (*types.PlatUser, error)
	ListRoleId(ctx context.Context, userId string) ([]string, error)
}
