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
		Bucket:    in.Bucket,
		Account:   in.Account,
		Name:      in.Name,
		Fields:    in.Fields,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserRegisterResponse{User: user}, nil
}
