package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/admin/manager/v1"
	apiTypes "rockimserver/apis/rockim/api/admin/manager/v1/types"
	"rockimserver/app/access/admin/module/manager/biz"
	"rockimserver/app/access/admin/module/manager/service/converter"
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
		Session: &apiTypes.SessionUser{Id: session.ID, Name: session.Name, AvatarUrl: session.AvatarUrl},
	}
	return
}

func (s *SessionService) ListResource(ctx context.Context, in *v1.SessionListResourceRequest) (reply *v1.SessionListResourceResponse, err error) {
	list, err := s.uc.ListResources(ctx)
	if err != nil {
		return
	}
	resources := make([]*apiTypes.SysResource, len(list))
	for i, item := range list {
		resources[i] = converter.ToManagerSysResource(item)
	}
	reply = &v1.SessionListResourceResponse{
		List: converter.BuildResourceTree(resources),
	}
	return
}
