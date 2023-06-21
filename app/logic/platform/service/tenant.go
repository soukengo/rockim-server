package service

import (
	"context"
	"github.com/soukengo/gopkg/util/copier"
	v1 "rockimserver/apis/rockim/service/platform/v1"
	"rockimserver/apis/rockim/shared"
	"rockimserver/app/logic/platform/biz"
	"rockimserver/app/logic/platform/biz/options"
)

type TenantService struct {
	v1.UnimplementedTenantAPIServer
	uc *biz.TenantUseCase
}

func NewTenantService(uc *biz.TenantUseCase) *TenantService {
	return &TenantService{uc: uc}
}

func (s *TenantService) Create(ctx context.Context, in *v1.TenantCreateRequest) (reply *v1.TenantCreateResponse, err error) {
	req := &options.TenantCreateOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.TenantCreateResponse{}
	return
}

func (s *TenantService) Update(ctx context.Context, in *v1.TenantUpdateRequest) (reply *v1.TenantUpdateResponse, err error) {
	req := &options.TenantUpdateOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.TenantUpdateResponse{}
	return
}

func (s *TenantService) Paging(ctx context.Context, in *v1.TenantPagingRequest) (reply *v1.TenantPagingResponse, err error) {
	req := &options.TenantPagingOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	if req.Paginate == nil {
		req.Paginate = &shared.Paginating{}
	}
	list, paginate, err := s.uc.Paging(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.TenantPagingResponse{List: list, Paginate: paginate}
	return
}

func (s *TenantService) Find(ctx context.Context, in *v1.TenantFindRequest) (*v1.TenantFindResponse, error) {
	record, err := s.uc.Find(ctx, in.Email)
	if err != nil {
		return nil, err
	}
	return &v1.TenantFindResponse{Tenant: record}, nil
}
