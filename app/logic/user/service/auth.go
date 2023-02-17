package service

import (
	"context"
	"rockimserver/apis/rockim/service/user/v1"
	"rockimserver/app/logic/user/biz"
)

type AuthService struct {
	uc *biz.UserUseCase
	v1.UnimplementedAuthAPIServer
}

func NewAuthService(uc *biz.UserUseCase) *AuthService {
	return &AuthService{uc: uc}
}

func (s *AuthService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginResponse, error) {
	return &v1.LoginResponse{Uid: "123456", Token: "abcdefghijklmnopqrstuvwxyz"}, nil
}
