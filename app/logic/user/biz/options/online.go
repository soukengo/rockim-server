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

type OnlineCheckOptions struct {
	ProductId string
	Uid       string
}
type OnlineBatchCheckOptions struct {
	ProductId string
	Uids      []string
}

type OnlineUserListOptions struct {
	ProductId string
	Uids      []string
}
