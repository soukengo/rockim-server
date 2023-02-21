package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/openapi/v1"
	"rockimserver/app/access/gateway/module/openapi/biz"
	"rockimserver/app/access/gateway/module/openapi/biz/options"
)

type AuthService struct {
	uc *biz.AuthUseCase
}

func NewAuthService(uc *biz.AuthUseCase) *AuthService {
	return &AuthService{uc: uc}
}

func (s *AuthService) CreateAuthCode(ctx context.Context, in *v1.AuthCodeRequest) (out *v1.AuthCodeResponse, err error) {
	ret, err := s.uc.CreateAuthCode(ctx, &options.AuthCodeCreateOptions{ProductId: in.ProductId, Account: in.Account})
	if err != nil {
		return
	}
	out = &v1.AuthCodeResponse{AuthCode: ret.Code, ExpireTime: ret.ExpireTime}
	return
}
