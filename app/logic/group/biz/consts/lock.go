package consts

import "rockimserver/pkg/component/lock"

const (
	LockKeyChatRoomCreate  lock.Key = "ChatRoom.Create"
	LockKeyChatRoomDismiss lock.Key = "ChatRoom.Dismiss"
	LockKeyChatRoomJoin    lock.Key = "ChatRoom.Join"
	LockKeyChatRoomQuit    lock.Key = "ChatRoom.Quit"
)
