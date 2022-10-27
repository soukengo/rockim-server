package biz

// UserStatus 用户状态
type UserStatus int32

const (
	UserStatus_USER_STATUS_INVALID UserStatus = 0 // 无效
	UserStatus_USER_STATUS_NORMAL  UserStatus = 1 // 正常
)

type User struct {
	Id         string            `json:"id,omitempty"`          // 用户ID
	CreateTime int64             `json:"create_time,omitempty"` // 创建时间
	AppId      string            `json:"app_id,omitempty"`      // 所属应用
	ChannelId  string            `json:"channel_id,omitempty"`  // 客户自定义渠道（可选）
	Account    string            `json:"account,omitempty"`     // 用户账户名，由接入方指定
	Name       string            `json:"name,omitempty"`        // 用户名称
	Status     UserStatus        `json:"status,omitempty"`      // 状态
	Fields     map[string]string `json:"fields,omitempty"`      // 客户自定义字段
}
