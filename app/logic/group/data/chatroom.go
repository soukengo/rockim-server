package data

import (
	"context"
	"rockimserver/app/logic/group/biz"
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/database/redis"
)

type chatRoomRepo struct {
	mgo *mongo.Client
	rds *redis.Client
}

func NewChatRoomRepo(mgo *mongo.Client, rds *redis.Client) biz.ChatRoomRepo {
	return &chatRoomRepo{mgo: mgo, rds: rds}
}

func (r *chatRoomRepo) AddMember(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
