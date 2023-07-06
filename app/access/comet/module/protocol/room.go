package protocol

import "rockimserver/apis/rockim/service/comet/v1/types"

const (
	roomIdSep = "#"
)

func EncodeRoomId(room *types.Room) string {
	return room.RoomType.String() + roomIdSep + room.BizId
}
