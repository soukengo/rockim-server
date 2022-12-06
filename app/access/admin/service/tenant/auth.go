package tenant

import (
	"context"
	v1 "rockim/api/rockim/admin/tenant/v1"
	"rockim/app/access/admin/biz"
)

type AuthService struct {
	uc *biz.AuthUseCase
}

func (s *AuthService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthService(uc *biz.AuthUseCase) *AuthService {
	return &AuthService{uc: uc}
}
