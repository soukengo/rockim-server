package manager

import (
	"context"
	v1 "rockim/api/rockim/admin/manager/v1"
	"rockim/api/rockim/admin/manager/v1/types"
	"rockim/app/access/admin/biz/manager"
)

type SessionService struct {
	uc *manager.SessionUseCase
}

func NewSessionService(uc *manager.SessionUseCase) *SessionService {
	return &SessionService{uc: uc}
}

func (s *SessionService) Check(ctx context.Context, request *v1.SessionCheckRequest) (reply *v1.SessionCheckResponse, err error) {
	session, err := s.uc.Check(ctx)
	if err != nil {
		return
	}
	reply = &v1.SessionCheckResponse{
		Session: &types.SessionUser{Id: session.ID, Name: session.Name, AvatarUrl: session.AvatarUrl},
	}
	return
}

func (s *SessionService) ListResource(ctx context.Context, request *v1.SessionListResourceRequest) (reply *v1.SessionListResourceResponse, err error) {
	list, err := s.uc.ListResources(ctx)
	resources := make([]*types.PlatResource, len(list))
	for i, item := range list {
		resources[i] = &types.PlatResource{
			Id:         item.Id,
			CreateTime: item.CreateTime,
			Category:   item.Category,
			Name:       item.Name,
			ParentId:   item.ParentId,
			Url:        item.Url,
			Icon:       item.Icon,
			Level:      item.Level,
			Leaf:       item.Leaf,
			Order:      item.Order,
		}
	}
	reply = &v1.SessionListResourceResponse{
		List: buildResourceTree(resources),
	}
	return
}
