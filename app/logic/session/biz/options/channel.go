package options

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

type OnlineCheckOptions struct {
	ProductId string
	Uid       string
}
type OnlineBatchCheckOptions struct {
	ProductId string
	Uids      []string
}

type UserChannelListOptions struct {
	ProductId string
	Uids      []string
}
