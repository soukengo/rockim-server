package data

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/access/admin/biz/manager"
)

type platUserRepo struct {
	uac v1.PlatUserAPIClient
}

func NewPlatUserRepo(uac v1.PlatUserAPIClient) manager.PlatUserRepo {
	return &platUserRepo{uac: uac}
}

func (r *platUserRepo) FindByAccount(ctx context.Context, account string) (*types.PlatUser, error) {
	ret, err := r.uac.Find(ctx, &v1.PlatUserFindRequest{Account: account})
	if err != nil {
		return nil, err
	}
	return ret.User, nil
}

func (r *platUserRepo) ListRoleId(ctx context.Context, userId string) ([]string, error) {
	ret, err := r.uac.ListRoleId(ctx, &v1.PlatUserRoleIdListRequest{
		UserIds: []string{userId},
	})
	if err != nil {
		return nil, err
	}
	return ret.List, nil
}
