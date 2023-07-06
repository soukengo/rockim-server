package cache

import (
	"github.com/soukengo/gopkg/component/cache"
)

const (
	keyGroup              cache.Key = "group"
	keyGroupCustomGroupId cache.Key = "group.custom_group_id"

	keyGroupMemberSet cache.Key = "group.member.set"
	keyUserGroupSet   cache.Key = "user.group.set"
	keyGroupMember    cache.Key = "group.member"

	keyChatRoomMemberSet cache.Key = "chatroom.member.set"
	keyUserChatRoomSet   cache.Key = "user.chatroom.set"
	keyChatRoomMember    cache.Key = "chatroom.member"
)
