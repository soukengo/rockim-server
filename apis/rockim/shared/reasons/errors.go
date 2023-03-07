package reasons

import "rockimserver/pkg/errors"

var (
	ErrUnSpecified = errors.BadRequest(ErrorReason_UN_SPECIFIED.String(), "未知错误")
)
