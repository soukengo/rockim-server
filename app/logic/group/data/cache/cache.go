package cache

import (
	"github.com/soukengo/gopkg/component/cache"
)

const (
	keyGroup      cache.Key = "group"
	keyGroupBizId cache.Key = "group.biz_id"

	keyGroupMemberSet  cache.Key = "group.member.set"
	keyGroupMemberInfo cache.Key = "group.member.info"

	keyChatRoomMemberSet  cache.Key = "chatroom.member.set"
	keyChatRoomMemberHash cache.Key = "chatroom.member.hash"
)
