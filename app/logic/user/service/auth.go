package service

import (
	"rockimserver/app/logic/user/biz"
)

type AuthService struct {
	uc *biz.UserUseCase
}
