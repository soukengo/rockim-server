package options

type ChatRoomCreateOptions struct {
	ProductId string
	BizId     string // cp自定义群组id
	Name      string
	IconUrl   string
	Fields    map[string]string
	Owner     string
}
