package service

import (
	"context"
	"rockimserver/apis/rockim/service/user/v1"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/app/logic/user/biz"
)

type UserService struct {
	v1.UnimplementedUserAPIServer
	uc *biz.UserUseCase
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Register(ctx context.Context, in *v1.UserRegisterRequest) (*v1.UserRegisterResponse, error) {
	user, err := s.uc.Register(ctx, &types.User{
		ProductId: in.ProductId,
		Account:   in.Account,
		Name:      in.Name,
		AvatarUrl: in.AvatarUrl,
		Fields:    in.Fields,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserRegisterResponse{User: user}, nil
}

func (s *UserService) Find(ctx context.Context, in *v1.UserFindRequest) (out *v1.UserFindResponse, err error) {
	user, err := s.uc.Find(ctx, in.ProductId, in.Account)
	if err != nil {
		return
	}
	return &v1.UserFindResponse{User: user}, nil
}
