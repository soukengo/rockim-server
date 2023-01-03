package tenant

import (
	"context"
	"rockim/api/rockim/service/platform/v1/types"
)

type TenantResourceRepo interface {
	List(context.Context) ([]*types.TenantResource, error)
}
