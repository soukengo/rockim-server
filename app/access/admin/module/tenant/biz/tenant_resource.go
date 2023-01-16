package biz

import (
	"context"
	"rockim/app/access/admin/module/tenant/biz/types"
)

type SysTenantResourceRepo interface {
	List(context.Context) ([]*types.SysTenantResource, error)
}
