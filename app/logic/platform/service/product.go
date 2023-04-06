package service

import (
	"context"
	"github.com/soukengo/gopkg/util/copier"
	v1 "rockimserver/apis/rockim/service/platform/v1"
	"rockimserver/app/logic/platform/biz"
	"rockimserver/app/logic/platform/biz/options"
)

type ProductService struct {
	v1.UnimplementedProductAPIServer
	uc *biz.ProductUseCase
}

func NewProductService(uc *biz.ProductUseCase) *ProductService {
	return &ProductService{uc: uc}
}

func (s *ProductService) Create(ctx context.Context, in *v1.ProductCreateRequest) (out *v1.ProductCreateResponse, err error) {
	req := &options.ProductCreateOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	out = &v1.ProductCreateResponse{}
	return
}

func (s *ProductService) Update(ctx context.Context, in *v1.ProductUpdateRequest) (out *v1.ProductUpdateResponse, err error) {
	req := &options.ProductUpdateOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	out = &v1.ProductUpdateResponse{}
	return
}

func (s *ProductService) Find(ctx context.Context, in *v1.ProductFindRequest) (out *v1.ProductFindResponse, err error) {
	p, err := s.uc.FindById(ctx, in.Id)
	if err != nil {
		return
	}
	out = &v1.ProductFindResponse{Product: p}
	return
}

func (s *ProductService) ListByTenant(ctx context.Context, in *v1.ProductListByTenantRequest) (out *v1.ProductListByTenantResponse, err error) {
	list, err := s.uc.ListByTenant(ctx, in.TenantId)
	if err != nil {
		return
	}
	out = &v1.ProductListByTenantResponse{List: list}
	return
}
