package biz

import (
	"github.com/soukengo/gopkg/errors"
	"rockimserver/apis/rockim/shared/reasons"
)

var (
	ErrTargetNotSupported    = errors.BadRequest(reasons.Message_TARGET_NOT_SUPPORTED.String(), "不支持的消息接收目标")
	ErrGenerateMessageFailed = errors.InternalServer(reasons.ErrorReason_UN_SPECIFIED.String(), "生成消息序号失败")
	ErrNotGroupMember        = errors.InternalServer(reasons.Message_NOT_GROUP_MEMBER.String(), "非群成员不允许发送消息")
)
