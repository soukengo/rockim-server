package service

import (
	"context"
	"rockim/api/rockim/service/platform/v1"
	"rockim/app/logic/platform/biz"
)

type PlatResourceService struct {
	v1.UnimplementedPlatResourceAPIServer
	uc *biz.PlatResourceUseCase
}

// NewPlatResourceService new a plat resource service.
func NewPlatResourceService(uc *biz.PlatResourceUseCase) *PlatResourceService {
	return &PlatResourceService{uc: uc}
}

func (s *PlatResourceService) List(ctx context.Context, request *v1.PlatResourceListRequest) (*v1.PlatResourceListResponse, error) {
	return &v1.PlatResourceListResponse{}, nil
}

func (s *PlatResourceService) ListByIds(ctx context.Context, request *v1.PlatResourceListByIdsRequest) (reply *v1.PlatResourceListByIdsResponse, err error) {
	list, err := s.uc.ListByIds(ctx, request.Ids)
	if err != nil {
		return
	}
	reply = &v1.PlatResourceListByIdsResponse{List: list}
	return
}
