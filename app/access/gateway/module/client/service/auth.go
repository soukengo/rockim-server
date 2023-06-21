package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/client/v1/protocol/http"
	"rockimserver/app/access/gateway/module/client/biz"
	"rockimserver/app/access/gateway/module/client/biz/options"
	"rockimserver/app/access/gateway/module/client/service/converter"
)

type AuthService struct {
	uc *biz.AuthUseCase
}

func NewAuthService(uc *biz.AuthUseCase) *AuthService {
	return &AuthService{uc: uc}
}

func (s *AuthService) Login(ctx context.Context, in *v1.LoginRequest) (out *v1.LoginResponse, err error) {
	ret, user, err := s.uc.Login(ctx, &options.LoginOptions{ProductId: in.Base.ProductId, AuthCode: in.AuthCode})
	if err != nil {
		return
	}
	out = &v1.LoginResponse{Token: ret.Token, ExpireTime: ret.ExpireTime, User: converter.ToClientUser(user)}
	return
}
