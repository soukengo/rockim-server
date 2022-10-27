package service

import (
	v1 "rockim/api/logic/user/v1"
	"rockim/app/logic/user/biz"
)

type AuthService struct {
	v1.UnimplementedUserServer
	uc *biz.UserUseCase
}
