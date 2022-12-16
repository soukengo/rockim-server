package service

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/app/logic/platform/biz"
	"rockim/pkg/util/copier"
)

type PlatResourceService struct {
	v1.UnimplementedPlatResourceAPIServer
	uc *biz.PlatResourceUseCase
}

// NewPlatResourceService new a plat resource service.
func NewPlatResourceService(uc *biz.PlatResourceUseCase) *PlatResourceService {
	return &PlatResourceService{uc: uc}
}

func (s *PlatResourceService) Create(ctx context.Context, in *v1.PlatResourceCreateRequest) (reply *v1.PlatResourceCreateResponse, err error) {
	req := &biz.PlatResourceCreateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.PlatResourceCreateResponse{}
	return
}

func (s *PlatResourceService) Update(ctx context.Context, in *v1.PlatResourceUpdateRequest) (reply *v1.PlatResourceUpdateResponse, err error) {
	req := &biz.PlatResourceUpdateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.PlatResourceUpdateResponse{}
	return
}

func (s *PlatResourceService) Delete(ctx context.Context, in *v1.PlatResourceDeleteRequest) (reply *v1.PlatResourceDeleteResponse, err error) {
	req := &biz.PlatResourceDeleteRequest{Id: in.Id}
	err = s.uc.Delete(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.PlatResourceDeleteResponse{}
	return
}

func (s *PlatResourceService) List(ctx context.Context, request *v1.PlatResourceListRequest) (reply *v1.PlatResourceListResponse, err error) {
	list, err := s.uc.List(ctx)
	if err != nil {
		return
	}
	return &v1.PlatResourceListResponse{List: list}, nil
}

func (s *PlatResourceService) ListByIds(ctx context.Context, request *v1.PlatResourceListByIdsRequest) (reply *v1.PlatResourceListByIdsResponse, err error) {
	list, err := s.uc.ListByIds(ctx, request.Ids)
	if err != nil {
		return
	}
	reply = &v1.PlatResourceListByIdsResponse{List: list}
	return
}
