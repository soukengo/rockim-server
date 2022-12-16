package data

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/access/admin/biz/manager"
)

type platResourceRepo struct {
	ac v1.PlatResourceAPIClient
}

func NewPlatResourceRepo(ac v1.PlatResourceAPIClient) manager.PlatResourceRepo {
	return &platResourceRepo{ac: ac}
}

func (r *platResourceRepo) Create(ctx context.Context, request *v1.PlatResourceCreateRequest) (err error) {
	_, err = r.ac.Create(ctx, request)
	return
}

func (r *platResourceRepo) Update(ctx context.Context, request *v1.PlatResourceUpdateRequest) (err error) {
	_, err = r.ac.Update(ctx, request)
	return
}

func (r *platResourceRepo) Delete(ctx context.Context, request *v1.PlatResourceDeleteRequest) (err error) {
	_, err = r.ac.Delete(ctx, request)
	return
}

func (r *platResourceRepo) List(ctx context.Context) ([]*types.PlatResource, error) {
	ret, err := r.ac.List(ctx, &v1.PlatResourceListRequest{})
	if err != nil {
		return nil, err
	}
	return ret.List, nil
}
func (r *platResourceRepo) ListByIds(ctx context.Context, ids []string) ([]*types.PlatResource, error) {
	ret, err := r.ac.ListByIds(ctx, &v1.PlatResourceListByIdsRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	return ret.List, nil
}
