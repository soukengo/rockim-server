package options

import (
	"rockimserver/apis/rockim/shared/enums"
)

type ChatRoomCreateOptions struct {
	ProductId     string
	CustomGroupId string // cp自定义群组id
	Name          string
	IconUrl       string
	Fields        map[string]string
	Owner         string
}

type ChatRoomDismissOptions struct {
	ProductId string
	GroupId   string
}

type ChatRoomJoinOptions struct {
	ProductId string
	GroupId   string
	Uid       string
}

type ChatRoomQuitOptions struct {
	ProductId string
	GroupId   string
	Uid       string
}

type ChatRoomMemberRemoveOptions struct {
	ProductId string
	GroupId   string
	Uid       string
	TargetUid string
}

type ChatRoomMemberAddItem struct {
	Uid    string
	Role   enums.Group_MemberRole
	Fields map[string]string
}
