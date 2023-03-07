package options

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
