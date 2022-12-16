package manager

import (
	"context"
	v1 "rockim/api/rockim/admin/manager/v1"
	"rockim/app/access/admin/biz/manager"
)

type AuthService struct {
	uc *manager.AuthUseCase
}

func NewAuthService(uc *manager.AuthUseCase) *AuthService {
	return &AuthService{uc: uc}
}

func (s *AuthService) Login(ctx context.Context, in *v1.LoginRequest) (reply *v1.LoginResponse, err error) {
	token, err := s.uc.Login(ctx, in.Account, in.Password)
	if err != nil {
		return
	}
	return &v1.LoginResponse{Token: token}, nil
}
