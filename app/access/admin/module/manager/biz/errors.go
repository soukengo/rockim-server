package biz

import (
	"rockim/api/rockim/shared/reasons"
	"rockim/pkg/errors"
)

var (
	ErrPasswordIncorrect    = errors.Conflict(reasons.AdminSysUserErrorReason_SYS_PASSWORD_INCORRECT.String(), "Password is incorrect")
	ErrAuthorizationInvalid = errors.Conflict(reasons.AdminSysUserErrorReason_SYS_AUTHORIZATION_INVALID.String(), "Authorization invalid")
)
