package service

import (
	"rockim/app/logic/user/biz"
)

type AuthService struct {
	uc *biz.UserUseCase
}
