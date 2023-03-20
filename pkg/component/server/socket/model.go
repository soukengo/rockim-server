package socket

import "rockimserver/pkg/component/server/socket/packet"

type PushGroupPacket struct {
	GroupId string
	Packet  packet.IPacket
}
