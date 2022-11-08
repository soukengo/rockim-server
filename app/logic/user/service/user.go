package service

import (
	"context"
	"rockim/api/logic/user/v1/types"

	v1 "rockim/api/logic/user/v1"
	"rockim/app/logic/user/biz"
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
		AppId:   in.AppId,
		Bucket:  in.Bucket,
		Account: in.Account,
		Name:    in.Name,
		Fields:  in.Fields,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserRegisterResponse{User: user}, nil
}
