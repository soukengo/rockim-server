package options

type ChatRoomCreateOptions struct {
	ProductId     string
	CustomGroupId string // cp自定义群组id
	Name          string
	IconUrl       string
	Fields        map[string]string
}

type ChatRoomDismissOptions struct {
	ProductId string
	GroupId   string
}
