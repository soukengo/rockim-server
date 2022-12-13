package tenant

import (
	"context"
	v1 "rockim/api/rockim/admin/tenant/v1"
	"rockim/app/access/admin/biz/manager"
)

type AuthService struct {
	uc *manager.AuthUseCase
}

func (s *AuthService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthService(uc *manager.AuthUseCase) *AuthService {
	return &AuthService{uc: uc}
}
