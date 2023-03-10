package data

import (
	"context"
	"rockimserver/app/logic/message/biz"
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/database/redis"
)

type messageRepo struct {
	mongoCli *mongo.Client
	redisCli *redis.Client
}

func NewMessageRepo(mongoCli *mongo.Client, redisCli *redis.Client) biz.MessageRepo {
	return &messageRepo{mongoCli: mongoCli, redisCli: redisCli}
}

func (r *messageRepo) Save(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
