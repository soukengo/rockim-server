package service

import (
	"context"
	"rockimserver/apis/rockim/api/client/v1"
	"rockimserver/app/access/gateway/module/client/biz"
	"rockimserver/app/access/gateway/module/client/service/converter"
)

type UserService struct {
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Find(ctx context.Context, in *v1.UserFindRequest) (out *v1.UserFindResponse, err error) {
	user, err := s.uc.Find(ctx, in.Base.ProductId, in.Account)
	if err != nil {
		return nil, err
	}
	return &v1.UserFindResponse{User: converter.ToClientUser(user)}, nil
}
