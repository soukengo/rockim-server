package biz

import (
	"context"
	"rockim/api/rockim/service/platform/v1/types"
)

type TenantRepo interface {
	FindByEmail(ctx context.Context, email string) (*types.Tenant, error)
}
