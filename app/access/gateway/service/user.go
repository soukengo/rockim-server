package service

import (
	"context"
	v1 "rockim/api/access/gateway/csdk/v1"
	v1Types "rockim/api/access/gateway/csdk/v1/types"
	"rockim/api/logic/user/v1/types"
	"rockim/app/access/gateway/biz"
)

type UserService struct {
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Register(ctx context.Context, req *v1.UserRegisterRequest) (*v1.UserRegisterResponse, error) {
	user, err := s.uc.RegisterUser(ctx, &types.User{
		AppId:   req.AppId,
		Bucket:  req.Bucket,
		Account: req.Account,
		Name:    req.Name,
		Fields:  req.Fields,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserRegisterResponse{User: &v1Types.User{
		Id:         user.Id,
		CreateTime: user.CreateTime,
		AppId:      user.AppId,
		Bucket:     user.Bucket,
		Account:    user.Account,
		Name:       user.Name,
		Fields:     user.Fields,
		Status:     v1Types.UserStatus(user.Status),
	}}, nil
}
