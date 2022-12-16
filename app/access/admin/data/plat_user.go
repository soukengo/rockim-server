package data

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/access/admin/biz/manager"
)

type platUserRepo struct {
	ac v1.PlatUserAPIClient
}

func NewPlatUserRepo(uac v1.PlatUserAPIClient) manager.PlatUserRepo {
	return &platUserRepo{ac: uac}
}

func (r *platUserRepo) Create(ctx context.Context, request *v1.PlatUserCreateRequest) (err error) {
	_, err = r.ac.Create(ctx, request)
	return
}

func (r *platUserRepo) Update(ctx context.Context, request *v1.PlatUserUpdateRequest) (err error) {
	_, err = r.ac.Update(ctx, request)
	return
}

func (r *platUserRepo) Delete(ctx context.Context, request *v1.PlatUserDeleteRequest) (err error) {
	_, err = r.ac.Delete(ctx, request)
	return
}

func (r *platUserRepo) Paging(ctx context.Context, request *v1.PlatUserPagingRequest) (reply *v1.PlatUserPagingResponse, err error) {
	return r.ac.Paging(ctx, request)
}

func (r *platUserRepo) FindByAccount(ctx context.Context, account string) (*types.PlatUser, error) {
	ret, err := r.ac.Find(ctx, &v1.PlatUserFindRequest{Account: account})
	if err != nil {
		return nil, err
	}
	return ret.User, nil
}

func (r *platUserRepo) ListRoleId(ctx context.Context, userId string) ([]string, error) {
	ret, err := r.ac.ListRoleId(ctx, &v1.PlatUserRoleIdListRequest{
		UserIds: []string{userId},
	})
	if err != nil {
		return nil, err
	}
	return ret.List, nil
}
