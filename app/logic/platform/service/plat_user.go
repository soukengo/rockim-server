package service

import (
	"context"
	"rockim/api/rockim/service/platform/v1"
	"rockim/app/logic/platform/biz"
)

type PlatUserService struct {
	v1.UnimplementedPlatUserAPIServer
	uc *biz.PlatUserUseCase
}

// NewPlatUserService new a plat user service.
func NewPlatUserService(uc *biz.PlatUserUseCase) *PlatUserService {
	return &PlatUserService{uc: uc}
}

func (s *PlatUserService) Find(ctx context.Context, req *v1.PlatUserFindRequest) (*v1.PlatUserFindResponse, error) {
	u, err := s.uc.Find(ctx, req.Account)
	if err != nil {
		return nil, err
	}
	return &v1.PlatUserFindResponse{User: u}, nil
}
