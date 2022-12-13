package service

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/app/logic/platform/biz"
)

type PlatRoleService struct {
	v1.UnimplementedPlatRoleAPIServer
	uc *biz.PlatRoleUseCase
}

// NewPlatRoleService new a plat user service.
func NewPlatRoleService(uc *biz.PlatRoleUseCase) *PlatRoleService {
	return &PlatRoleService{uc: uc}
}

func (s *PlatRoleService) ListResourceId(ctx context.Context, request *v1.PlatRoleResourceIdListRequest) (*v1.PlatRoleResourceIdListResponse, error) {
	list, err := s.uc.ListResourceId(ctx, request.RoleIds)
	if err != nil {
		return nil, err
	}
	return &v1.PlatRoleResourceIdListResponse{List: list}, nil
}
