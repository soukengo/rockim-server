package options

type UserRegisterOptions struct {
	// 所属应用
	ProductId string
	// 用户所属分桶（可选）
	Bucket string
	// 用户账户名，由接入方指定
	Account string
	// 用户名称
	Name string
	// 用户自定义字段
	Fields map[string]string
}
