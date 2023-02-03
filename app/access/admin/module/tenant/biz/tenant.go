package biz

import (
	"context"
	"rockimserver/apis/rockim/service/platform/v1/types"
)

type TenantRepo interface {
	FindByEmail(ctx context.Context, email string) (*types.Tenant, error)
}
