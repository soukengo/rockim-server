package service

import (
	"context"
	"rockimserver/api/rockim/client/v1"
	clienttypes "rockimserver/api/rockim/client/v1/types"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/app/access/gateway/biz"
)

type UserService struct {
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Register(ctx context.Context, req *v1.UserRegisterRequest) (*v1.UserRegisterResponse, error) {
	user, err := s.uc.RegisterUser(ctx, &types.User{
		ProductId: req.ProductId,
		Bucket:    req.Bucket,
		Account:   req.Account,
		Name:      req.Name,
		Fields:    req.Fields,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserRegisterResponse{User: &clienttypes.User{
		Id:         user.Id,
		CreateTime: user.CreateTime,
		ProductId:  user.ProductId,
		Bucket:     user.Bucket,
		Account:    user.Account,
		Name:       user.Name,
		Fields:     user.Fields,
		Status:     user.Status,
	}}, nil
}
