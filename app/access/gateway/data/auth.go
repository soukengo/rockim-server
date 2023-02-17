package data

import (
	v1 "rockimserver/apis/rockim/service/user/v1"
)

type authRepo struct {
	ac v1.AuthAPIClient
}

//func NewUserRepo(uac v1.UserAPIClient) biz.UserRepo {
//	return &userRepo{uac: uac}
//}
