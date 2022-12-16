package manager

import (
	"context"
	adminV1 "rockim/api/rockim/admin/manager/v1"
	adminTypes "rockim/api/rockim/admin/manager/v1/types"
	"rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/access/admin/biz/manager"
	"rockim/pkg/util/copier"
)

type PlatUserService struct {
	uc *manager.PlatUserUseCase
}

func NewPlatUserService(uc *manager.PlatUserUseCase) *PlatUserService {
	return &PlatUserService{uc: uc}
}

func (s *PlatUserService) Create(ctx context.Context, in *adminV1.PlatUserCreateRequest) (reply *adminV1.PlatUserCreateResponse, err error) {
	req := &v1.PlatUserCreateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.PlatUserCreateResponse{}
	return
}

func (s *PlatUserService) Update(ctx context.Context, in *adminV1.PlatUserUpdateRequest) (reply *adminV1.PlatUserUpdateResponse, err error) {
	req := &v1.PlatUserUpdateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.PlatUserUpdateResponse{}
	return
}

func (s *PlatUserService) Delete(ctx context.Context, in *adminV1.PlatUserDeleteRequest) (reply *adminV1.PlatUserDeleteResponse, err error) {
	req := &v1.PlatUserDeleteRequest{Id: in.Id}
	err = s.uc.Delete(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.PlatUserDeleteResponse{}
	return
}

func (s *PlatUserService) Paging(ctx context.Context, in *adminV1.PlatUserPagingRequest) (reply *adminV1.PlatUserPagingResponse, err error) {
	req := &v1.PlatUserPagingRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	ret, err := s.uc.Paging(ctx, req)
	if err != nil {
		return
	}
	list := make([]*adminTypes.PlatUser, len(ret.List))
	for i, item := range ret.List {
		list[i] = convertUser(item)
	}
	return &adminV1.PlatUserPagingResponse{List: list, Paginate: ret.Paginate}, nil
}

func (s *PlatUserService) ListRoleId(ctx context.Context, in *adminV1.PlatUserRoleIdListRequest) (reply *adminV1.PlatUserRoleIdListResponse, err error) {
	list, err := s.uc.ListRoleId(ctx, in.UserId)
	if err != nil {
		return
	}
	return &adminV1.PlatUserRoleIdListResponse{List: list}, nil
}

func convertUser(source *types.PlatUser) *adminTypes.PlatUser {
	return &adminTypes.PlatUser{
		Id:         source.Id,
		CreateTime: source.CreateTime,
		Account:    source.Account,
		Name:       source.Name,
	}
}
