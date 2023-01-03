package tenant

import (
	"context"
	v1 "rockim/api/rockim/admin/tenant/v1"
	"rockim/app/access/admin/biz/tenant"
)

type AuthService struct {
	uc *tenant.AuthUseCase
}

func NewAuthService(uc *tenant.AuthUseCase) *AuthService {
	return &AuthService{uc: uc}
}

func (s *AuthService) Login(ctx context.Context, in *v1.LoginRequest) (reply *v1.LoginResponse, err error) {
	token, err := s.uc.Login(ctx, in.Account, in.Password)
	if err != nil {
		return
	}
	return &v1.LoginResponse{Token: token}, nil
}
