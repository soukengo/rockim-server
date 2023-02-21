package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/openapi/v1"
	"rockimserver/app/access/gateway/module/openapi/biz"
	"rockimserver/app/access/gateway/module/openapi/biz/options"
)

type UserService struct {
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Register(ctx context.Context, req *v1.UserRegisterRequest) (*v1.UserRegisterResponse, error) {
	user, err := s.uc.Register(ctx, &options.UserRegisterOptions{
		ProductId: req.ProductId,
		Bucket:    req.Bucket,
		Account:   req.Account,
		Name:      req.Name,
		Fields:    req.Fields,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserRegisterResponse{Uid: user.Id}, nil
}
