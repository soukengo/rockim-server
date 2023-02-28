package cache

import "rockimserver/pkg/component/database/redis"

type ChatRoomData struct {
	rds *redis.Client
}

func NewChatRoomData(rds *redis.Client) *ChatRoomData {
	return &ChatRoomData{rds: rds}
}
