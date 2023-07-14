package options

type ChannelAuthOptions struct {
	ProductId string
	Token     string
}

type ChannelAddOptions struct {
	ProductId string
	Uid       string
	ServerId  string
	ChannelId string
}

type ChannelRefreshOptions struct {
	ProductId string
	Uid       string
	ServerId  string
	ChannelId string
}

type ChannelDeleteOptions struct {
	ProductId string
	Uid       string
	ServerId  string
	ChannelId string
}

type ListRoomOptions struct {
	ProductId string
	Uid       string
}
