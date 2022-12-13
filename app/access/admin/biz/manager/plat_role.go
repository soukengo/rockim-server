package manager

import (
	"context"
)

type PlatRoleRepo interface {
	ListResourceId(ctx context.Context, roleIds []string) ([]string, error)
}
