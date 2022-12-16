package service

import (
	"context"
	"rockim/api/rockim/service/platform/v1"
	"rockim/app/logic/platform/biz"
	"rockim/pkg/util/copier"
)

type PlatUserService struct {
	v1.UnimplementedPlatUserAPIServer
	uc *biz.PlatUserUseCase
}

// NewPlatUserService new a plat user service.
func NewPlatUserService(uc *biz.PlatUserUseCase) *PlatUserService {
	return &PlatUserService{uc: uc}
}

func (s *PlatUserService) Create(ctx context.Context, in *v1.PlatUserCreateRequest) (reply *v1.PlatUserCreateResponse, err error) {
	req := &biz.PlatUserCreateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.PlatUserCreateResponse{}
	return
}

func (s *PlatUserService) Update(ctx context.Context, in *v1.PlatUserUpdateRequest) (reply *v1.PlatUserUpdateResponse, err error) {
	req := &biz.PlatUserUpdateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.PlatUserUpdateResponse{}
	return
}

func (s *PlatUserService) Delete(ctx context.Context, in *v1.PlatUserDeleteRequest) (reply *v1.PlatUserDeleteResponse, err error) {
	req := &biz.PlatUserDeleteRequest{Id: in.Id}
	err = s.uc.Delete(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.PlatUserDeleteResponse{}
	return
}

func (s *PlatUserService) Paging(ctx context.Context, in *v1.PlatUserPagingRequest) (reply *v1.PlatUserPagingResponse, err error) {
	req := &biz.PlatUserPagingRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	res, err := s.uc.Paging(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.PlatUserPagingResponse{List: res.List, Paginate: res.Paginate}
	return
}

func (s *PlatUserService) Find(ctx context.Context, req *v1.PlatUserFindRequest) (*v1.PlatUserFindResponse, error) {
	u, err := s.uc.Find(ctx, req.Account)
	if err != nil {
		return nil, err
	}
	return &v1.PlatUserFindResponse{User: u}, nil
}

func (s *PlatUserService) ListRoleId(ctx context.Context, request *v1.PlatUserRoleIdListRequest) (*v1.PlatUserRoleIdListResponse, error) {
	list, err := s.uc.ListRoleId(ctx, request.UserIds)
	if err != nil {
		return nil, err
	}
	return &v1.PlatUserRoleIdListResponse{List: list}, nil
}
