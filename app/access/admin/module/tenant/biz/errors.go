package biz

import (
	"rockim/api/rockim/shared/reasons"
	"rockim/pkg/errors"
)

var (
	ErrPasswordIncorrect    = errors.Conflict(reasons.AdminTenantErrorReason_TENANT_PASSWORD_INCORRECT.String(), "Password is incorrect")
	ErrAuthorizationInvalid = errors.Conflict(reasons.AdminTenantErrorReason_TENANT_AUTHORIZATION_INVALID.String(), "Authorization invalid")
)
