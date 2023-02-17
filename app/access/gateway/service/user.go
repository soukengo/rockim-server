package service

import (
	"context"
	"rockimserver/apis/rockim/api/client/v1"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/app/access/gateway/biz"
	"rockimserver/app/access/gateway/service/converter"
)

type UserService struct {
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Register(ctx context.Context, req *v1.UserRegisterRequest) (*v1.UserRegisterResponse, error) {
	user, err := s.uc.Register(ctx, &types.User{
		ProductId: req.ProductId,
		Bucket:    req.Bucket,
		Account:   req.Account,
		Name:      req.Name,
		Fields:    req.Fields,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserRegisterResponse{User: converter.ToClientUser(user)}, nil
}

func (s *UserService) Find(ctx context.Context, in *v1.UserFindRequest) (out *v1.UserFindResponse, err error) {
	user, err := s.uc.Find(ctx, in.ProductId, in.Account)
	if err != nil {
		return nil, err
	}
	return &v1.UserFindResponse{User: converter.ToClientUser(user)}, nil
}
