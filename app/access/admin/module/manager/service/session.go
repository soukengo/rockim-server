package service

import (
	"context"
	v1 "rockim/api/rockim/admin/manager/v1"
	apiTypes "rockim/api/rockim/admin/manager/v1/types"
	"rockim/app/access/admin/module/manager/biz"
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
		resources[i] = convertResource(item)
	}
	reply = &v1.SessionListResourceResponse{
		List: buildResourceTree(resources),
	}
	return
}
