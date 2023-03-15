package biz

import (
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/pkg/errors"
)

var (
	ErrTargetNotSupported    = errors.BadRequest(reasons.Message_TARGET_NOT_SUPPORTED.String(), "不支持的消息接收目标")
	ErrGenerateMessageFailed = errors.InternalServer(reasons.ErrorReason_UN_SPECIFIED.String(), "生成消息序号失败")
)
