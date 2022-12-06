package service

import (
	"context"
	"rockim/api/rockim/client/v1"
	clienttypes "rockim/api/rockim/client/v1/types"
	"rockim/api/rockim/service/user/v1/types"
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
	return &v1.UserRegisterResponse{User: &clienttypes.User{
		Id:         user.Id,
		CreateTime: user.CreateTime,
		AppId:      user.AppId,
		Bucket:     user.Bucket,
		Account:    user.Account,
		Name:       user.Name,
		Fields:     user.Fields,
		Status:     user.Status,
	}}, nil
}
