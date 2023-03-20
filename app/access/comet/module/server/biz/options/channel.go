package options

import "rockimserver/apis/rockim/shared/enums"

type PushOptions struct {
	// 推送类型
	Operation enums.Network_PushOperation
	// 连接id
	ChannelIds []string
	// 推送内容
	Body []byte
}

type PushGroupOptions struct {
	// 推送类型
	Operation enums.Network_PushOperation
	// 组ID
	GroupId string
	// 推送内容
	Body []byte
}
