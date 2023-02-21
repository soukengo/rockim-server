package biz

import (
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/pkg/errors"
)

var (
	ErrPasswordIncorrect    = errors.BadRequest(reasons.Admin_SYS_PASSWORD_INCORRECT.String(), "Password is incorrect")
	ErrAuthorizationInvalid = errors.BadRequest(reasons.Admin_SYS_AUTHORIZATION_INVALID.String(), "Authorization invalid")
)
