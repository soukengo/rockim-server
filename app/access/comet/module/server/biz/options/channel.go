package options

import (
	"rockimserver/apis/rockim/service/comet/v1/types"
	"rockimserver/apis/rockim/shared/enums"
)

type PushOptions struct {
	// 推送类型
	Operation enums.Comet_PushOperation
	// 连接id
	ChannelIds []string
	// 推送内容
	Body []byte
}

type PushRoomOptions struct {
	// 推送类型
	Operation enums.Comet_PushOperation
	// 房间ID
	Room *types.Room
	// 推送内容
	Body []byte
}
