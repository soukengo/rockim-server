package options

type ClientAuthOptions struct {
	// 需要客户端传的参数
	ProductId string
	Token     string
	// 以下参数服务端自动填充
	ServerId  string
	ChannelId string
}

type ClientHeartBeatOptions struct {
	// 以下参数服务端自动填充
	ProductId string
	Uid       string
	ServerId  string
	ChannelId string
}

type ClientDisConnectOptions struct {
	// 以下参数服务端自动填充
	ProductId string
	Uid       string
	ServerId  string
	ChannelId string
}
