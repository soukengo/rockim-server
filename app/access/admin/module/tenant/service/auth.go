package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/admin/tenant/v1"
	"rockimserver/app/access/admin/module/tenant/biz"
)

type AuthService struct {
	uc *biz.AuthUseCase
}

func NewAuthService(uc *biz.AuthUseCase) *AuthService {
	return &AuthService{uc: uc}
}

func (s *AuthService) Login(ctx context.Context, in *v1.LoginRequest) (reply *v1.LoginResponse, err error) {
	token, err := s.uc.Login(ctx, in.Account, in.Password)
	if err != nil {
		return
	}
	return &v1.LoginResponse{Token: token}, nil
}
