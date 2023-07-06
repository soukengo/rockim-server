package options

type ChannelAuthOptions struct {
	// 需要客户端传的参数
	ProductId string
	Token     string
}

type OnlineAddOptions struct {
	ProductId string
	Token     string
	ServerId  string
	ChannelId string
}

type OnlineRefreshOptions struct {
	ProductId string
	Uid       string
	ServerId  string
	ChannelId string
}

type OnlineDeleteOptions struct {
	ProductId string
	Uid       string
	ServerId  string
	ChannelId string
}

type ListRoomOptions struct {
	ProductId string
	Uid       string
}
