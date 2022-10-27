package service

import (
	"context"
	v1 "rockim/api/access/gateway/v1"
	"rockim/app/access/gateway/biz"
)

type UserService struct {
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Register(ctx context.Context, request *v1.UserRegisterRequest) (*v1.UserRegisterResponse, error) {
	s.uc.Register()
	return nil, nil
}
