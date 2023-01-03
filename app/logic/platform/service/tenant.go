package service

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/app/logic/platform/biz"
	"rockim/pkg/util/copier"
)

type TenantService struct {
	v1.UnimplementedTenantAPIServer
	uc *biz.TenantUseCase
}

func NewTenantService(uc *biz.TenantUseCase) *TenantService {
	return &TenantService{uc: uc}
}

func (s *TenantService) Create(ctx context.Context, in *v1.TenantCreateRequest) (reply *v1.TenantCreateResponse, err error) {
	req := &biz.TenantCreateRequest{}
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
	req := &biz.TenantUpdateRequest{}
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
	req := &biz.TenantPagingRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	res, err := s.uc.Paging(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.TenantPagingResponse{List: res.List, Paginate: res.Paginate}
	return
}

func (s *TenantService) Find(ctx context.Context, in *v1.TenantFindRequest) (*v1.TenantFindResponse, error) {
	record, err := s.uc.Find(ctx, in.Email)
	if err != nil {
		return nil, err
	}
	return &v1.TenantFindResponse{Tenant: record}, nil
}
