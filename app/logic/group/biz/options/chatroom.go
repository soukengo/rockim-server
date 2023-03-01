package options

import "rockimserver/apis/rockim/shared"

type ChatRoomCreateOptions struct {
	ProductId     string
	CustomGroupId string // cp自定义群组id
	Name          string
	IconUrl       string
	Fields        map[string]string
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

type ChatRoomMemberCheckOptions struct {
	ProductId string
	GroupId   string
	Uid       string
}

type ChatRoomMemberFindOptions struct {
	ProductId string
	GroupId   string
	Uid       string
}
type ChatRoomMemberListOptions struct {
	ProductId string
	GroupId   string
	Uids      []string
}

type ChatRoomMemberIdPaginateOptions struct {
	ProductId string
	GroupId   string
	Paginate  *shared.Paginating
}
