package biz

import (
	v1 "rockim/api/rockim/admin/tenant/v1"
	"rockim/pkg/errors"
)

var (
	ErrPasswordIncorrect    = errors.Conflict(v1.ErrorReason_PASSWORD_INCORRECT.String(), "Password is incorrect")
	ErrAuthorizationInvalid = errors.Conflict(v1.ErrorReason_AUTHORIZATION_INVALID.String(), "Authorization invalid")
)
