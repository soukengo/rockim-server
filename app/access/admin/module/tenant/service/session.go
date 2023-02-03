package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/admin/tenant/v1"
	"rockimserver/apis/rockim/api/admin/tenant/v1/types"
	"rockimserver/app/access/admin/module/tenant/biz"
	"rockimserver/app/access/admin/module/tenant/service/converter"
)

type SessionService struct {
	uc *biz.SessionUseCase
}

func NewSessionService(uc *biz.SessionUseCase) *SessionService {
	return &SessionService{uc: uc}
}

func (s *SessionService) Check(ctx context.Context, in *v1.SessionCheckRequest) (reply *v1.SessionCheckResponse, err error) {
	session, err := s.uc.Check(ctx)
	if err != nil {
		return
	}
	reply = &v1.SessionCheckResponse{
		Session: &types.SessionUser{Id: session.ID, Name: session.Name, AvatarUrl: session.AvatarUrl},
	}
	return
}

func (s *SessionService) ListResource(ctx context.Context, in *v1.SessionListResourceRequest) (reply *v1.SessionListResourceResponse, err error) {
	list, err := s.uc.ListResources(ctx)
	if err != nil {
		return
	}
	resources := make([]*types.SysTenantResource, len(list))
	for i, item := range list {
		resources[i] = converter.ToSysTenantResource(item)
	}
	reply = &v1.SessionListResourceResponse{
		List: converter.BuildTenantResourceTree(resources),
	}
	return
}
