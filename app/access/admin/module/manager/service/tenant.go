package service

import (
	"context"
	"github.com/soukengo/gopkg/util/copier"
	adminV1 "rockimserver/apis/rockim/api/admin/manager/v1"
	apiTypes "rockimserver/apis/rockim/api/admin/manager/v1/types"
	"rockimserver/apis/rockim/service/platform/v1"
	"rockimserver/apis/rockim/service/platform/v1/types"
	"rockimserver/app/access/admin/module/manager/biz"
)

type TenantService struct {
	uc *biz.TenantUseCase
}

func NewTenantService(uc *biz.TenantUseCase) *TenantService {
	return &TenantService{uc: uc}
}

func (s *TenantService) Create(ctx context.Context, in *adminV1.TenantCreateRequest) (reply *adminV1.TenantCreateResponse, err error) {
	req := &v1.TenantCreateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.TenantCreateResponse{}
	return
}

func (s *TenantService) Update(ctx context.Context, in *adminV1.TenantUpdateRequest) (reply *adminV1.TenantUpdateResponse, err error) {
	req := &v1.TenantUpdateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.TenantUpdateResponse{}
	return
}

func (s *TenantService) Paging(ctx context.Context, in *adminV1.TenantPagingRequest) (reply *adminV1.TenantPagingResponse, err error) {
	req := &v1.TenantPagingRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	ret, err := s.uc.Paging(ctx, req)
	if err != nil {
		return
	}
	list := make([]*apiTypes.Tenant, len(ret.List))
	for i, item := range ret.List {
		list[i] = convertTenant(item)
	}
	return &adminV1.TenantPagingResponse{List: list, Paginate: ret.Paginate}, nil
}

func convertTenant(source *types.Tenant) *apiTypes.Tenant {
	return &apiTypes.Tenant{
		Id:         source.Id,
		CreateTime: source.CreateTime,
		Name:       source.Name,
		Email:      source.Email,
		Status:     source.Status,
	}
}
