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

type DispatchOptions struct {
	// 连接ID
	ChannelIds []string
	// 数据类型
	DataType enums.Comet_DataType
	// 推送消息
	Push *types.PushMessage
	// 控制消息
	Control *types.ControlMessage
}

type ControlOptions struct {
	// 控制类型
	ControlType types.ControlMessage_ControlType
	// 连接id
	ChannelIds []string
	// 推送内容
	Body []byte
}

type DispatchRoomOptions struct {
	// 连接ID
	Room *types.Room
	// 推送消息
	Push *types.PushMessage
}
