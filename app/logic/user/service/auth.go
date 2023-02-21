package service

import (
	"context"
	"rockimserver/apis/rockim/service/user/v1"
	"rockimserver/app/logic/user/biz"
	"rockimserver/app/logic/user/biz/options"
)

type AuthService struct {
	uc *biz.AuthUseCase
	v1.UnimplementedAuthAPIServer
}

func NewAuthService(uc *biz.AuthUseCase) *AuthService {
	return &AuthService{uc: uc}
}

func (s *AuthService) CreateAuthCode(ctx context.Context, in *v1.AuthCodeRequest) (out *v1.AuthCodeResponse, err error) {
	code, err := s.uc.CreateAuthCode(ctx, &options.AuthCodeCreateOptions{ProductId: in.ProductId, Account: in.Account})
	if err != nil {
		return
	}
	return &v1.AuthCodeResponse{AuthCode: code.Code, ExpireTime: code.ExpireTime}, nil
}

func (s *AuthService) Login(ctx context.Context, in *v1.LoginRequest) (out *v1.LoginResponse, err error) {
	token, err := s.uc.Login(ctx, &options.LoginOptions{ProductId: in.ProductId, AuthCode: in.AuthCode})
	if err != nil {
		return
	}
	return &v1.LoginResponse{Token: token.Token, ExpireTime: token.ExpireTime}, nil
}

func (s *AuthService) CheckToken(ctx context.Context, in *v1.TokenCheckRequest) (out *v1.TokenCheckResponse, err error) {
	uid, err := s.uc.CheckToken(ctx, &options.TokenCheckOptions{ProductId: in.ProductId, Token: in.Token})
	if err != nil {
		return
	}
	return &v1.TokenCheckResponse{Uid: uid}, nil
}
