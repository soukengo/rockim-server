package manager

import (
	"context"
	"rockim/api/rockim/service/platform/v1/types"
)

type PlatResourceRepo interface {
	ListByIds(ctx context.Context, ids []string) ([]*types.PlatResource, error)
}
