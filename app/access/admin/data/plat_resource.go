package data

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/access/admin/biz/manager"
)

type platResourceRepo struct {
	uac v1.PlatResourceAPIClient
}

func NewPlatResourceRepo(uac v1.PlatResourceAPIClient) manager.PlatResourceRepo {
	return &platResourceRepo{uac: uac}
}

func (r *platResourceRepo) ListByIds(ctx context.Context, ids []string) ([]*types.PlatResource, error) {
	ret, err := r.uac.ListByIds(ctx, &v1.PlatResourceListByIdsRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	return ret.List, nil
}
