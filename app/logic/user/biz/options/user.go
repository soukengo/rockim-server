package options

type UserCreateOptions struct {
	ProductId string            // 所属应用
	Account   string            // 用户账户名，由接入方指定
	Name      string            // 用户名称
	AvatarUrl string            // 头像地址
	Fields    map[string]string // 客户自定义字段
}
